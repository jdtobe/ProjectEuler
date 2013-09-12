package main

import (
	big "big"
	fmt "fmt"
)





func main() {
	maxA := 0
	maxB := 0
	maxSum := 0

	s := 0
	n := big.NewInt(0)
	for a:=1; a<100; a++ {
		for b:=1; b<100; b++ {
			n.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			s = sumDigits(n)

			if maxSum < s {
				maxA = a
				maxB = b
				maxSum = s
			}
		}
	}

	fmt.Println(maxA)
	fmt.Println(maxB)
	fmt.Println(maxSum)
}

func sumDigits(in *big.Int) (out int) {
	s := in.String()

	for _,c := range s {
		out = out + (c - '0')
	}

	return
}
