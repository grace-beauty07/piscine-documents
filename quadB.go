package piscine

import "github.com/01-edu/z01"

// Quad(5,3)
// /***\
// *   *
// \***/

// Quad(5,1)
// /***\

// Quad(1,5)
// /
// *
// *
// *
// \

func QuadB(x int, y int) {
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 { // Top left corner
				z01.PrintRune('/')
			} else if row == 1 && col == x { // Top right corner
				z01.PrintRune('\\')
			} else if row == y && col == 1 { // Bottom left corner
				z01.PrintRune('\\')
			} else if row == y && col == x { // Bottom right corner
				z01.PrintRune('/')
			} else if (col > 1 || col < x) && (row == 1 || row == y) { // remaining top and bottom rows
				z01.PrintRune('*')
			} else if (row > 1 || row < y) && (col == 1 || col == x) { // remaining left and right columns
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ') // every other space
			}
		}
		z01.PrintRune('\n') // go to new line after every row print
	}
	z01.PrintRune('\n')
}
