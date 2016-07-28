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

package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xfix/showdown2irc/showdown"
)

func TestGetMessage(t *testing.T) {
	input := showdown.ToID("Bulbasaur")
	assert.Equal(t, GetPokemon(input).Types, []Type{Grass, Poison}, "GetPokemon(%#q)", input)
}
