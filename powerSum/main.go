package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	var sum int
	fmt.Println("Power Sum")
	fmt.Println(math.Pow(2, 1000))
	n := big.NewInt(2)
	m := big.NewInt(1000)
	l := big.NewInt(0)
	z := n.Exp(n, m, l)
	fmt.Println(n.Text(10))
	digits := z.Text(10)
	for x := 0; x < len(digits); x++ {
		// fmt.Printf("%s", string(digits[x]))
		i, err := strconv.Atoi(string(digits[x]))
		if err != nil {
			fmt.Println(err)
		}
		sum += i
	}
	fmt.Println(sum)
}
