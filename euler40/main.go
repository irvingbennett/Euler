package main
import(
	"fmt"
	"strings"
	"strconv"
)
func main(){
	fmt.Println("Champernowne's constant")
	digits := []string{}
	for x := 1; x < 999999; x++ {
		s := fmt.Sprintf("%d",x)
		digits = append(digits, s)
	}
	c := strings.Join(digits,"")
	fmt.Println(strconv.Atoi(c[:10]))
	fmt.Println(len(c))
	d1, _ := strconv.Atoi(string(c[1-1]))
	d10, _ := strconv.Atoi(string(c[10-1]))
	d100, _ := strconv.Atoi(string(c[100-1]))
	d1000, _ := strconv.Atoi(string(c[1000-1]))
	d10000, _ := strconv.Atoi(string(c[10000-1]))
	d100000, _ := strconv.Atoi(string(c[100000-1]))
	d1000000, _ := strconv.Atoi(string(c[1000000-1]))
	result := d1*d10*d100*d1000*d10000*d100000*d1000000
	fmt.Println(string(c[1-1]), string(c[10-1]), string(c[100-1]))
	fmt.Println(string(c[1000-1]), string(c[10000-1]), string(c[100000-1]))	
	fmt.Println(result)
}