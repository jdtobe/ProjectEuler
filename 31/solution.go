package main

import (
	fmt "fmt"
	strings "strings"
)

var coinNames = []string{
	"£2",
	"£1",
	"50p",
	"20p",
	"10p",
	"5p",
	"2p",
	"1p",
}

var coinValues = []int{
	200,
	100,
	50,
	20,
	10,
	5,
	2,
	1,
}

func main() {
	cs := findCombinations(200)
	for i, tc := range cs {
		fmt.Print(i+1, ": ")
		printCombination(tc)
	}
}

// find denomination combinations that sum to t pence
func findCombinations(total int) (r [][]int) {
	var countDown func(t, i int, v []int)
	countDown = func(t, i int, v []int) {
		if i == len(coinValues)-1 {
			nt := make([]int, len(v)+1)
			copy(nt, append(v, t))
			r = append(r, nt)

			return
		}

		cv := coinValues[i]
		for n := t / cv; n >= 0; n-- {
			countDown(t-(n*cv), i+1, append(v, n))
		}
	}

	countDown(total, 0, []int{})

	return
}

func printCombination(c []int) {
	p := []string{}

	for i, n := range c {
		if n > 0 {
			p = append(p, fmt.Sprintf("%vx%v", n, coinNames[i]))
		}
	}

	fmt.Println(strings.Join(p, " + "))
}
