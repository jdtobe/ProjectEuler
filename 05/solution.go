package main

import (
	"fmt"
)


func main() {
	i := 2520;
	done := false;

	for i = 2520; !done; i++ {
		done = true
		for c := 1; c <= 20; c++ {
			if i % c != 0 {
				done = false;
				break;
			}
		}

		if done {
			break;
		}
	}

	fmt.Println(i);
}
