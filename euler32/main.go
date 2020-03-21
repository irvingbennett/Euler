package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Pandigital Product")
	count := 0
	perms := []string{}
	Perm([]rune("123456789"), func(a []rune) {
		perms = append(perms, string(a))
	})
	fmt.Println(len(perms))
	for x := 0; x < len(perms); x++ {
		if check(perms[x]) {
			count++
		}
	}
	fmt.Println(count)
}

func check(s string) bool {
	for x := 1; x < 7; x++ {
		for y := 1; y < 7; y++ {
			if x+y > len(s) {
				break
			}
			m, _ := strconv.Atoi(s[0:x])
			n, _ := strconv.Atoi(s[x : x+y])
			o, _ := strconv.Atoi(s[x+y:])
			// fmt.Printf("%d x %d = %d\n", m, n, o)
			if m*n == o {
				fmt.Printf("%d x %d = %d\n", m, n, o)
				return true
			}
		}
	}
	return false
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
