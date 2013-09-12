package main

import (
	fmt "fmt"
	strings "strings"
)

func main() {
	p := map[int]bool{}
	t := 0
	z := 0
	for x := 1; ; x++ {
		for y := 1; ; y++ {
			z = x * y

			s := fmt.Sprintf("%d%d%d", x, y, z)

			if len(s) < 9 {
				continue
			}
			if len(s) > 9 {
				y = 0
				break
			}

			if isPandigital(s) {
				if _, ok := p[z]; !ok {
					p[z] = true

					t += z

					fmt.Println("t =", t)
					fmt.Println(x, "*", y, "=", z)
				}
			}
		}

		if len(fmt.Sprint((z - x))) > 9 {
			break
		}
	}

	fmt.Println("Total:", t)
}

var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func isPandigital(in string) bool {
	for _, s := range numbers {
		if !strings.Contains(in, s) {
			return false
		}
	}

	return true
}
