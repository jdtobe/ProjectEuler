package main

import (
	"fmt"
	"math"
)


func main() {
	t := 0;
	f := []int{};
	for i := 1; len(f) < 500; i++ {
		t += i;
		f = factors(t);
	}

	fmt.Println(t);
	fmt.Println(f);
}

func factors(n int) []int {
	f := []int{};

	for i := int(math.Sqrt(float64(n))); i > 0; i-- {
		if n % i == 0 {
			f = append([]int{i}, f...);
			if n / i != i {
				f = append(f, n / i);
			}
		}
	}

	return f
}