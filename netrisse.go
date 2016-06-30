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
import "fmt"

func main() {

	var s shape

	load()

	var gc gameContainer
	gc.draw()

	s = shapeI
	s.draw()

	// s = shapeJ
	// s.rotate(true)
	// s.draw()N

	// ---------- BEGIN MATRIX TRANSFORMATION INFO

	// printMatrix(s.points)
	// s.rotate(true)
	// fmt.Println("")
	// printMatrix(s.points)
	//
	// s.rotate(false)
	// fmt.Println("")
	// printMatrix(s.points)
	// s.rotate(false)
	// fmt.Println("")
	// printMatrix(s.points)

	// ---------- END MATRIX TRANSFORMATION INFO

	readKey()

	close()

}

func printMatrix(points []point) {

	for _, p := range points {
		fmt.Println(p)
	}

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

func reverseColumns(points []point) {

	var intCurr int

	for col := 0; col < len(points[0]); col++ {

		for row := 0; row < len(points)/2; row++ {

			intCurr = points[row][col]
			points[row][col] = points[len(points)-row-1][col]
			points[len(points)-row-1][col] = intCurr

		}

	}

}

func (s *shape) rotate(isLeft bool) {

	s.points = transpose(s.points)

	if isLeft {
		reverseColumns(s.points)
	} else {
		reverseRows(s.points)
	}

}

func (s *shape) draw() {

	var currOGX int

	for rowNum, row := range s.points {

		for colNum, col := range row {

			if col > 0 {

				currOGX = colNum

				if rowNum > 0 {
					termbox.SetCell(colNum+currOGX+containerXOffset, containerYOffset, '[', termbox.ColorBlack, s.color)
					termbox.SetCell(colNum+currOGX+containerXOffset+1, containerYOffset, ']', termbox.ColorBlack, s.color)
				}

			}

		}

	}

	termbox.Flush()

}

var shapeI = shape{
	termbox.ColorBlue,
	[]point{
		{0, 0, 0, 0},
		{1, 1, 1, 1}}}

var shapeJ = shape{
	termbox.ColorYellow,
	[]point{
		{1, 1, 1},
		{0, 0, 1}}}

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
