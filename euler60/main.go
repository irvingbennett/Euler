package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kavehmz/prime"
)

var mp map[int]bool

func main() {
	fmt.Println("Prime Pair Sets")
	p := prime.Primes(10000000000) // generate a slice of primes
	mp = make(map[int]bool, 0)     // create a map of generated primes
	for x := range p {
		mp[int(p[x])] = true
	}
	var x int
	for x = 0; x < len(p); x++ {
		if p[x] == 673 {
			break
		}
	}
	p = p[x+1:]
	fmt.Println(p[0], mp[int(p[0])])

	l := len(p)
	max := p[l-1]

	fmt.Printf("There are %d primes under a trillion\n", l)
	fmt.Println("The largest prime in the range is", max)
	x = 0
	for {
		// fmt.Println(p[x])
		check := tryPrimes(p[x])
		if check {
			fmt.Println("3, 7, 109, 673 and", p[x], "adds to", 3+7+109+673+p[x])
			break
		}
		x++
		if x == l {
			break
		}
	}

}

func tryPrimes(n uint64) bool {
	p3 := "3"
	p7 := "7"
	p109 := "109"
	p673 := "673"
	pn := fmt.Sprint(n)

	p, err := strconv.Atoi(strings.Join([]string{p3, pn}, ""))
	// fmt.Print(p, " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !mp[p] {
		// fmt.Println()
		return false
	}

	p, err = strconv.Atoi(strings.Join([]string{p7, pn}, ""))
	// fmt.Print(p, " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !mp[p] {
		// fmt.Println()
		return false
	}

	p, err = strconv.Atoi(strings.Join([]string{p109, pn}, ""))
	// fmt.Print(p, " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !mp[p] {
		// fmt.Println()
		return false
	}

	p, err = strconv.Atoi(strings.Join([]string{p673, pn}, ""))
	// fmt.Print(p, " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !mp[p] {
		// fmt.Println()
		return false
	}

	p, err = strconv.Atoi(strings.Join([]string{pn, p3}, ""))
	// fmt.Print(p, " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !mp[p] {
		// fmt.Println()
		return false
	}

	p, err = strconv.Atoi(strings.Join([]string{pn, p7}, ""))
	// fmt.Print(p, " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !mp[p] {
		// fmt.Println()
		return false
	}

	p, err = strconv.Atoi(strings.Join([]string{pn, p109}, ""))
	// fmt.Print(p, " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !mp[p] {
		// fmt.Println()
		return false
	}

	p, err = strconv.Atoi(strings.Join([]string{pn, p673}, ""))
	// fmt.Print(p, " ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !mp[p] {
		// fmt.Println()
		return false
	}
	// fmt.Println()
	return true
}
