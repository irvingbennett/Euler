package main

import (
	"fmt"

	"github.com/jbarham/primegen.go"
)

type pandigital struct {
	number  int
	digits  []string
	product string
}

func main() {
	fmt.Println("Pandigital Prime")
	fmt.Println("12345 is pandigital is", isPandigital("12345"))

	pps := PandigitalPrimes(999999999)
	for x := range pps {
		fmt.Println(pps[x])
	}

}

// SumPrimes adds all primes below a value
func PandigitalPrimes(n int) (pps []int) {
	if n == 1 {
		return []int{2}
	}
	p := primegen.New()
	var next int
	for {
		next = int(p.Next())
		if next < n {
			s := fmt.Sprintf("%d", next)
			if isPandigital(s) > 0 {
				pps = append(pps, next)
			}
		} else {
			break
		}
	}

	return
}

func isPandigital(s string) int {
	// fmt.Println(s)
	d := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
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
		"0": true,
	}
	for _, c := range s {
		if allDigits[string(c)] == false {
			allDigits[string(c)] = true
		} else {
			return -1
		}
	}
	l := len(s)
	n := 0
	for x := 0; x < l; x++ {
		if n == l {
			break
		}
		// fmt.Println(x, d[x], allDigits[d[x]])
		if !allDigits[d[x]] {
			return -1
		}
		n++
	}
	// fmt.Println(s)
	return n
}
