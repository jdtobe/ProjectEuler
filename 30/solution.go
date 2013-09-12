package main

import (
	strconv "strconv"
	fmt "fmt"
)


func main() {
	t := 0
	for i:=2; i<1000000; i++ {
		s := fmt.Sprint(i)
		p := 0
//fmt.Println("string: ", s)
		for _,c := range s {
//fmt.Println("trying: ", c)
			n,_ := strconv.Atoi(string(c))
			p = p + n*n*n*n*n
		}
		if p == i {
			fmt.Println("prime: ", p)
			t = t + p
		}
	}

	fmt.Println("total: ", t)
}
