package main

import (
	big "big"
	fmt "fmt"
)





func main() {
	t := 0
	for n:=1; n<10000; n++ {
		fmt.Print(n, " ")
		if isLychrel(n) {
			t = t + 1
//			fmt.Print(" yup!")
		}
		fmt.Println()
	}

	fmt.Println(t)
}


func reverse(in *big.Int) (out *big.Int) {
	n := big.NewInt(0).Set(in)
	out = big.NewInt(0)
	zero := big.NewInt(0)
	for n.Cmp(zero) > 0 {
		t := big.NewInt(10)
		out.Mul(out, t)

		n.DivMod(n, t, t)
		out.Add(t, out)
	}

	return
}

func isPalindrome(in *big.Int) (out bool) {
	s := fmt.Sprint(in)
	l := len(s)
	for i:=0; i<(l/2); i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}

func isLychrel(in int) (out bool) {
	n := big.NewInt(int64(in))

	for i:=0; i<50; i++ {
		r := reverse(n)
		n.Add(n, r)
		if isPalindrome(n) {
			fmt.Print(i, ":", n)
			return false
		}
	}

	return true
}
