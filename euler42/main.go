package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Triangle Words")
	n, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(n))
	words := strings.Split(string(n), ",")
	for x := range words {
		words[x] = strings.ReplaceAll(words[x], "\"", "")
		// fmt.Println(names[x])
	}
	fmt.Println("t of n10 is", tn(10))
	values := make(map[int]bool)
	for x := 0; x < 100; x++ {
		values[tn(x)] = true
	}
	tws := []string{}
	for x := range words {
		if values[value(words[x])] {
			tws = append(tws, words[x])
		}
	}
	fmt.Println("There are", len(tws), "in", tws)
	fmt.Println("SKY is", value("SKY"))
}

func value(n string) (w int) {
	for x := range n {
		w += int(n[x]) - int('@')
	}
	return
}

func tn(n int) int {
	return n * (n + 1) / 2
}
