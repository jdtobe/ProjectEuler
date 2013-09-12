package main

import (
	strconv "strconv"
	big "big"
	fmt "fmt"
)

func main() {
	t := big.NewInt(1)

	for n:=100; n>1; n=n-1 {
		t.Mul(t, big.NewInt(int64(n)))

		fmt.Println(n)
	}

	fmt.Println(t)

	s := fmt.Sprint(t)
	sum := 0
	for l:=len(s); l>0; l=len(s) {
		d,_ := strconv.Atoi(string(s[l-1]))

		s = s[0:l-1]
		sum = sum + d
	}

	fmt.Println(sum)
}
