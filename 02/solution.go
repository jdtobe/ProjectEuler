package main

import "fmt"

func main() {
	a := 1
	b := 2
	t := 2
	f := 0

	for f < 4000000 {
		f = a + b

		if f%2 == 0 {
			t += f
		}

		a = b
		b = f
	}

	fmt.Println(t)
}
