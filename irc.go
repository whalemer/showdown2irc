package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/xfix/showdown2irc/protocol"
)

var tokenRegexp = regexp.MustCompile(`:[^\r\n]*|[^\s:]+`)

type connection struct {
	tcp          net.Conn
	nickname     string
	showdown     *protocol.BotConnection
	loginData    protocol.LoginData
	nickObtained bool
	userObtained bool
	closing      bool
}

func (c *connection) parseIRCLine(tokens []string) {
	commandName := strings.ToUpper(tokens[0])
	if command, ok := ircCommands[commandName]; ok {
		command(c, tokens[1:])
	} else if len(tokens) >= 2 && len(tokens[1]) > 0 && tokens[1][0] == '#' {
		room := c.showdown.Room(protocol.RoomID(tokens[1][1:]))
		room.SendCommand(commandName, strings.Join(tokens[2:], ""))
	} else {
		c.showdown.SendGlobalCommand(commandName, strings.Join(tokens[1:], " "))
	}
}

func (c *connection) send(parts ...string) {
	result := toIRC(parts)
	log.Print(result)
	c.tcp.Write([]byte(result))
}

func (c *connection) sendGlobal(parts ...string) {
	newParts := make([]string, len(parts)+1)
	newParts[0] = "showdown"
	copy(newParts[1:], parts)
	c.send(newParts...)
}

func (c *connection) runShowdownCommand(command, argument string, room *protocol.Room) {
	if callback, ok := showdownCommands[command]; ok {
		callback(c, argument, room)
	}
}

func (c *connection) continueConnection() {
	showdownConnection, connectionSuccess, err := protocol.ConnectToServer(c.loginData, "showdown", c.runShowdownCommand)
	if err != nil {
		c.sendGlobal("NOTICE", "#", err.Error())
		c.close()
		return
	}
	c.showdown = showdownConnection
	select {
	case <-connectionSuccess:
		c.sendGlobal("NICK", c.nickname)
		c.sendGlobal("001", c.nickname, "Welcome to Showdown proxy!")
		c.sendGlobal("005", c.nickname, "PREFIX=(qraohv)~#&@%+")
		c.sendGlobal("422", c.nickname, "No MoTD here. Go on.")
	case <-time.After(10 * time.Second):
		c.sendGlobal("NOTICE", "#", "Authentication did not succeed in 10 seconds")
		c.close()
		return
	}
}

func (c *connection) close() {
	c.sendGlobal("QUIT", c.nickname)
	c.closing = true
	if c.showdown != nil {
		c.showdown.Close()
	}
}

var ircCommands = map[string]func(*connection, []string){
	"CAP": func(c *connection, command []string) {
		// Not implemented, does nothing
	},
	"PASS": func(c *connection, command []string) {
		c.loginData.Password = command[0]
	},
	"NICK": func(c *connection, command []string) {
		if c.userObtained && !c.nickObtained {
			c.continueConnection()
		}
		c.nickObtained = true
	},
	"USER": func(c *connection, command []string) {
		c.loginData.Nickname = command[3]
		c.nickname = escapeUser(command[3])
		if !c.userObtained && c.nickObtained {
			c.continueConnection()
		}
		c.userObtained = true
	},
	"USERHOST": func(c *connection, command []string) {
		for _, arg := range command {
			c.sendGlobal("302", c.nickname, escapeUserWithHost(arg))
		}
	},
	"PING": func(c *connection, command []string) {
		args := make([]string, len(command)+2)
		args[0] = "PONG"
		args[1] = "showdown"
		copy(args[2:], command)
		c.sendGlobal(args...)
	},
	"PRIVMSG": func(c *connection, command []string) {
		if command[0][0] == '#' {
			room := c.showdown.Room(protocol.RoomID(command[0][1:]))
			room.Reply(unescapeUser(command[1]))
		} else if command[1] != "NickServ" {
			c.showdown.SendGlobalCommand("pm", fmt.Sprintf("%s,%s", command[0], command[1]))
		}
	},
	"JOIN": func(c *connection, command []string) {
		for _, room := range strings.Split(command[0], ",") {
			c.showdown.SendGlobalCommand("join", room)
		}
	},
	"PART": func(c *connection, command []string) {
		room := c.showdown.Room(protocol.RoomID(command[0][1:]))
		room.SendCommand("part", "")
	},
	"MODE": func(c *connection, command []string) {
		if len(command) == 1 {
			c.sendGlobal("352", c.nickname, command[0], "+ntc")
		}
	},
	"QUIT": func(c *connection, command []string) {
		c.close()
	},
}

var rankMap = map[rune]byte{'~': 'g', '#': 'r', '&': 'a', '@': 'o', '%': 'h', '+': 'v'}

func escapeUser(name string) string {
	return strings.Replace(name, " ", "\u00A0", -1)
}

func unescapeUser(name string) string {
	return strings.Replace(name, "\u00A0", " ", -1)
}

// Some IRC clients expect host for an user during room joining operations. This generates a fake one for their purpose
func escapeUserWithHost(name string) string {
	return fmt.Sprintf("%s!%s@showdown", escapeUser(name), protocol.ToID(name))
}

func escapeRoom(room protocol.RoomID) string {
	return "#" + string(room)
}

func fromIRC(line string) []string {
	lines := tokenRegexp.FindAllString(line, -1)
	for i, line := range lines {
		if line[0] == ':' {
			lines[i] = line[1:]
		}
	}
	return lines
}

func toIRC(tokens []string) string {
	spaceTokenFound := false
	var result bytes.Buffer
	for i, token := range tokens {
		if spaceTokenFound {
			panic(errors.New("Tokens found after a token containing a space"))
		}
		if i == 0 {
			result.WriteByte(':')
		} else {
			result.WriteByte(' ')
			if token == "" || strings.Contains(token, " ") {
				result.WriteByte(':')
				spaceTokenFound = true
			}
		}
		result.WriteString(token)
	}
	result.WriteString("\r\n")
	return result.String()
}

func handleIRC(rawConnection net.Conn) {
	defer rawConnection.Close()
	lines := bufio.NewReader(rawConnection)
	var c connection

	c = connection{tcp: rawConnection}
	for !c.closing {
		line, err := lines.ReadString('\n')
		if err != nil {
			log.Print(err)
			return
		}
		tokens := fromIRC(line)
		log.Print(tokens)
		c.parseIRCLine(tokens)
	}
}

func listen() {
	socket, err := net.Listen("tcp", "localhost:6667")
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := socket.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleIRC(connection)
	}
}
