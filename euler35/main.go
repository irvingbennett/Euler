//Package prime returns the nth prime
package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jbarham/primegen.go"
)

func main() {
	fmt.Println("Circular Primes")
	circular := []int{}
	p := primegen.New()
	var next int
	for {
		next = int(p.Next())
		if next < 1000000 {
			if isCircular(next) {
				circular = append(circular, next)
			}
		} else {
			break
		}
	}
	// fmt.Println("2,000,000 prime is", nth)
	for x := range circular {
		fmt.Println(x, circular[x])
	}
}

// return true if a prime is circular
func isCircular(n int) bool {
	if n < 10 {
		return true
	}
	perms := []string{}
	s := fmt.Sprintf("%d", n)
	r := []rune(s)
	for x := 0; x < len(r); x++ {
		perms = append(perms, string(r))
		r = append(r[1:], r[0])
	}
	// fmt.Println(perms)
	for x := 0; x < len(perms); x++ {
		p, _ := strconv.Atoi(perms[x])
		if !IsPrimeSqrt(p) {
			return false
		}
	}
	return true
}

// IsPrimeSqrt return true if a number is prime
func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
