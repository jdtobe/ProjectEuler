package main

import (
	fmt "fmt"
)



var m = []int{
	31,
	28,
	31,
	30,
	31,
	30,
	31,
	31,
	30,
	31,
	30,
	31,
}



func main() {
	y := 1900
	d := 1

	s := 0
	for ; y<=2000; y=y+1 {
		for tm:=0; tm<12; tm=tm+1 {
			d = d + m[tm]
			d = d % 7

			if d == 0 {
				s = s + 1
			}
		}
	}

	fmt.Println(y)
	fmt.Println(s-2)
}
