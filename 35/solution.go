package main

import (
	strconv "strconv"
	math "math"
	fmt "fmt"
)


func main() {
	t := 0

	for i:=2; i<1000000; i++ {
		if isCircularPrime(i) {
			t = t + 1
			fmt.Println(i, " is a circular prime")
		}
	}

	fmt.Println(t, " primes found")
}

func isPrime(n int) bool {
	for i:=int(math.Sqrt(float64(n))); i>1; i-- {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func isCircularPrime(n int) bool {
	s := fmt.Sprint(n)
	prime := true

	l := len(s)
	for i:=0; i<l; i++ {
		t,_ := strconv.Atoi(s)

		if !isPrime(t) {
			prime = false
			break
		}

		s = s[1:] + string(s[0])
	}

	return prime
}
