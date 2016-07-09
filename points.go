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

import "bytes"

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

	if !isMax {
		maxX = 33 // set initial value for minX, otherwise will never be less than 0
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
