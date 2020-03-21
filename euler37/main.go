//Package prime returns the nth prime
package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jbarham/primegen.go"
)

func main() {
	fmt.Println("Truncatable Primes")
	truncatable := []int{}
	p := primegen.New()
	var next int
	sum := 0
	for {
		next = int(p.Next())
		if next < 1000000 {
			if isTruncatable(next) {
				truncatable = append(truncatable, next)
			}
		} else {
			break
		}
	}
	// fmt.Println("2,000,000 prime is", nth)
	for x := range truncatable {
		fmt.Println(x, truncatable[x])
		sum += truncatable[x]
	}
	fmt.Println("Sum is", sum)
}

// return true if a prime is truncatable
func isTruncatable(n int) bool {
	if n < 10 {
		return false
	}
	perms := []string{}
	s := fmt.Sprintf("%d", n)
	r := []rune(s)
	l := len(r)
	for x := 0; x < l; x++ {
		perms = append(perms, string(r))
		r = r[1:]
	}

	r = []rune(s)
	for x := 0; x < l; x++ {
		perms = append(perms, string(r))
		r = r[:len(r)-1]
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
