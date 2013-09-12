package main

import (
	"fmt"
)


func main() {
	sum_of_squares := 0;
	square_of_sums := 0;

	for i := 1; i <= 100; i++ {
		sum_of_squares += i * i;
		square_of_sums += i;
	}
	square_of_sums = square_of_sums * square_of_sums;

	fmt.Println(square_of_sums - sum_of_squares);
}
