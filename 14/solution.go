package main

import (
	"fmt"
)





func main() {
	ml := 0;
	mr :=[]int{}

	for i := 0; i < 1000000; i++ {
		res := count(i)

		if ml < len(res) {
			ml = len(res)
			mr = res
		}
	}

	fmt.Println(ml)
	fmt.Println(mr)
}

func count(n int) (ret []int) {
	odd := func(n int) int { return 3*n + 1 }
	even := func(n int) int { return n/2 }

	ret = append(ret, n)

	for n > 1 {
		if n%2 == 0 {
			n = even(n)
		} else {
			n = odd(n)
		}

		ret = append(ret, n)
	}

	return
}


/*
	l := 0;
	r := []int64{};

	var i int64;
	for i = 0; i < 1000000; i++ {
		fmt.Printf("Trying: %d\n", i);

		t := chain(i);
		if l < len(t) {
			l = len(t);
			r = t;
		}
	}

	fmt.Printf("%d: ", l);
	fmt.Println(r);
}

func chain(n int64) []int64 {
	if n <= 1 {
		return []int64{n};
	}

	c := []int64{n};
	if n % 2 == 0 {
		return append(c, chain(n/2)...);
	}

	return append(c, chain(3*n + 1)...);
}
*/