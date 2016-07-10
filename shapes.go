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

var shapes = []shape{shapeI, shapeJ, shapeL, shapeO, shapeS, shapeT, shapeZ}

var shapeI = shape{
	color: termbox.ColorBlue,
	shapePoints: points{
		{0, 0, 0, 0},
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
