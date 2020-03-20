package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Reciprocal cycles")
	fmt.Println("=======================")
	fmt.Println()
	const prec = 500
	d := new(big.Float).SetPrec(prec).SetInt64(2)
	decimals := []string{}
	dMax := 200
	// d := big.NewFloat(2)
	one := big.NewFloat(1)
	t := new(big.Float)
	fmt.Println(len("010309278350515463917525773195876288659793814432989690721649484536082474226804123711340206185567"))
	for x := 2; x <= dMax; x++ {
		t.Quo(one, d)
		decimals = append(decimals, fmt.Sprintf("%1.200f", t))
		fmt.Printf("1/%3d is %s\n", x, decimals[x-2])
		d.Add(d, one)
	}

}

func rc(n string) (l int) {
	return l
}
