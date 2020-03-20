package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Distinct Powers")
	fmt.Println("----------------------")
	fmt.Println()
	one := big.NewInt(1)
	a := big.NewInt(2)
	b := big.NewInt(2)
	z := big.NewInt(0)
	powers := make(map[[2]int]string, 99999)
	for x := 2; x <= 100; x++ {
		for y := 2; y <= 100; y++ {
			_, found := powers[[2]int{x, y}]
			if !found {
				z = z.Exp(a, b, nil)
				powers[[2]int{x, y}] = z.String()
			}
			b.Add(b, one)
		}
		b = big.NewInt(2)
		a.Add(a, one)
	}
	values := make(map[string][][2]int, 99999)
	for k, v := range powers {
		values[v] = append(values[v], k)
	}
	/*
		v := make([]string, 0)
		for k := range values {
			v = append(v, k)
		}
			v = sort.StringSlice(v)
			for _, k := range v {
				fmt.Printf("%s, ", k)
			}
			fmt.Println()
	*/
	fmt.Println("Values:", len(values))
}
