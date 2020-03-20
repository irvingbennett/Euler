package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Digit fifth powers")
	power := 5
	total := 0
	for x := 2; x <= 999999; x++ {
		if check(x, power) {
			fmt.Println(x)
			total += x
		}
	}
	fmt.Println(total)
}

func check(n, p int) bool {
	if n == 1 {
		return false
	}
	s := fmt.Sprintf("%d", n)
	// fmt.Print(s, " -> ")
	l := len(s)
	sum := 0
	for x := 0; x < l; x++ {
		a := float64(int(s[x]) - int('0'))
		// fmt.Printf("%d", int(a))
		b := float64(p)
		v := int(math.Pow(a, b))
		sum += v
	}
	// fmt.Printf(" -> %d\n", sum)
	if sum == n {
		return true
	}
	return false
}
