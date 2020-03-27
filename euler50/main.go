package main

import (
	"fmt"

	"github.com/kavehmz/prime"
)

func main() {
	fmt.Println("Prime Permutations")
	p := prime.Primes(1000000)
	mp := make(map[int]bool, 0)
	for x := range p {
		mp[int(p[x])] = true
	}
	most := 1
	x := 0
	y := 0
	l := len(mp)
	max := p[l-1]
	last := 0
	fmt.Printf("There are %d primes under a million\n", l)
	fmt.Println("The largest prime in the range is", max)

	for {
		sum := 0
		y = 0
		for y = 0; y < most && x+y < l && sum+int(p[x+y]) < int(max); y++ {
			sum += int(p[x+y])
		}
		if mp[sum] && y == most {
			last = most
			println(most, sum, x, y)
		}

		most++
		if most == l {
			x++
			most = last
		}
		if x == l {
			break
		}
	}
}
