package main

import (
	fmt "fmt"
)

func main() {
	s := 0
	ch := fib()
	for v := range ch {
		if v > 4000000 {
			break
		}

		if v%2 == 0 {
			s += v
		}
	}

	fmt.Println(s)
}

func fib() <-chan int {
	a := 1
	b := 1

	ch := make(chan int)

	go func() {
		for {
			a, b = b, a+b

			ch <- a
		}
	}()

	return ch
}
