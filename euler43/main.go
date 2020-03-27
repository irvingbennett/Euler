package main

import (
	"fmt"
	"strconv"
)

type pandigital struct {
	number  int
	digits  []string
	product string
}

func main() {
	fmt.Println("Sub-string Divisibility")
	fmt.Println("1406357289 is pandigital is", isPandigital("1406357289"))
	fmt.Println(check("1406357289"))

	sum := 0
	perms := []string{}
	Perm([]rune("1234567890"), func(a []rune) {
		perms = append(perms, string(a))
	})
	fmt.Println(len(perms))
	for x := 0; x < len(perms); x++ {
		if check(perms[x]) {
			v, err := strconv.Atoi(perms[x])
			if err != nil {
				fmt.Println(err)
			}
			sum += v
		}
	}
	fmt.Println("The sum is", sum)
}

func check(s string) bool {
	n1, _ := strconv.Atoi(s[1:4])
	if n1%2 != 0 {
		return false
	}
	n2, _ := strconv.Atoi(s[2:5])
	if n2%3 != 0 {
		return false
	}
	n3, _ := strconv.Atoi(s[3:6])
	if n3%5 != 0 {
		return false
	}
	n4, _ := strconv.Atoi(s[4:7])
	if n4%7 != 0 {
		return false
	}
	n5, _ := strconv.Atoi(s[5:8])
	if n5%11 != 0 {
		return false
	}
	n6, _ := strconv.Atoi(s[6:9])
	if n6%13 != 0 {
		return false
	}
	n7, _ := strconv.Atoi(s[7:10])
	if n7%17 != 0 {
		return false
	}

	return true

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

func isPandigital(s string) int {
	// fmt.Println(s)
	d := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	allDigits := map[string]bool{
		"0": false,
		"1": false,
		"2": false,
		"3": false,
		"4": false,
		"5": false,
		"6": false,
		"7": false,
		"8": false,
		"9": false,
	}
	for _, c := range s {
		if allDigits[string(c)] == false {
			allDigits[string(c)] = true
		} else {
			return -1
		}
	}
	l := len(s)
	n := 0
	for x := 0; x < l; x++ {
		if n == l {
			break
		}
		// fmt.Println(x, d[x], allDigits[d[x]])
		if !allDigits[d[x]] {
			return -1
		}
		n++
	}
	// fmt.Println(s)
	return n
}
