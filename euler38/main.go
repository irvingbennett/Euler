package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pandigital struct {
	number  int
	digits  []string
	product string
}

func main() {
	pds := []pandigital{}
	fmt.Println("Pandigital Multiples")
	fmt.Println("123456789 is pandigital is", isPandigital("123456789"))
	for n := 1; n <= 9999999999; n++ {
		products := []string{}
		for m := 1; m <= 9; m++ {
			p := strconv.Itoa(n * m)
			products = append(products, p)
		}
		pd, ok := check(products, n)
		if ok {
			// fmt.Println("It's true...")
			pds = append(pds, pd)
		}
	}
	for x := range pds {
		fmt.Println(pds[x])
	}

}

func check(p []string, n int) (pd pandigital, ok bool) {
	// fmt.Println(p)
	digits := []string{}
	for x := 1; x <= len(p); x++ {
		digits = append(digits, fmt.Sprintf("%d", x))
		c := strings.Join(p[:x], "")
		if len(c) != 9 {
			continue
		}
		if isPandigital(c) {
			ok = true
			pd = pandigital{
				n,
				digits,
				c,
			}
			fmt.Println(pd)
			return
		}
	}

	ok = false
	return
}

func isPandigital(s string) bool {
	// fmt.Println(s)
	allDigits := map[string]bool{
		"1": false,
		"2": false,
		"3": false,
		"4": false,
		"5": false,
		"6": false,
		"7": false,
		"8": false,
		"9": false,
	}
	if len(s) != 9 {
		return false
	}
	for _, c := range s {
		if allDigits[string(c)] == false {
			allDigits[string(c)] = true
		} else {
			return false
		}
	}
	for x := range allDigits {
		if !allDigits[x] {
			return false
		}
	}
	return true
}
