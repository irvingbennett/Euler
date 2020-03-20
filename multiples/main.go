package main
import "fmt"

// SumMultiples returns the sum of multiples up to limit
func SumMultiples(l int, n ...int) (s int) {
	m := make(map[int]bool)

	for x := 1; x < l; x++ {
		for _, i := range n {
			if i != 0 && x%i == 0 {
				m[x] = true
			}
		}
	}
	for x := range m {
		s += x
	}

	return
}

func main() {
	fmt.Println(SumMultiples(1000, 3, 5))
}