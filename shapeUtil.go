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
