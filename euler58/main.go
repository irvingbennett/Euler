package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Spiral Primes")
	sides := 999999999999
	sum, primes, count, last := 0, 0, 0, 0
	root := 3
	power := 9
	gap := 2
	corner := true
	diagonals := []int{}
	var r float32
	l := 0
	for x := 1; x <= (sides * sides); x++ {
		if x-last == gap || last == 0 {
			// fmt.Println("X:", x, "Last:", last, "Gap:", gap)
			corner = true
		}

		if corner {
			/*
				fmt.Println("Sum:", sum, "X:", x, "Gap:", gap,
					"Root:", root, "Power:", power, "Last:", last, "Count:", count)
			*/
			corner = false
			diagonals = append(diagonals, x)
			l = len(diagonals)
			if IsPrime(x) {
				primes++
			}
			// one full wrapt
			if count%4 == 0 {
				r = float32(primes) / float32(l) * 100
				/*
					fmt.Printf("The ratio of primes in diagonal is %d/%d, %2.2f%% on side %d\n",
						primes, l, r, root)
				*/
				if r < 10 && x > 4 {
					break
				}

			}
			sum += x
			count++
			last = x
			if x == power {
				gap = root + 1
				root += 2
				power = root * root
			}

		}
	}
	fmt.Printf("The ratio of primes in diagonal is %d/%d, %2.2f%% on side %d\n",
		primes, l, r, root)
	fmt.Println("Sum", sum, "Count:", count)
}

// IsPrime returns true if value is prime
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
