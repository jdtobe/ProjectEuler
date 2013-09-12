package main

import (
	fmt "fmt"
	math "math"
)

func main() {
//	fmt.Println(getSides(120))
//	yeilds: [[20 48 52] [24 45 51] [30 40 50]]

	ms := 0
	mp := 0
	for p:=1; p<=1000; p++ {
		s := getSides(p)
		ls := len(s)
		if ls > ms {
			ms = ls
			mp = p
		}
	}

	fmt.Println("Max Sides:", ms)
	fmt.Println("Mas Perimeter:", mp)
}

func getSides(p int) (r [][]int) {

	for a := 1; a < p/3-2; a++ {
		for b := 1; b < p-a-1; b++ {
			c2 := a*a + b*b

			fc := math.Sqrt(float64(c2))
			ic := int(fc)
			if a+b+ic != p {
				continue
			}

			if fc == float64(ic) {
				r = append(r, []int{a, b, ic})
			}
		}
	}

	return
}
