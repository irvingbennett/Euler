package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Lychrel Numbers")
	p := new(big.Int)

	l := []int{}
	for x := 0; x <= 10000; x++ {
		s := big.NewInt(int64(x))
		for y := 0; y <= 50; y++ {
			r := []byte(s.String())
			reverse([]byte(r))
			reversed := new(big.Int)
			reversed, ok := reversed.SetString(string(r), 10)
			if !ok {
				fmt.Println("Error reversing:", r)
			}
			p = p.Add(s, reversed)
			if isPalindrome(p.String()) {
				/*
					fmt.Print(x, " ")
					fmt.Println("Is palindrome:", p, isPalindrome(p.String()), y+1, "reverses")
				*/
				break
			}
			s = p
		}
		if !isPalindrome(p.String()) {
			l = append(l, x)
		}
	}
	fmt.Println("Lychrels is ", len(l))
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return
}

func isPalindrome(s string) bool {
	b := []byte(s)
	reverse(b)
	if s == string(b) {
		return true
	}
	return false

}
