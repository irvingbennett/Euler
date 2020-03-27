package main

import (
	"fmt"

	"github.com/kavehmz/prime"
)

func main() {
	fmt.Println("Prime Digit Replacements")
	p := prime.Primes(1000000)
	x := 0
	for x = range p {
		if p[x] > 100000 {
			break
		}
	}
	p = p[x:]
	x = 0
	/*
		for x = range p {
			if p[x] > 199993 {
				break
			}
		}
		p = p[:x]
	*/
	mp := make(map[int]bool, 0)
	for x := range p {
		s := fmt.Sprintf("%d", p[x])
		// if string(s[0]) == "1" && string(s[5]) == "3" {
		/*
			if strings.Contains(s[2:5], "010") ||
				strings.Contains(s[2:5], "111") ||
				strings.Contains(s[2:5], "212") ||
				strings.Contains(s[2:5], "313") ||
				strings.Contains(s[2:5], "414") ||
				strings.Contains(s[2:5], "515") ||
				strings.Contains(s[2:5], "616") ||
				strings.Contains(s[2:5], "717") ||
				strings.Contains(s[2:5], "818") ||
				strings.Contains(s[2:5], "919")*/
		if string(s[1]) == "2" && string(s[3]) == "3" && string(s[5]) == "3" {

			mp[int(p[x])] = true
		}
		// }
	}

	y := 0
	l := len(mp)
	max := p[l-1]

	fmt.Printf("There are %d primes under a million\n", l)
	fmt.Println("The largest prime in the range is", max)

	for x = range p {
		if mp[int(p[x])] {
			fmt.Print(p[x], " ")
			y++

			if y%20 == 0 {
				fmt.Println()
			}
		}
	}
}
