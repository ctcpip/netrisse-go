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

import "github.com/nsf/termbox-go"

type board struct {
	top, right, bottom, left int
	occupied                 points
	border                   dimensions
}

func (b *board) draw() {

	//top
	scr.writeText("+--------------------+", b.border.left, b.border.top)

	//left
	for i := b.border.top + 1; i < b.border.bottom; i++ {
		termbox.SetCell(b.border.left, i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	//right
	for i := b.border.top + 1; i < b.border.bottom; i++ {
		termbox.SetCell(b.border.right, i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	//bottom
	scr.writeText("+--------------------+", b.border.left, b.border.bottom)

	termbox.Flush()

}
