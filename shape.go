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

type shape struct {
	color          termbox.Attribute
	points         []point
	position       []point
	centerPosition []int
	xOffset        int
	yOffset        int
}

func (s *shape) rotate(isLeft bool) {

	s.points = transpose(s.points)

	if isLeft {
		reverseColumns(s.points)
	} else {
		reverseRows(s.points)
	}

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
	var currCoords []int
	position := make([]point, 0, 4)

loopyMcLoopface:
	for rowNum, row := range s.points {

		for colNum, col := range row {

			if col > 0 {

				currCoords = []int{colNum + colNum + s.xOffset, rowNum + s.yOffset}
				position = append(position, currCoords)

				if col == 3 { // check if current block is the center/pivot block

					if s.centerPosition == nil {
						s.centerPosition = currCoords // one-time initial setting
					} else if !(s.centerPosition[0] == currCoords[0] && s.centerPosition[1] == currCoords[1]) {
						// shape rotation caused the center/pivot block to move out of position
						// fix offsets and start over!
						//fmt.Print("s.centerPosition: " s.centerPosition[1])
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
	points: []point{
		{1, 3, 1, 1}}}

var shapeJ = shape{
	color: termbox.ColorYellow,
	points: []point{
		{1, 3, 1},
		{0, 0, 1}}}

var shapeL = shape{
	color: termbox.ColorCyan,
	points: []point{
		{1, 3, 1},
		{1, 0, 0}}}

var shapeO = shape{
	color: termbox.ColorMagenta,
	points: []point{
		{1, 1},
		{1, 1}}}

var shapeS = shape{
	color: termbox.ColorGreen,
	points: []point{
		{0, 3, 1},
		{1, 1, 0}}}

var shapeT = shape{
	color: termbox.ColorWhite,
	points: []point{
		{1, 3, 1},
		{0, 1, 0}}}

var shapeZ = shape{
	color: termbox.ColorRed,
	points: []point{
		{1, 3, 0},
		{0, 1, 1}}}
