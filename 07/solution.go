package main

import (
	"fmt"
	"math"
)


func main() {
	p := 0;
	pn := 0;

	for i := 1; p <= 10001; i++ {
		f := factors(i);
		if len(f) == 2 {
			p += 1;
			pn = i;
		}
	}

	fmt.Println(pn);
}

func factors(n int) []int {
	f := []int{};

	for i := int(math.Sqrt(float64(n))); i > 0; i-- {
		if n % i == 0 {
			f = append([]int{i}, f...);
			f = append(f, n / i);
		}
	}

	return f
}