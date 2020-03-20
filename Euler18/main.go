package main

import (
	"fmt"
	"strconv"
	"strings"
)

const data = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

const data2 = `3
7 4
2 4 6
8 5 9 3`

func main() {
	var t = [][]int{}
	fmt.Println("Maximum Path Sum")
	t1 := strings.Split(data, "\n")
	for x := range t1 {
		fmt.Println(t1[x])
		t2 := strings.Split(t1[x], " ")
		t3 := []int{}
		for n := range t2 {
			v, _ := strconv.Atoi(t2[n])
			t3 = append(t3, v)
		}
		t = append(t, t3)
	}
	for x := range t {
		for y := range t[x] {
			fmt.Printf("%2d ", t[x][y])
		}
		fmt.Println()
	}
	fmt.Println(max(t, 0, 0))
}

func max(t [][]int, x, y int) int {
	var sum = 0
	sum += t[x][y]
	x++
	if x == len(t) {
		return sum
	}
	left := max(t, x, y)
	right := max(t, x, y+1)
	if left > right {
		sum += left
	} else {
		sum += right
	}

	return sum
}