package main

import (
	"fmt"
	"math"
)


func main() {
	var s int64 = 0;

	for i := 2; i < 2000000; i++ {
		f := factors(i);
		if len(f) == 2 {
			s += int64(i);
		}
	}

	fmt.Println(s);
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