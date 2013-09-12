package main

import (
	"fmt"
)


func main() {
	max := 0;

	for x := 100; x < 1000; x++ {
		for y := 100; y < 1000; y++{
			z := x * y;
			if max < z {
				if is_palindrome(z) {
					max = z;
				}
			}
		}
	}

	fmt.Println(max);
}

func is_palindrome(n int) bool {
	s := fmt.Sprint(n);

	l := len(s);

	for i := 0; i < l; i++ {
		e := l - i - 1;

		if s[i] != s[e] {
			return false;
		}
	}

	return true;
}