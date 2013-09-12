package main

import (
	big "big"
	fmt "fmt"
)



func main() {
//	fmt.Println("10 =?", C(5, 3))
//	fmt.Println("1144066 =?", C(23, 10))

	oneMillion := big.NewInt(1000000)
	num := 0
	for n:=0; n<=100; n++ {
		for r:=0; r<=n; r++ {
			if C(n, r).Cmp(oneMillion) >= 0 {
				num = num + 1

				fmt.Println(n, r)
			}
		}
	}

	fmt.Println()
	fmt.Println("Answer:", num)
}

func C(n, r int) *big.Int {
	if r > n || n < 0 || r < 0 {
		return nil
	}

	a := fact(n)
	b := fact(r)
	c := fact(n - r)

	b.Mul(b, c)
	a.Div(a, b)

	return a
}

func fact(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}

	r := big.NewInt(int64(n))

	return r.Mul(r, fact(n-1))
}

