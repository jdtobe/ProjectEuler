package main

import (
	fmt "fmt"
	time "time"
)

func main() {
	tch := make(chan int)
	go triangle(tch)
	t := <-tch

	pch := make(chan int)
	go pentagonal(pch)
	p := <-pch

	hch := make(chan int)
	go hexagonal(hch)
	h := <-hch

	for {
		if t == p && p == h {
			fmt.Println(t, p, h)

			t = <-tch
			p = <-pch
			h = <-hch
			time.Sleep(250 * time.Millisecond)
		}

		if t < p || t < h {
			t = <-tch
			continue
		}

		if p < t || p < h {
			p = <-pch
			continue
		}

		if h < t || h < p {
			h = <-hch
			continue
		}
	}
}

func triangle(ch chan int) {
	//Triangle	 	Tn=n(n+1)/2	 	1, 3, 6, 10, 15, ...
	for n := 1.0; ; n++ {
		ch <- int(n * (n + 1) / 2)
	}
}

func pentagonal(ch chan int) {
	//Pentagonal	 	Pn=n(3n-1)/2	 	1, 5, 12, 22, 35, ...
	for n := 1.0; ; n++ {
		ch <- int(n * (3*n - 1) / 2)
	}
}

func hexagonal(ch chan int) {
	//Hexagonal	 	Hn=n(2n-1)	 	1, 6, 15, 28, 45, ...
	for n := 1.0; ; n++ {
		ch <- int(n * (2*n - 1))
	}
}
