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

type keyboard struct{ keyDrop, keyRight, keyDown, keyLeft, keyRotateLeft rune }

func (k *keyboard) read() {

loopyMcLoopface:
	for {

		switch e := termbox.PollEvent(); e.Type {

		case termbox.EventKey:

			switch {
			case e.Key == termbox.KeyCtrlC:
				break loopyMcLoopface
			case e.Ch == k.keyDrop:
				return
			case e.Ch == k.keyRight:
				s.move(RIGHT)
			case e.Ch == k.keyDown:
				s.move(DOWN)
				g.timer.Reset(g.interval)
			case e.Ch == k.keyLeft:
				s.move(LEFT)
			case e.Ch == k.keyRotateLeft:

				if s.rotate(true) {
					s.erase()
					s.setPosition()
					s.draw()
					termbox.Flush()
				}

				//g.timer.Reset(g.interval)

			}

		case termbox.EventError:
			panic(e.Err)
		}

	}

}
