package main

import (
	fmt "fmt"
)


func main() {
	max := 28123
	abundants := abundantNumbersBelow(max)

	s := 0

	for i:=0; i<max; i++ {
		if !isSumOfAbundants(i, abundants) {
			s = s + i
		}
	}

	fmt.Println(s)
}





func isSumOfAbundants(n int, abundants []bool) (r bool) {
	for i:=0; i<n; i++ {
		if abundants[i] && abundants[n-i] {
			return true
		}
	}

	return
}

func abundantNumbersBelow(below int) (r []bool) {
	r = make([]bool, below)

	for i:=0; i<below; i++ {
		s := 0

		for j:=1; j<i; j++ {
			if i%j == 0 {
				s = s + j
			}
		}

		if i < s {
			r[i] = true
		}
	}

	return
}
