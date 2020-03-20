package main

import (
	"fmt"
)

func main() {
	fmt.Println("Non-abundant Sums")
	fmt.Println()
	var abundant string
	var abundantsMap = map[int]int{}
	var abundants = []int{}
	var deficients = []int{}
	var perfects int
	for x := 1; x <= 28123; x++ {
		fs := factors(x)
		sumFs := sumOfFactors(fs)
		switch {
		case x == sumFs:
			abundant = "Perfect!"
			perfects++
		case sumFs > x:
			abundant = "abundant"
			abundantsMap[x] = x
			abundants = append(abundants, x)
		case sumFs < x:
			abundant = "deficient"
			deficients = append(deficients, x)
		}
		if x <= 100 {
			fmt.Printf("%3d has sum %3d and is %s\n", x, sumFs, abundant)
		}
	}
	fmt.Println()
	odd := 0
	odds := []int{}
	for x := range abundants {
		if abundants[x]%2 != 0 {
			odd++
			odds = append(odds, abundants[x])
		}
	}
	fmt.Println("Abundants", "evens are", len(abundants)-odd)
	fmt.Println("=========================================")
	fmt.Println(abundants[:20])
	fmt.Println("Deficients", len(deficients))
	fmt.Println("=========================================")
	fmt.Println(deficients[:20])
	fmt.Println("Abundants", "odds are", odd)
	fmt.Println("=========================================")
	fmt.Println(odds)
	// fmt.Println(abundantsMap[12])
	fmt.Println()
	fmt.Println("Starting trial")
	var sum int
	var count int
	var Sums int
	for x := 1; x < 28123; x++ {
		if !hasSum(x, abundants, abundantsMap) {
			sum += x
			count++
		} else {
			Sums++
		}
	}
	fmt.Println("=========================================")
	fmt.Println("Sum is", sum)
	fmt.Println("Count is", count)
	fmt.Println("Has sums", Sums)
	fmt.Println("Deficients is", len(deficients))
	fmt.Println("Total of count plus sums", count+Sums)
	fmt.Println("Abundants are", len(abundants))
	fmt.Println("Total:", len(deficients)+len(abundants))
	fmt.Println("Perfects are:", perfects)
	fmt.Println()
	fmt.Println()

}

func hasSum(d int, a []int, m map[int]int) bool {
	x := 0
	// y := 1
	first := a[x]
	// second := a[y]
	for {
		if d < first {
			return false
		}
		if d > first {
			if d == first+m[d-first] {
				// fmt.Println(d, "has sum", first, m[d-first])
				return true
			}
		}
		x++
		if x > len(a) {
			break
		}
		first = a[x]
	}
	return false
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
