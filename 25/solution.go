package main

import (
	big "big"
	fmt "fmt"
)


func main() {
	fiboCache = make(map[string]*big.Int)

	for i,a,running:=big.NewInt(1),big.NewInt(1),true; running; i.Add(i, a) {
		f := fibo(i)

		l := len(f.String())

		if l >= 1000 {
			running = false

			fmt.Println(i)
		}
	}
}



var fiboCache map[string]*big.Int
func fibo(n *big.Int) *big.Int {
	s := n.String()
	if r,ok := fiboCache[s]; ok {
		return r
	}

	i0 := big.NewInt(0)
	if n.Cmp(i0) <= 0 {
		return i0
	}
	if n.Cmp(big.NewInt(2)) <= 0 {
		return big.NewInt(1)
	}

	n1 := big.NewInt(-1)
	n1.Add(n1, n)
	n1 = fibo(n1)

	n2 := big.NewInt(-2)
	n2.Add(n2, n)
	n2 = fibo(n2)

	i0.Add(n1, n2)
	fiboCache[s] = i0
	return i0
}
