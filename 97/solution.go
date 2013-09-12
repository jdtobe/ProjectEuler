package main

import (
	big "math/big"
	fmt "fmt"
)





func main() {
// 28433 * 2^(7830457)+1

	r := big.NewInt(2)

	xr := big.NewInt(2)

	x := 7830457
	for t:=1; t<x; t++ {
		r.Mul(r, xr)
	}

	r.Mul(r, big.NewInt(28433))

	r.Add(r, big.NewInt(1))

	fmt.Println(r)
}
