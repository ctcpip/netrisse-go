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
	"math/rand"
	"time"
)

var s shape

type game struct {
	interval time.Duration
	board    board
	timer    *time.Timer
}

func (g *game) start() {

	getNewShape := true

	g.board = board{top: 2, right: 21, bottom: 23, left: 0}
	g.board.draw()

	rand.Seed(time.Now().Unix())

	if g.interval <= 0 {
		g.interval = time.Duration(int(.5*1000)) * time.Millisecond
	}

	for {

		g.timer = time.NewTimer(g.interval)
		<-g.timer.C

		if getNewShape {
			getNewShape = false
			s = shapes[rand.Intn(6)]
			s.board = &g.board
			s.xOffset = g.board.left + initialXOffset
			s.yOffset = g.board.top - 1
			s.setPosition()
		}

		if !s.move(DOWN) {
			getNewShape = true
		}

	}

}
