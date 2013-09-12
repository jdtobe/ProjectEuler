package main

import "fmt"

func main() {
	r := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			r += i
		}
	}

	fmt.Println(r)
}
