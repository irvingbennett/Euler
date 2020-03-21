package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Binary palindromes")
	fmt.Println(585, "is palindrome", isPalindrome(585))
	fmt.Println(1001001001, "is palindrome", isPalindrome(1001001001))
	fmt.Println()
	// palindromes := []int{}
	sum := 0
	for x := 0; x < 1000000; x++ {
		if isPalindrome(x) {
			fmt.Printf("%d is %b\n", x, x)
			sum += x
		}
	}
	fmt.Println("Sum is", sum)
}
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func isPalindrome(i int) bool {
	s := strconv.Itoa(i)
	b := []byte(s)
	reverse(b)

	binary := fmt.Sprintf("%b", i)
	reversed := []byte(binary)
	reverse(reversed)

	if s == string(b) && binary == string(reversed) {
		return true
	}
	return false

}
