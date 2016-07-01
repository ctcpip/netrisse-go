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
	color    termbox.Attribute
	points   []point
	position []point
	xOffset  int
	yOffset  int
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
		s.yOffset++
	}

}

func (s *shape) setPosition() {

	position := make([]point, 0, 4)

	for rowNum, row := range s.points {

		for colNum, col := range row {

			if col > 0 {
				position = append(position, []int{colNum + colNum + s.xOffset, rowNum + s.yOffset})
			}

		}

	}

	s.position = position

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
		{1, 1, 1, 1}}}

var shapeJ = shape{
	color: termbox.ColorYellow,
	points: []point{
		{1, 1, 1},
		{0, 0, 1}}}

var shapeL = shape{
	color: termbox.ColorCyan,
	points: []point{
		{1, 1, 1},
		{1, 0, 0}}}

var shapeO = shape{
	color: termbox.ColorMagenta,
	points: []point{
		{1, 1},
		{1, 1}}}

var shapeS = shape{
	color: termbox.ColorGreen,
	points: []point{
		{0, 1, 1},
		{1, 1, 0}}}

var shapeT = shape{
	color: termbox.ColorWhite,
	points: []point{
		{1, 1, 1},
		{0, 1, 0}}}

var shapeZ = shape{
	color: termbox.ColorRed,
	points: []point{
		{1, 1, 0},
		{0, 1, 1}}}
