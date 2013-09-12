package main

import (
	big "big"
	fmt "fmt"
)





func main() {
	t := big.NewInt(0)
	for n:=1; n<=1000; n++ {
		bn := big.NewInt(1)
		for t:=0; t<n; t++ {
			bn = bn.Mul(bn, big.NewInt(int64(n)))
		}

		t = t.Add(t, bn)
	}

	fmt.Println("Total: ", t)
}
