package main

import (
	bufio "bufio"
	fmt "fmt"
	io "io"
)


func main() {
	r, w := io.Pipe()
	defer r.Close()

	go func(i int) {
		defer w.Close()

		fmt.Fprint(w, "0.")
		for {
			i++
			fmt.Fprint(w, i)

			if i == 1000000 {
				break
			}
		}
	}(0)

	i := -2
	rr := bufio.NewReader(r)
	for {
		i++
		rune, size, err := rr.ReadRune()
		if err != nil {
			fmt.Println()
			fmt.Println("Error:", err)
			break
		}
		if size == 0 {
			fmt.Println()
			fmt.Println("Size:", size)
			break
		}

/*
		if i==1 || i==10 || i==100 || i==1000 || i==10000 || i==100000 || i==1000000 {
			fmt.Printf("\nd(%d) = '%d'\n", i, string(rune))
		}
*/

		switch i {
		case 1,10,100,1000,10000,100000,1000000:
			fmt.Printf("\nd(%d) = '%d'\n", i, string(rune))
		}

//		fmt.Print(string(rune))
	}

	// ANSWER:
	fmt.Println(1*1*5*3*7*2*1)

}
