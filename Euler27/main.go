package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
	fmt.Println("Quadratic primes")
	fmt.Println("====================")
	fmt.Println()

	maxPrimes := 0
	best := []int{1, 1}
	for a := -1000; a < 1000; a++ {
		for b := -999; b <= 1000; b++ {
			n := Primes(a, b)
			if n > maxPrimes {
				maxPrimes = n
				best[0] = a
				best[1] = b

				fmt.Print("Max ", maxPrimes, " -> ")
				fmt.Println("Best", best)
			}
		}
	}
	fmt.Println("Max", maxPrimes)
	fmt.Println("Best", best)
	fmt.Println()
	fmt.Println("Product", best[0]*best[1])
}

func qp(n, a, b int) (p int) {
	p = (n * n) + n*a + b
	return
}

// IsPrimeSqrt returns true if a value is prime
func IsPrimeSqrt(value int) bool {
	if value == 2 || value == 3 || value == 5 || value == 7 {
		return true
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// Primes returns the count of consecutive primes
func Primes(a, b int) (p int) {
	for n := 0; n < 1000; n++ {
		if IsPrimeSqrt(qp(n, a, b)) {
			p++
		} else {
			break
		}
	}
	return
}
