package main

import (
	fmt "fmt"
)





func main() {
	fmt.Println(permutation(9, 1000000))
}





func fact(n int) (r int) {
	r = 1

	for i:=n; i>1; i-- {
		r = r * i
	}

	return
}


func span(s, e int) (r []int) {
	for n:=s; n<e; n++ {
		r = append(r, n)
	}

	return
}


func permutation(size, index int) (r []int) {
	i := index-1
	n := size
	p := span(0, n+1)

	for _,k := range span(0, n+1) {
		f := fact(n-k)
		d := i / f
		i = i % f

		r = append(r, p[d])
		p = append(p[:d], p[d+1:]...)
	}

	return
}

