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
	"sort"

	"github.com/nsf/termbox-go"
)

type points []point

func (pts points) maxY() int {
	return getLimit(pts, true, false)
}

func (pts points) maxX() int {
	return getLimit(pts, true, true)
}

func (pts points) minX() int {
	return getLimit(pts, false, true)
}

func getLimit(pts points, isMax bool, isX bool) int {

	var currX, maxX, i int

	if isX {
		i = 0
	} else {
		i = 1
	}

	for _, p := range pts {

		currX = p[i]

		if isMax {

			if currX > maxX {
				maxX = currX
			}

		} else {

			if currX < maxX {
				maxX = currX
			}

		}

	}

	return maxX

}

func (pts points) Len() int {
	return len(pts)
}

func (pts points) Less(i, j int) bool {
	return pts[i][0] < pts[j][0] // points are sorted on x coord
}

func (pts points) Swap(i, j int) {
	pts[i], pts[j] = pts[j], pts[i]
}

func (pts points) toString() string {

	var b bytes.Buffer

	b.WriteString("\n")

	for i, p := range pts {

		b.WriteString(p.toString())

		if i < len(pts)-1 {
			b.WriteString("\n")
		}

	}

	return b.String()

}

type shape struct {
	board          *board
	color          termbox.Attribute
	shapePoints    points
	position       points
	centerPosition point
	xOffset        int
	yOffset        int
	toggle         bool
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

		//logger.Print(s.shapePoints.toString())

		if s.toggle {
			isLeft = s.color == termbox.ColorRed // shapes S and Z perform opposite rotations
			s.shapePoints[1][0] = 1
			s.shapePoints[1][1] = 3
		} else {
			isLeft = s.color != termbox.ColorRed
			s.shapePoints[0][1] = 3
			s.shapePoints[1][1] = 1
		}

		//logger.Print(s.shapePoints.toString())

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

type direction int

const (
	// RIGHT move shape right
	RIGHT direction = iota
	//DOWN move shape down
	DOWN
	// LEFT move shape left
	LEFT
)

func (s *shape) move(d direction) bool {

	sNew := *s
	booContinue := true

	switch d {
	case DOWN:

		destinationY := s.position.maxY() + 1

		if destinationY < s.board.bottom {
			sNew.yOffset = destinationY
			sNew.centerPosition[1]++
		} else {
			booContinue = false
		}

	case LEFT:

	case RIGHT:

		destinationX := s.position.maxX() + 2

		if destinationX < s.board.right-2 {
			sNew.xOffset = destinationX
			sNew.centerPosition[0] = sNew.centerPosition[0] + 2
		}

	}

	if booContinue {

		sNew.setPosition()

	loopyMcLoopface:
		for _, bp := range s.board.occupied {

			for _, sp := range sNew.position {

				if bp[0] == sp[0] && bp[1] == sp[1] {
					booContinue = false
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
	} else {
		s.board.occupied = append(s.board.occupied, s.position...)
	}

	return booContinue

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

					if s.centerPosition == nil {
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
					//logger.Print("center: " + strconv.Itoa(currCoords[0]) + " , " + strconv.Itoa(currCoords[1]))
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

var shapes = []shape{shapeI, shapeJ, shapeL, shapeO, shapeS, shapeT, shapeZ}

var shapeI = shape{
	color: termbox.ColorBlue,
	shapePoints: points{
		{1, 3, 1, 1}}}

var shapeJ = shape{
	color: termbox.ColorYellow,
	shapePoints: points{
		{1, 3, 1},
		{0, 0, 1}}}

var shapeL = shape{
	color: termbox.ColorCyan,
	shapePoints: points{
		{1, 3, 1},
		{1, 0, 0}}}

var shapeO = shape{
	color: termbox.ColorMagenta,
	shapePoints: points{
		{1, 3},
		{1, 1}}}

var shapeS = shape{
	color: termbox.ColorGreen,
	shapePoints: points{
		{0, 1, 1},
		{1, 3, 0}}}

var shapeT = shape{
	color: termbox.ColorWhite,
	shapePoints: points{
		{1, 3, 1},
		{0, 1, 0}}}

var shapeZ = shape{
	color: termbox.ColorRed,
	shapePoints: points{
		{1, 1, 0},
		{0, 3, 1}}}
