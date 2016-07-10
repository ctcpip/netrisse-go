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

import "time"

type game struct {
	interval time.Duration
	board    board
	timer    *time.Timer
}

func (g *game) start() {

	booGetNewShape := true

	g.board = board{top: 2, right: 21, bottom: 23, left: 0}
	g.board.draw()

	if g.interval <= 0 {
		g.interval = time.Duration(int(.5*1000)) * time.Millisecond
	}

	for {

		if booGetNewShape {
			booGetNewShape = false
			getNewShape()
		}

		if !s.move(DOWN) {
			booGetNewShape = true
		}

		g.timer = time.NewTimer(g.interval)
		<-g.timer.C

	}

}
