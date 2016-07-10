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
	"sort"

	"github.com/nsf/termbox-go"
)

type shape struct {
	board          *board
	color          termbox.Attribute
	shapePoints    points
	position       points
	centerPosition point
	xOffset        int
	yOffset        int
	toggle         bool
	movable        bool
}

func (s *shape) rotate(isLeft bool) bool {

	booContinue := true

	switch s.color {
	case termbox.ColorMagenta: // O shape cannot rotate
		booContinue = false
	case termbox.ColorBlue: // shape I can only toggle betwixt 2 rotated views

		isLeft = s.toggle
		s.toggle = !s.toggle

	case termbox.ColorGreen, termbox.ColorRed: // shapes S and Z can only toggle betwixt 2 rotated views, and the center/pivot block changes

		if s.toggle {
			isLeft = s.color == termbox.ColorRed // shapes S and Z perform opposite rotations
			s.shapePoints[1][0] = 1
			s.shapePoints[1][1] = 3
		} else {
			isLeft = s.color != termbox.ColorRed
			s.shapePoints[0][1] = 3
			s.shapePoints[1][1] = 1
		}

		s.toggle = !s.toggle

	}

	if booContinue {

		s.shapePoints = transpose(s.shapePoints)

		if isLeft {
			reverseColumns(s.shapePoints)
		} else {
			reverseRows(s.shapePoints)
		}

	}

	return booContinue

}

func (s *shape) move(d direction) bool {

	booContinue, booLock := true, false
	sNew := copyShape(s)

	switch d {

	case DROP:

	case DOWN:

		destinationY := s.position.maxY() + 1

		if destinationY < s.board.bottom {
			sNew.yOffset = destinationY
			sNew.centerPosition[1]++
		} else {
			booContinue = false
			booLock = true
		}

	case LEFT:

		destinationX := s.position.minX() - 1

		if destinationX > s.board.left+1 {
			sNew.xOffset = destinationX - 1
			sNew.centerPosition[0] = sNew.centerPosition[0] - 2
		} else {
			booContinue = false
		}

	case RIGHT:

		destinationX := s.position.maxX() + 1

		if destinationX < s.board.right-2 {
			sNew.xOffset = destinationX + 1
			sNew.centerPosition[0] = sNew.centerPosition[0] + 2
		} else {
			booContinue = false
		}

	case ROTATE:

		if sNew.rotate(true) {

			sNew.setPosition()

			if sNew.position.maxX() >= s.board.right-2 ||
				sNew.position.minX() <= s.board.left+1 ||
				sNew.position.maxY() >= s.board.bottom {
				booContinue = false
			}

		} else {
			booContinue = false
		}

	}

	if booContinue {

		if d != ROTATE {
			// ROTATE already called setPosition() to check against board boundaries
			sNew.setPosition()
		}

	loopyMcLoopface:
		for _, bp := range s.board.occupied {

			for _, sp := range sNew.position {

				if bp[0] == sp[0] && bp[1] == sp[1] {
					booContinue = false
					booLock = true
					break loopyMcLoopface
				}

			}

		}

	}

	if booContinue {

		s.erase()
		*s = sNew
		s.draw()
		termbox.Flush()

	} else if booLock {
		s.movable = false
		s.board.occupied = append(s.board.occupied, s.position...)
		g.timer.Reset(0)
	}

	return !booLock

}

func (s *shape) setPosition() {

	var booStartOver bool
	var currCoords point
	position := make(points, 0, 4)

loopyMcLoopface:
	for rowNum, row := range s.shapePoints {

		for colNum, col := range row {

			if col > 0 {

				currCoords = point{colNum + s.xOffset, rowNum + s.yOffset}
				position = append(position, currCoords)

				if col == 3 { // check if current block is the center/pivot block

					if len(s.centerPosition) == 0 {
						s.centerPosition = currCoords // one-time initial setting
					} else if !(s.centerPosition[0] == currCoords[0] && s.centerPosition[1] == currCoords[1]) {
						// shape rotation caused the center/pivot block to move out of position
						// fix offsets and start over!

						if s.centerPosition[1] < currCoords[1] {
							s.yOffset -= currCoords[1] - s.centerPosition[1]
						} else {
							s.yOffset += s.centerPosition[1] - currCoords[1]
						}

						if s.centerPosition[0] < currCoords[0] {
							s.xOffset -= currCoords[0] - s.centerPosition[0]
						} else {
							s.xOffset += s.centerPosition[0] - currCoords[0]
						}

						booStartOver = true
						break loopyMcLoopface

					}
				}
			}
		}
	}

	if booStartOver {
		s.setPosition()
	} else {
		sort.Sort(position)
		s.position = position
	}

}

func (s *shape) erase() { drawShape(s, true) }
func (s *shape) draw()  { drawShape(s, false) }

func drawShape(s *shape, erase bool) {

	var fg, bg termbox.Attribute
	var r1, r2 rune
	var x, y, centerX, currX int

	centerX = s.centerPosition[0]

	if erase {
		fg, bg = termbox.ColorDefault, termbox.ColorDefault
		r1, r2 = ' ', ' '
	} else {
		fg, bg = termbox.ColorBlack, s.color
		r1, r2 = '[', ']'
	}

	// if !erase {
	// 	logger.Print(s.position.toString())
	// }

	for _, p := range s.position {

		if p[1] > s.board.top {

			// if !erase {
			// 	logger.Print("col adjust: " + strconv.Itoa(p[0]) + " = " + strconv.Itoa(cols[p[0]]))
			// 	logger.Print("drawing:    " + strconv.Itoa(p[0]+cols[p[0]]) + " & " + strconv.Itoa(p[0]+cols[p[0]]+1))
			// }

			currX = p[0]
			x = currX - (centerX - currX)
			y = p[1]

			termbox.SetCell(x, y, r1, fg, bg)
			termbox.SetCell(x+1, y, r2, fg, bg)

		}

	}

}
