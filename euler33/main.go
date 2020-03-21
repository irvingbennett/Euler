package main

import (
	"fmt"
)

// Fraction represents ratio.
type Fraction struct {
	Numerator, Denominator int
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

// Reduce makes given fraction reduced.
func (f Fraction) Reduce() Fraction {
	gcd := GCDEuclidean(f.Numerator, f.Denominator)
	f.Numerator /= gcd
	f.Denominator /= gcd

	return f
}

// MultiplyByNumber multiplies fraction by given number.
func (f Fraction) MultiplyByNumber(m int) Fraction {
	f.Numerator *= m

	return f
}

// MultiplyByFraction multiplies fraction by given fraction.
func (f Fraction) MultiplyByFraction(m Fraction) Fraction {
	f.Numerator *= m.Numerator
	f.Denominator *= m.Denominator

	return f
}

func main() {
	fmt.Println("Digit cancelling fractions")
	pairs := []Fraction{}
	for x := 1; x < 10; x++ {
		for y := 1; y < 10; y++ {
			num := x*10 + y
			dem := y*10 + x
			if num >= dem {
				continue
			}
			f := Fraction{num, dem}
			pairs = append(pairs, f)
		}
	}
	// ten := 0
	for x := range pairs {
		fmt.Print(pairs[x].String(), "  ")
		f := pairs[x].Reduce()
		fmt.Println(f.String())
		/*
			ten++
			if ten == 10 {
				fmt.Println()
				ten = 0
			}
		*/

	}
}

// GCDEuclidean calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
