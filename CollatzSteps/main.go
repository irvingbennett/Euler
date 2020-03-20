package main

import (
	"fmt"
)

// CollatzConjecture (input int) (steps int, err error)
func CollatzConjecture(input int) (steps int, err error) {
	if input == 0 || input < 0 {
		err = fmt.Errorf("improper input")
		return
	}

	var n = float64(input)
	steps = 0
	// fmt.Println(input)
	for n != 1 {
		if int(n)%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}

	return steps, nil
}

func main() {
	var longest, number int

	for x := 10; x < 1000000; x++ {
		if x%100 == 0 {
			fmt.Print(x / 100)
		}
		if x%1000 == 0 {
			fmt.Println()
			fmt.Println(x / 100)
		}
		steps, err := CollatzConjecture(x)
		if err != nil {
			fmt.Println(err)
		}
		if steps > longest {
			longest = steps
			number = x
		}
	}
	fmt.Println("Number:", number, "Steps:", longest)
}
