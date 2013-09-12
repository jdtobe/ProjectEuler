package main

import (
	strings "strings"
	sort "sort"
	fmt "fmt"
)



func main() {
	for x:=1; ; x++ {
		x2 := x*2
		x3 := x*3
		x4 := x*4
		x5 := x*5
		x6 := x*6
		if sameDigits(x, x2) && sameDigits(x2, x3) && sameDigits(x3, x4) && sameDigits(x4, x5) && sameDigits(x5, x6) {
			fmt.Println(x)
			break
		}
	}
}

func sameDigits(i1,i2 int) bool {
	s1 := fmt.Sprint(i1)
	s2 := fmt.Sprint(i2)

	if len(s1) != len(s2) {
		return false
	}

	l1 := strings.Split(s1, "")
	sort.Strings(l1)

	l2 := strings.Split(s2, "")
	sort.Strings(l2)

	for k,_ := range l1 {
		if l1[k] != l2[k] {
			return false
		}
	}

	return true
}

