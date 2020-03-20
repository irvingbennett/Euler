package main

import (
	"fmt"
)

var matrix [][]int
var size int

// board := make([]string, m*n)
//  board[i*m + j] = "abc" // like board[i][j] = "abc"

// int sizes[width*height];
// void setPoint(int x, int y, int val) {
//  sizes[x*width + y] = val;
// }

func main() {
	fmt.Println("Spiral Diagonals")
	fmt.Println("======================")
	fmt.Println()
	size = 5
	matrix = make([][]int, 5)
	for x := 0; x < size; x++ {
		line := make([]int, 5)
		for y := 0; y < size; y++ {
			line[y] = 0
		}
		matrix[x] = line
	}
	matrix[2][2] = 1
	for x := 0; x < size; x++ {
		fmt.Println(matrix[x])
	}

}
