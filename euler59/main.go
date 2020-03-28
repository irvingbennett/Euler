package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("XOR decrytion")
	fmt.Println()
	letters := []byte{}
	for x := 97; x < 123; x++ {
		letters = append(letters, byte(x))
	}
	fmt.Println(string(letters))
	content, err := ioutil.ReadFile("cipher.txt")
	if err != nil {
		fmt.Println(err)
	}
	c := strings.Split(string(content), ",")
	cBytes := []byte{}
	for x := range c {
		a, err := strconv.Atoi(c[x])
		if err != nil {
			fmt.Println("Error:", err)
		}
		cBytes = append(cBytes, byte(a))
	}
	fmt.Println()
	words := make(map[string][]string)
	keys := []byte{}
	max := 0
	var best string
	for x := 0; x < len(letters); x++ {
		keys = append(keys, letters[x])
		for y := 0; y < len(letters); y++ {
			keys = append(keys, letters[y])
			for z := 0; z < len(letters); z++ {
				keys := []byte{}
				keys = append(keys, letters[x])
				keys = append(keys, letters[y])
				keys = append(keys, letters[z])
				w := strings.Split(EncryptDecrypt(string(cBytes), string(keys)), " ")
				words[string(keys)] = w
				if len(w) > max {
					max = len(w)
					best = string(keys)
				}
				// fmt.Println("============================")
			}
		}
	}
	fmt.Println(max, best)
	fmt.Println("============================")
	w := EncryptDecrypt(string(cBytes), best)
	fmt.Println(w)
	sum := 0
	for x := range w {
		sum += int(w[x])
	}
	fmt.Println("Sum:", sum)
}

// EncryptDecrypt runs a XOR encryption on the input string, encrypting it if it hasn't already been,
// and decrypting it if it has, using the key provided.
func EncryptDecrypt(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output
}
