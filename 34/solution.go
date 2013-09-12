package main

import (
	fmt "fmt"
)


func main() {
	s := 0
	for i:=3; i<1000000; i++ {
		t := 0
		o := i
		for o > 0 {
			f := fact(o % 10)

			t = t + f
			o = o / 10
		}

		if t == i {
			s = s + t
			fmt.Println(t)
		}
	}

	fmt.Println(s)
}

func fact(n int) (r int) {
	r = 1

	for n > 0 {
		r = r * (n%10)

		n = n - 1
	}

	return
}
