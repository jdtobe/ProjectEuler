package main

import (
	fmt "fmt"
)


func main() {
	fmt.Println(spiralSum(1001))
}

func spiralSum(s int) (r int) {
	for x,i,c:=1,2,0; x<=s*s; x=x+i {
		r = r + x

		if c%4 == 0 && c != 0 {
			i = i + 2
		}
		c++
	}

	return
}
