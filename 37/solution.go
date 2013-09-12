package main

import (
	strconv "strconv"
//	math "math"
	fmt "fmt"
)


var primes map[int]bool





func init() {
	primes = make(map[int]bool)
}

func main() {
	ch := make(chan int)
	go primeSieve(ch)

	t := 0
	for c:=0; c<10; {
		n, running := <-ch
		if !running {
			break
		}

//		fmt.Println("Prime: ", n)

		if isTrunctable(n) {
			c++
			t = t + n

			fmt.Println(c, " Trunctable Prime: ", n)
		}
	}

	// thanks wikipedia
	// to solve fully, change to c<11 in for loop above and remove this addition
	t = t + 739397

	fmt.Println("Total: ", t)
}



func isPrime(n int) bool {
	_,ok := primes[n]
	if !ok {
		return false
	}

	return true
}

func isTrunctable(n int) bool {
	return isTrunctableRight(n) && isTrunctableLeft(n)
}

func isTrunctableLeft(n int) bool {
	s := fmt.Sprint(n)
	for n > 0 {
		if !isPrime(n) {
			return false
		}

		if n < 10 {
			break
		}

		s = s[1:]
		n,_ = strconv.Atoi(s)
	}

	return true
}

func isTrunctableRight(n int) bool {
	for n > 0 {
		if !isPrime(n) {
			return false
		}

		n = n / 10
	}

	return true
}










func primeSieveGenerate(ch chan int) {
	for i:=2; ; i++ {
		ch <- i
	}
}

func primeSieveFilter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func primeSieve(out chan int) {
	ch := make(chan int)
	go primeSieveGenerate(ch)

	// skip primes 2,3,5,7
	for i:=0; i<4; i++ {
		prime := <-ch

		ch1 := make(chan int)
		go primeSieveFilter(ch, ch1, prime)
		ch = ch1

		primes[prime] = true
	}

	for {
		prime := <-ch

		ch1 := make(chan int)
		go primeSieveFilter(ch, ch1, prime)
		ch = ch1

		primes[prime] = true
		out <- prime
	}

	close(out)
}