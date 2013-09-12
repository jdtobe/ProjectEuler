package main

import (
	fmt "fmt"
)



func main() {
	t := 0

	for n:=1; n<10000; n++ {
		if isAmicable(n) {
			t = t + n
		}
	}

	fmt.Println("Total: ", t)
}

func isAmicable(n int) bool {
	f := factorSum(n)

	if n != f && n == factorSum(f) {
		return true
	}

	return false
}


func factor(n int) (r []int) {
	for i:=1; i<n; i++ {
		if n%i == 0 {
			r = append(r, i)
		}
	}

	return
}

func factorSum(n int) int {
	factors := factor(n)
	s := 0

	for _,f := range factors {
		s = s + f
	}

	return s
}
