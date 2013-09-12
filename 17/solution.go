package main

import (
	math "math"
	fmt "fmt"
)





func main() {
	ch, done, ret := makeCounter()
	for i:=1; i<=1000; i++ {
		s := num2str(i)
		ch <- s
		fmt.Println(s)
	}

	done<-true
	fmt.Println(<-ret)
}










func makeCounter() (ch chan string, done chan bool, ret chan int) {
	ch = make(chan string)
	done = make(chan bool)
	ret = make(chan int)

	go func() {
		l := 0

		for {
			select {
			case s:=<-ch:
				for _,c := range s {
					if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
						l++
					}
				}
			case <-done:
				ret <- l
				break
			}
		}
	}()

	return
}





var numbers = []string{
	"",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = []string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

var groups = []string{
	"",
	"thousand",
	"million",
	"billion",
	"trillion",
}

func num2str(n int) (r string) {
	if n < 0 {
		return
	}

	if n >= 1000 {
		group := int(math.Floor(math.Log(float64(n)) / math.Log(1000)))
		r = num2str(n / 1000) +" "+ groups[group] +" "+ num2str(n % 1000)
	} else if n >= 100 {
		r = num2str(n / 100) +" hundred"

		if n%100 > 0 {
			r = r +" and "+ num2str(n % 100)
		}
	} else if n >= 20 {
		s := tens[n/10]

		r = s

		if n%10 > 0 {
			r = r +"-"+ num2str(n % 10)
		}
	} else {	
		s := numbers[n]

		r = s
		n = 0
	}

	return
}
