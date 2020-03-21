package main

// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

// Find the sum of all numbers which are equal to the sum of the
// factorial of their digits.

// Note: as 1! = 1 and 2! = 2 are not sums they are not included.

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("Digit Factorial")
	numbers := []int{}
	sum := 0
	for x := 10; x < 999999; x++ {
		n := big.NewInt(int64(x))
		if digitFactorial(n) {
			numbers = append(numbers, x)
			sum += x
			fmt.Println(x)
		}
	}
	fmt.Println(sum)

}

func factorial(f int) *big.Int {
	n := big.NewInt(1)
	m := big.NewInt(2)

	for x := 0; x < f-1; x++ {
		n = n.Mul(n, m)
		m.Add(m, big.NewInt(1))
	}
	return n
}

func digitFactorial(n *big.Int) bool {
	var sum int64
	digits := n.Text(10)
	for x := 0; x < len(digits); x++ {
		// fmt.Printf("%s", string(digits[x]))
		i, err := strconv.Atoi(string(digits[x]))
		if err != nil {
			fmt.Println(err)
		}
		sum += factorial(i).Int64()
	}
	if int64(sum) == n.Int64() {
		return true
	}
	return false
}
