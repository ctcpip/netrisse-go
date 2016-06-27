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
	"time"

	"github.com/nsf/termbox-go"
)

func main() {

	load()

	var gc gameContainer
	gc.draw()

	time.Sleep(3 * time.Second)

	close()

}

func load() {

	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	writeText("netrisse 0.1.0 (C) 2016  Chris de Almeida     \"netrisse -h\" for more info", 0, 0)
	termbox.Flush()

}

type gameContainer struct{}

func (c *gameContainer) draw() {

	//top
	writeText("+--------------------+", 0, 2)

	//left
	for i := 3; i < 23; i++ {
		termbox.SetCell(0, i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	//right
	for i := 3; i < 23; i++ {
		termbox.SetCell(21, i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	//bottom
	writeText("+--------------------+", 0, 23)

	termbox.Flush()

}

func writeText(text string, startX, y int) {

	currX := startX

	for i := 0; i < len(text); i++ {
		termbox.SetCell(currX, y, rune(text[i]), termbox.ColorDefault, termbox.ColorDefault)
		currX++
	}

}

func close() { defer termbox.Close() }
