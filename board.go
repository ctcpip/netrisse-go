/*

netrisse - a network version of tetris for the console/terminal
Copyright (C) 2016  Chris de Almeida

http://github.com/ctcpip/netrisse

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

package main

import (
	"bytes"
	"strconv"

	"github.com/nsf/termbox-go"
)

const initialXOffset = 9

type point []int

func (p point) toString() string {

	var b bytes.Buffer

	b.WriteString("{ ")

	for _, x := range p {
		b.WriteString(strconv.Itoa(x))
		b.WriteString(" ")
	}

	b.WriteString("}")

	return b.String()

}

type board struct {
	top, right, bottom, left int
}

func (b *board) draw() {

	//top
	writeText("+--------------------+", b.left, b.top)

	//left
	for i := 3; i < 23; i++ {
		termbox.SetCell(0, i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	//right
	for i := 3; i < 23; i++ {
		termbox.SetCell(21, i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	//bottom
	writeText("+--------------------+", b.left, b.bottom)

	termbox.Flush()

}
