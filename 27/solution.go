package main

import (
//	math "math"
//	big "big"
	fmt "fmt"
)


func main() {
// n² + an + b, where |a|  1000 and |b|  1000

	minA := -1000
	maxA := 1000

	minB := -1000
	maxB := 1000

	minN := 0
	maxN := 1000


	fa := 0
	fb := 0
	fn := 0


	for a:=minA; a<=maxA; a++ {
		for b:=minB; b<=maxB; b++ {
			for n:=minN; n<=maxN; n++ {
				tn := (n*n) + (a*n) + b

				if !isPrime(tn) {
					break
				}


				if n  >= fn {
					fa = a
					fb = b
					fn = n
				}
			}
		}
	}

	fmt.Println(fa, fb, fn)
	fmt.Println(fa*fb)
}





func isPrime(n int) (r bool) {
	if n < 0 {
		n = -n
	}

	if n <= 1 {
		return false
	}

	for i:=2; i<=n/2; i++ {
		if n%i == 0 || n%(n-i) == 0 {
			return false
		}
	}

	return true
}
