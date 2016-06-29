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
	s.draw()

	s = shapeJ
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

type point struct {
	x, y int
	char rune
}

type shape struct {
	color  termbox.Attribute
	points []point
}

const containerXOffset = 9
const containerYOffset = 3

func (s *shape) rotate(isLeft bool) {

	if isLeft {
	}

	// for _, p := range s.points {
	//
	// 	termbox.SetCell(p.x+containerXOffset, p.y+containerYOffset, p.char, termbox.ColorBlack, s.color)
	//
	// }

}

func (s *shape) draw() {

	var currContainerYOffset int
	isIShape := s.color == termbox.ColorBlue // hacky but avoids the need for another property or having to do additional processing

	if isIShape {
		currContainerYOffset = containerYOffset
	} else {
		currContainerYOffset = containerYOffset - 1
	}

	for _, p := range s.points {

		p.x = p.x + containerXOffset
		p.y = p.y + currContainerYOffset

		if p.y > 2 {
			termbox.SetCell(p.x, p.y, p.char, termbox.ColorBlack, s.color)
		}

	}

	termbox.Flush()

}

var shapeI = shape{
	termbox.ColorBlue,
	[]point{
		{0, 0, '['},
		{1, 0, ']'},
		{2, 0, '['},
		{3, 0, ']'},
		{4, 0, '['},
		{5, 0, ']'},
		{6, 0, '['},
		{7, 0, ']'}}}

var shapeJ = shape{
	termbox.ColorYellow,
	[]point{
		{0, 0, '['},
		{1, 0, ']'},
		{2, 0, '['},
		{3, 0, ']'},
		{4, 0, '['},
		{5, 0, ']'},
		{4, 1, '['},
		{5, 1, ']'}}}

// var shapeJ = shape{[]point{
// 	{0, 0},
// 	{1, 0},
// 	{2, 0},
// 	{2, 1}}}
//
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
