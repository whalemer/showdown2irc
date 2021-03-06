// showdown2irc - use Showdown chat with an IRC client
// Copyright (C) 2016 Konrad Borowski
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package showdown

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// LoginData represents authentication information for a bot
type LoginData struct {
	Nickname string
	Password string
	Rooms    []string
}

// BotConnection represents a websocket communication with a bot
type BotConnection struct {
	loginData       LoginData
	rooms           map[RoomID]*Room
	commandCallback func(command, argument string, room *Room)
	onSuccess       chan<- struct{}
	*connection
}

// Room retrieves a room an user is connected to.
//
// If a room doesn't exist, a placeholder room is provided on which
// operations can be provided, in event you need a room to run command
// in, but not actually need to be in it.
//
// This command doesn't handle room aliases.
func (bc *BotConnection) Room(id RoomID) *Room {
	if roomWithUsers, ok := bc.rooms[id]; ok {
		return roomWithUsers
	}
	return &Room{BotConnection: bc, ID: id}
}

func (bc *BotConnection) parseMessage(message string) {
	var roomID RoomID
	if message[0] == '>' {
		parts := strings.SplitN(message[1:], "\n", 2)
		roomID = RoomID(parts[0])
		message = parts[1]
	}
	for {
		messages := strings.SplitN(message, "\n", 2)
		currentMessage := messages[0]
		if currentMessage == "" {
			return
		} else if currentMessage[0] == '|' {
			parts := strings.SplitN(currentMessage[1:], "|", 2)
			command := parts[0]
			var argument string
			if len(parts) > 1 {
				argument = parts[1]
			}

			if handler, ok := serverCommandHandlers[command]; ok {
				handler(argument, bc.Room(roomID))
			}
			bc.commandCallback(command, argument, bc.Room(roomID))
		} else {
			bc.commandCallback("", currentMessage, bc.Room(roomID))
		}
		if len(messages) != 2 {
			return
		}
		message = messages[1]
	}
}

// SendGlobalCommand sends a command does not care about a room in which
// it is used.
func (bc *BotConnection) SendGlobalCommand(command string, value string) {
	bc.sendCommand(command, value, "")
}

// SendCommand uses a command in a room.
func (bc *BotConnection) sendCommand(command string, value string, room RoomID) {
	if room == "lobby" {
		room = ""
	}
	bc.write(fmt.Sprintf("%s|/%s %s", room, command, value))
}

// Say says a text message in a specified room.
func (bc *BotConnection) say(message string, room RoomID) {
	if message == "" {
		return
	}
	if message[0] == '/' {
		message = "/" + message
	} else if message[0] == '!' || strings.HasPrefix(message, ">> ") || strings.HasPrefix(message, ">>> ") {
		message = " " + message
	}
	bc.write(fmt.Sprintf("%s|%s", room, message))
}

func handleConnection(botConnection *BotConnection) {
	for message := range botConnection.messageChannel {
		log.Println(message)
		botConnection.parseMessage(message)
	}
}

// ConnectToServer connects to a Showdown server by using its
// client location or its name.
func ConnectToServer(loginData LoginData, name string, commandCallback func(command, argument string, room *Room)) (*BotConnection, <-chan struct{}, error) {
	conf, err := findConfiguration(name)
	if err != nil {
		return nil, nil, err
	}
	return ConnectToKnownServer(loginData, conf, commandCallback)
}

// ConnectToKnownServer connects to a Showdown server with known
// configuration.
func ConnectToKnownServer(loginData LoginData, conf ServerAddress, commandCallback func(command, argument string, room *Room)) (*BotConnection, <-chan struct{}, error) {
	connection, err := webSocketConnect(&conf)
	if err != nil {
		return nil, nil, err
	}
	onSuccess := make(chan struct{}, 1)
	botConnection := &BotConnection{
		connection:      connection,
		loginData:       loginData,
		rooms:           map[RoomID]*Room{},
		commandCallback: commandCallback,
		onSuccess:       onSuccess,
	}
	go handleConnection(botConnection)
	return botConnection, onSuccess, nil
}

var serverCommandHandlers = map[string]func(string, *Room){
	"challstr": challStr,
	"init":     initializeChatRoom,
	"title":    setTitle,
	"users":    setUsers,
	"j":        joinRoom,
	"J":        joinRoom,
	"l":        leaveRoom,
	"L":        leaveRoom,
	"N":        renameNick,
}

const actionURL = "https://play.pokemonshowdown.com/action.php"

func challStr(challenge string, room *Room) {
	botConnection := room.BotConnection
	loginData := botConnection.loginData

	parameters := url.Values{}
	parameters.Set("act", "login")
	parameters.Set("name", loginData.Nickname)
	parameters.Set("pass", loginData.Password)
	parameters.Set("challstr", challenge)

	res, err := http.PostForm(actionURL, parameters)
	if err != nil {
		log.Fatal(err)
	}

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	var assertion struct {
		Assertion string
	}
	json.Unmarshal(contents[1:], &assertion)

	value := fmt.Sprintf("%s,0,%s", loginData.Nickname, assertion.Assertion)
	botConnection.SendGlobalCommand("trn", value)

	botConnection.onSuccess <- struct{}{}

	for _, room := range loginData.Rooms {
		botConnection.SendGlobalCommand("join", room)
	}
}
func initializeChatRoom(rawMessage string, room *Room) {
	room.BotConnection.rooms[room.ID] = room
}

func setTitle(rawMessage string, room *Room) {
	room.Title = rawMessage
}

func setUsers(rawMessage string, room *Room) {
	room.onUserList(rawMessage)
}

func joinRoom(rawMessage string, room *Room) {
	room.onJoin(rawMessage)
}

func leaveRoom(rawMessage string, room *Room) {
	room.onLeave(rawMessage)
}

func renameNick(rawMessage string, room *Room) {
	parts := strings.SplitN(rawMessage, "|", 2)
	room.onRename(parts[0], UserID(parts[1]))
}
