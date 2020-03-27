package main

import (
	"fmt"
	"sort"
)

func main() {
	for x := 1; x < 9999999; x++ {
		if isPermutation(x, x*2, x*3, x*4, x*5, x*6) {
			fmt.Println(x, x*2, x*3, x*4, x*5, x*6)
		}
	}
}

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

func isPermutation(x1, x2, x3, x4, x5, x6 int) bool {
	s1 := SortStringByCharacter(fmt.Sprintf("%d", x1))
	if SortStringByCharacter(fmt.Sprintf("%d", x2)) == s1 &&
		SortStringByCharacter(fmt.Sprintf("%d", x3)) == s1 &&
		SortStringByCharacter(fmt.Sprintf("%d", x4)) == s1 &&
		SortStringByCharacter(fmt.Sprintf("%d", x5)) == s1 &&
		SortStringByCharacter(fmt.Sprintf("%d", x6)) == s1 {
		return true
	}
	return false
}
