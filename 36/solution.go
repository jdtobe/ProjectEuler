package main

import (
	fmt "fmt"
)


func main() {
	t := 0
	for i:=0; i<1000000; i++ {
		if isPalindrome(i) && isBinaryPalindrome(i) {
			t = t + i
		}
	}

	fmt.Println(t)
}

func isPalindrome(n int) bool {
	s := fmt.Sprint(n)
	l := len(s)

	for i,c := range s {
		if c != int(s[l-i-1]) {
			return false
		}
	}

	return true
}

func isBinaryPalindrome(n int) bool {
	s:=""
	for n > 0 {
		if n&1 == 1 {
			s = "1"+ s
		} else {
			s = "0"+ s
		}

		n = n >> 1
	}

	l := len(s)
	for i,c := range s {
		if c != int(s[l-i-1]) {
			return false
		}
	}

	return true
}
