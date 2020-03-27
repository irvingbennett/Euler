package main

import (
	"fmt"
	"math"
)

var primes, squaresX2 []int

func main() {
	fmt.Println("Composite Numbers")
	primes = Sieve(1000000)
	squaresX2 = []int{}
	for x := 1; x < 1000000; x++ {
		squaresX2 = append(squaresX2, x*x*2)
	}
	fmt.Println(len(primes))
	n := 0
	for x := range primes {
		fmt.Println(primes[x], squaresX2[n])
		n++
		if x > 10 {
			break
		}
	}
	p := 3
	fmt.Println(primes[p])

	for x := 9; x < 999999999; x += 2 {
		//fmt.Println(x, "fits Goldbach's is", check(n))
		if !check(x) {
			fmt.Println(x, "fits Goldbach's is false")
			break
		}
	}

}

func check(n int) bool {
	for x := range primes {
		if primes[x] == n {
			return true
		}
		if primes[x] > n {
			for z := x - 1; z > 0; z-- {
				for y := range squaresX2 {
					if n-(primes[z]+squaresX2[y]) == 0 {
						fmt.Println(n, "fits Goldbach's is true", primes[z], squaresX2[y])
						return true
					}
					if (primes[z] + squaresX2[y]) > n {
						break
					}
				}
			}
		}
	}
	return false
}

// Package sieve does...

// Sieve returns all the primes from 2 up to a given number.
func Sieve(n int) (primes []int) {
	primes = make([]int, n+16)
	// fmt.Println(n)
	iter := primesOdds(uint(n + 1))
	x := 0
	for v := iter(); v <= uint(n); v = iter() {
		p := int(v)
		// fmt.Printf("%d, ", p)
		if p > 0 {
			primes[x] = p
			x++
		}
	}
	primes = primes[:x]
	return
}

func primesOdds(top uint) func() uint {
	topndx := int((top - 3) / 2)
	topsqrtndx := (int(math.Sqrt(float64(top))) - 3) / 2
	cmpsts := make([]uint, top)
	// fmt.Println(top, topndx, topsqrtndx, (topndx/32)+1)
	for i := 0; i <= topsqrtndx; i++ {
		if cmpsts[i>>5]&(uint(1)<<(uint(i)&0x1F)) == 0 {
			p := (i << 1) + 3
			for j := (p*p - 3) >> 1; j <= topndx; j += p {
				cmpsts[j>>5] |= 1 << (uint(j) & 0x1F)
			}
		}
	}
	i := -1
	return func() uint {
		oi := i
		if i <= topndx {
			i++
		}
		for i <= topndx && cmpsts[i>>5]&(1<<(uint(i)&0x1F)) != 0 {
			i++
		}
		if oi < 0 {
			return 2
		} else {
			return (uint(oi) << 1) + 3
		}
	}
}
