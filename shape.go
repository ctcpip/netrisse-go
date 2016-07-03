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

	"github.com/nsf/termbox-go"
)

type points []point

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

		logger.Print(s.shapePoints.toString())

		if s.toggle {
			isLeft = s.color == termbox.ColorRed // shapes S and Z perform opposite rotations
			s.shapePoints[1][0] = 1
			s.shapePoints[1][1] = 3
		} else {
			isLeft = s.color != termbox.ColorRed
			s.shapePoints[0][1] = 3
			s.shapePoints[1][1] = 1
		}

		logger.Print(s.shapePoints.toString())

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

func (s *shape) move() {

	for _, p := range s.position {
		p[1]++
	}

	s.yOffset++
	s.centerPosition[1]++

}

func (s *shape) setPosition() {

	var booStartOver bool
	var currCoords point
	position := make(points, 0, 4)

loopyMcLoopface:
	for rowNum, row := range s.shapePoints {

		for colNum, col := range row {

			if col > 0 {

				currCoords = point{colNum + colNum + s.xOffset, rowNum + s.yOffset}
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

				}

			}

		}

	}

	if booStartOver {
		s.setPosition()
	} else {
		s.position = position
	}

}

func (s *shape) erase() {

	for _, p := range s.position {

		if p[1] > 2 {
			termbox.SetCell(p[0], p[1], ' ', termbox.ColorDefault, termbox.ColorDefault)
			termbox.SetCell(p[0]+1, p[1], ' ', termbox.ColorDefault, termbox.ColorDefault)
		}

	}

	termbox.Flush()

}

func (s *shape) draw() {

	for _, p := range s.position {

		if p[1] > 2 {
			termbox.SetCell(p[0], p[1], '[', termbox.ColorBlack, s.color)
			termbox.SetCell(p[0]+1, p[1], ']', termbox.ColorBlack, s.color)
		}

	}

	termbox.Flush()

}

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
		{1, 1},
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
