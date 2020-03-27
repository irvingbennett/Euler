package main

import (
	"fmt"

	"github.com/irvingbennett/prime-factors"
)

func main() {
	fmt.Println("Distinct Primes Factors")
	fmt.Println(prime.Factors(14))
	for x := 1; x < 99999999; x++ {
		if distinct(x) == 4 &&
			distinct(x+1) == 4 &&
			distinct(x+2) == 4 &&
			distinct(x+3) == 4 {

			f1 := prime.Factors(int64(x))
			f2 := prime.Factors(int64(x + 1))
			f3 := prime.Factors(int64(x + 2))
			f4 := prime.Factors(int64(x + 3))
			fmt.Println(x, x+1, x+2, x+3)
			fmt.Println(f1, f2, f3, f4)
			fmt.Println(distinct(x), distinct(x+1), distinct(x+2), distinct(x+3))
			break
		}

	}
}

func distinct(n int) int {
	f := prime.Factors(int64(n))
	d := make(map[int64]bool, 0)
	for p := range f {
		d[f[p]] = true
	}
	return len(d)
}
