package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Self Powers")
	one := big.NewInt(1)
	sum := big.NewInt(0)
	exp := big.NewInt(1)
	z := big.NewInt(0)
	for x := 0; x < 1000; x++ {
		sum.Add(sum, z.Exp(exp, exp, nil))
		fmt.Println(exp, sum)
		exp.Add(exp, one)
	}
}
