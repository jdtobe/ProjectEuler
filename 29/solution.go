package main

import (
	big "math/big"
	fmt "fmt"
)


func main() {
	r := []*big.Int{}

	for a:=2; a<=100; a++ {
		for b:=2; b<=100; b++ {
			n := pow(a, b)

			found := false
			for _,t := range r {
				if n.Cmp(t) == 0 {
					found = true
					break
				}
			}

			if !found {
				r = append(r, n)
			}
		}
	}

	fmt.Println(len(r))
}


func pow(a,b int) (r *big.Int) {
	ba := big.NewInt(int64(a))
	r = big.NewInt(int64(a))

	for i:=1; i<b; i++ {
		r.Mul(r, ba)
	}

	return
}
