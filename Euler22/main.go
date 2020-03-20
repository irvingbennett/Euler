package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	var sum int
	fmt.Println("Name Score")
	n, err := ioutil.ReadFile("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(n))
	names := strings.Split(string(n), ",")
	for x := range names {
		names[x] = strings.ReplaceAll(names[x], "\"", "")
		// fmt.Println(names[x])
	}
	sort.Strings(names)

	for x := range names {
		names[x] = strings.ReplaceAll(names[x], "\"", "")
		sum += worth(names[x]) * (x + 1)
	}
	fmt.Println("Sum", sum)
}

func worth(n string) (w int) {
	for x := range n {
		w += int(n[x]) - int('@')
	}
	return
}
