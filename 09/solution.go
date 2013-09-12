package main

import (
	"fmt"
)


func main() {
	i := 0;

	for a := 0; a < 1000; a++ {
		for b := a; b < 1000 - a; b++ {
			c := 1000 - a - b;
			if a*a + b*b != c*c {
				continue;
			}

			i = a * b * c;
		}
	}

	fmt.Println(i);
}
