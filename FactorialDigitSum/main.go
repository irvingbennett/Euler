package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var sum int

	n := big.NewInt(1)
	m := big.NewInt(2)
	fmt.Println("Factorial Digit Sum")
	for x := 0; x < 99; x++ {
		n = n.Mul(n, m)
		m.Add(m, big.NewInt(1))
	}

	fmt.Println(n.Text(10))
	digits := n.Text(10)
	for x := 0; x < len(digits); x++ {
		// fmt.Printf("%s", string(digits[x]))
		i, err := strconv.Atoi(string(digits[x]))
		if err != nil {
			fmt.Println(err)
		}
		sum += i
	}
	fmt.Println(sum)
}
