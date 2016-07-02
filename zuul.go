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

import "fmt"

func printMatrix(points points) {

	for _, p := range points {
		fmt.Println(p)
	}

}

func testSomeStuff() {

	var s shape
	s = shapeJ
	s.rotate(true)
	s.draw()

	// ---------- BEGIN MATRIX TRANSFORMATION INFO

	printMatrix(s.shapePoints)
	s.rotate(true)
	fmt.Println("")
	printMatrix(s.shapePoints)

	s.rotate(false)
	fmt.Println("")
	printMatrix(s.shapePoints)
	s.rotate(false)
	fmt.Println("")
	printMatrix(s.shapePoints)

	// ---------- END MATRIX TRANSFORMATION INFO

}
