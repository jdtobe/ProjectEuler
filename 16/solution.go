package main

import (
	"math"
	"fmt"
)





func main() {
	s := fmt.Sprint(math.Pow(2, 1000))

	fmt.Println(len(s))

	t := 0
	for _, v := range s {
		if tmp := int(v - '0'); tmp >= 0 {
			t += t + int(tmp)
		}
	}

	fmt.Println(t)
}
