package main

import (
	"fmt"
)

type Coins struct {
	Total int
	P200  int
	P100  int
	P50   int
	P20   int
	P10   int
	P5    int
	P2    int
	P1    int
}

func (c *Coins) Sum() int {
	c.Total = c.P200*200 + c.P100*100 +
		c.P50*50 + c.P20*20 + c.P10*100 +
		c.P5*5 + c.P2*2 + c.P1*1
	return c.Total
}

func main() {
	fmt.Println("Coin Sums")
	b := new(Coins)
	b.P200 = 1
	fmt.Println(b.Sum())
	fmt.Println(b)
	s := []int{1, 2, 5, 10, 20, 50, 100, 200}
	n := 200
	m := 8
	fmt.Println(count(s, m, n))
}

func count(s []int, m, n int) int {
	// If n is 0 then there is 1 solution
	// (do not include any coin)
	if n == 0 {
		return 1
	}
	// If n is less than 0 then no
	// solution exists
	if n < 0 {
		return 0
	}
	// If there are no coins and n
	// is greater than 0, then no
	// solution exist
	if m <= 0 && n >= 1 {
		return 0
	}
	return count(s, m-1, n) + count(s, m, n-s[m-1])
}
