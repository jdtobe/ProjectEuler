package main

import (
	"fmt"
	"math"
)


func main() {
	var n int64 = 600851475143;

	f := factors(n);

	var max int64 = 0;
	for _,v := range f {
		sf := factors(v);
		if len(sf) == 2 && v > max {
			max = v;
		}
	}

	fmt.Printf("Max: %d\n", max);
}

func factors(n int64) []int64 {
	f := []int64{};

	for i := int64(math.Sqrt(float64(n))); i > 0; i-- {
		if n % i == 0 {
			f = append([]int64{i}, f...);
			f = append(f, n / i);
		}
	}

	return f
}