package main

import (
	"fmt"
)

func main() {
	fmt.Println("Amicable Numbers")
	f1 := []int{}
	f2 := []int{}
	fSum := 0
	for x := 1; x < 10000; x++ {
		f1 = factors(x)
		sum := sumOfFactors(f1)
		f2 = factors(sum)
		sum2 := sumOfFactors(f2)
		if sum2 == x && x != sum {
			// if sum2 == x {
			fSum += x
			fmt.Print(x, f1, sum, f2, sum2)
			fmt.Println(" --> Friendly")
		}
	}
	fmt.Println(fSum)
}

func factors(n int) (f []int) {
	for x := 1; x <= n/2; x++ {
		if n%x == 0 {
			f = append(f, x)
		}
	}
	// f = append(f, n)
	return
}

func sumOfFactors(fs []int) (sum int) {
	for _, f := range fs {
		sum += f
	}
	return
}
