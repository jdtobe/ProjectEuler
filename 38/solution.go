package main

import (
	fmt "fmt"
)


func main() {
	for i:=1; i<10000; i++ {
		s := ""
		for n:=1; ; n++ {
			s = s + fmt.Sprint(n*i)

			ls := len(s)
			if ls > 9 {
				break
			} else if ls == 9 {
				if isPandigital(s) {
					fmt.Println(i, s)
				}

				break
			}
		}
	}
}

func isPandigital(n string) bool {
	if len(n) != 9 {
		return false
	}

	for i:=1; i<10; i++ {
		s := rune('0' + i)
		found := false
		for _,c := range n {
			if (s == c) {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
