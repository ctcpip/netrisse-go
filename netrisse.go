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

func main() {

	var s shape

	load()

	var gc gameContainer
	gc.draw()

	s = shapeI
	//s.draw()

	s = shapeJ
	//s.rotate(false)
	s.draw()

	readKey()

	close()

}

func readKey() {

loopyMcLoopface:
	for {

		switch e := termbox.PollEvent(); e.Type {

		case termbox.EventKey:

			if e.Key == termbox.KeyCtrlC {
				break loopyMcLoopface
			}

		case termbox.EventError:
			panic(e.Err)
		}

	}

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

// type point struct {
// 	x, y int
// }

type point []int

type shape struct {
	color  termbox.Attribute
	points []point
}

const containerXOffset = 9
const containerYOffset = 5 // s/b 2 - using 5 for testing

func transpose(points []point) []point {

	p := make([]point, len(points[0]))

	for x := range p {
		p[x] = make(point, len(points))
	}

	for y, a := range points {
		for x, b := range a {
			p[x][y] = b
		}
	}

	return p

}

func reverseRows(points []point) {

	for _, a := range points {

		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}

	}

}

func (s *shape) rotate(isLeft bool) {

	if isLeft {
		//s.points = reverseColumns(transpose(s.points))
	} else {
		s.points = transpose(s.points)
		//reverseRows(s.points)
	}

}

func (s *shape) draw() {

	var currOGX int

	for _, p := range s.points {

		currOGX = p[0]
		p[0] += containerXOffset
		p[1] += containerYOffset

		if p[1] > 2 {
			termbox.SetCell(p[0]+currOGX, p[1], '[', termbox.ColorBlack, s.color)
			termbox.SetCell(p[0]+currOGX+1, p[1], ']', termbox.ColorBlack, s.color)
		}

	}

	termbox.Flush()

}

var shapeI = shape{
	termbox.ColorBlue,
	[]point{
		{0, 1},
		{1, 1},
		{2, 1},
		{3, 1}}}

var shapeJ = shape{
	termbox.ColorYellow,
	[]point{
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1}}}

// var shapeL = shape{[]point{
// 	{0, 0},
// 	{1, 0},
// 	{2, 0},
// 	{0, 1}}}
//
// var shapeO = shape{[]point{
// 	{0, 0},
// 	{0, 1},
// 	{1, 0},
// 	{1, 1}}}
//
// var shapeS = shape{[]point{
// 	{0, 0},
// 	{1, 0},
// 	{2, 0},
// 	{3, 0}}}
//
// var shapeT = shape{[]point{
// 	{0, 0},
// 	{1, 0},
// 	{2, 0},
// 	{3, 0}}}
//
// var shapeZ = shape{[]point{
// 	{0, 0},
// 	{1, 0},
// 	{2, 0},
// 	{3, 0}}}

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
