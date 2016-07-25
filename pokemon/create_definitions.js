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

"use strict"

const {writeFile} = require('fs')

const dataDirectory = './../vendor/github.com/Zarel/Pokemon-Showdown/data/'

function toId(str) {
	return str.toLowerCase().replace(/[^a-z0-9]/g, "")
}

function values(obj) {
	const values = []
	for (const key in obj) {
		values.push(obj[key])
	}
	return values
}

function loadFile(name, key) {
	return values(require(dataDirectory + name)[key])
}

function upperFirst(str) {
	return str.charAt(0).toUpperCase() + str.slice(1)
}

function J(obj) {
	if (typeof obj !== 'object') {
		return JSON.stringify(obj)
	}
	let output = '{'
	for (const key in obj) {
		if (!obj[key]) continue
		output += upperFirst(key) + ': '
		output += J(obj[key])
		output += ', '
	}
	return output.slice(0, -2) + '}'
}

const pokedex = loadFile('pokedex', 'BattlePokedex')
const formatsData = require(dataDirectory + 'formats-data').BattleFormatsData
const moves = loadFile('moves', 'BattleMovedex')
const abilities = loadFile('abilities', 'BattleAbilities')

const header = `// Copyright (c) 2011-2016 Guangcong Luo and other contributors
// http://pokemonshowdown.com/
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package pokemon

import "github.com/xfix/showdown2irc/showdown"

`

function convertAbilities(abilities) {
	const result = []
	for (const possibleAbility of [0, 1, 'H']) {
		let ability = abilities[possibleAbility]
		if (ability) {
			let prefix = possibleAbility === 'H' ? 'Hidden: ' : ""
			result.push(`${prefix}${J(ability)}`)
		}
	}
	return `[]string{${result.join(", ")}}`
}

function parsePokedex(pokedex) {
	let output = header
	output += "var pokemon = map[showdown.UserID]*Pokemon{\n"
	for (const {species, types, abilities, baseStats, baseSpecies} of pokedex) {
		output += `\t${J(toId(species))}: {
\t\tSpecies:   ${J(species)},
\t\tTier:      ${J(formatsData[toId(baseSpecies || species)].tier)},
\t\tTypes:     []Type{${types.join(", ")}},
\t\tAbilities: ${convertAbilities(abilities)},
\t\tBaseStats: Stats${J(baseStats)},
\t},
`
	}
	output += "}\n"
	return output
}

function parseMoves(moves) {
	let output = header
	output += "var moves = map[showdown.UserID]*Move{\n"
	for (const {name, type, category, basePower, accuracy, pp, shortDesc} of moves) {
		output += `\t${J(toId(name))}: {
\t\tName:           ${J(name)},
\t\tType:           ${type},
\t\tDamageCategory: ${category},
`
		if (basePower) {
			output += `\t\tBasePower:      ${basePower},\n`
		}
		if (typeof accuracy === 'number') {
			output += `\t\tAccuracy:       ${accuracy},\n`
		}
		output += `\t\tPP:             ${pp},
\t\tDescription:    ${J(shortDesc)},
\t},
`
	}
	output += "}\n"
	return output
}

function parseAbilities(abilities) {
	let output = header
	output += "var abilities = map[showdown.UserID]*Ability{\n"
	for (const {name, shortDesc} of abilities) {
		output += `\t${J(toId(name))}: {
\t\tName:        ${J(name)},
\t\tDescription: ${J(shortDesc)},
\t},
`
	}
	output += "}\n"
	return output
}

writeFile("pokemon.go", parsePokedex(pokedex))
writeFile("moves.go", parseMoves(moves))
writeFile("abilities.go", parseAbilities(abilities))
