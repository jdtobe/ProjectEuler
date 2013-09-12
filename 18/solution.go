package main

import (
	fmt "fmt"
)



var triangle = [][]int{
/*
	[]int{3},
	[]int{7, 4},
	[]int{2, 4, 6},
	[]int{8, 5, 9, 3},
*/

	[]int{75},
	[]int{95, 64},
	[]int{17, 47, 82},
	[]int{18, 35, 87, 10},
	[]int{20,  4, 82, 47, 65},
	[]int{19,  1, 23, 75,  3, 34},
	[]int{88,  2, 77, 73,  7, 63, 67},
	[]int{99, 65,  4, 28,  6, 16, 70, 92},
	[]int{41, 41, 26, 56, 83, 40, 80, 70, 33},
	[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	[]int{63, 66,  4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	[]int{ 4, 62, 98, 27, 23,  9, 70, 98, 73, 93, 38, 53, 60,  4, 23},
}





func main() {
	ch := make(chan int, 1)

	chain(ch, 0, 0, 0)

	fmt.Println(<-ch)
}

func chain(ch chan int, t, x, y int) {
	t = t + triangle[y][x]
	y = y + 1

	if y == len(triangle) {
		ch  <- t
		return
	}

	chl := make(chan int, 1)
	chain(chl, t, x, y)

	chr := make(chan int, 1)
	chain(chr, t, x+1, y)

	chlv := <-chl
	chrv := <-chr

	if chlv < chrv {
		ch <- chrv
	} else {
		ch <- chlv
	}
}
