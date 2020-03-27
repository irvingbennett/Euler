package main
import (
	"fmt"
	"github.com/kavehmz/prime"
)
func main(){
	fmt.Println("Prime Permutations")
	p:=prime.Primes(10000)
	for x := 0; x < len(p); x++ {
		if p[x] > 1000 {
			p = p[x:]
			break
		}
	}
	mp := make(map[int]bool,0)
	for x:= range p {
		mp[int(p[x])]=true
	}
	for x := range p {
		if mp[int(p[x])] && mp[int(p[x])+3330] && mp[int(p[x])+6660] {
			fmt.Println(p[x], p[x]+3330, p[x]+3330*2)
		}
		
	}
}
