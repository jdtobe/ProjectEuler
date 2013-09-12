package main

import (
	fmt "fmt"
	big "math/big"
)

func main() {
	//    40!
	// ---------
	// 20! * 20!

	// 40!
	f := fact(40)

	fmt.Println(f)

	// 20! * 20!
	t := fact(20)
	t = t.Mul(t, t)

	fmt.Println(t)

	// result
	f = f.Div(f, t)

	fmt.Println(f)
}

func fact(x int) *big.Int {
	r := big.NewInt(int64(x))

	for i := 1; i < x; i++ {
		r = r.Mul(r, big.NewInt(int64(i)))
	}

	return r
}
