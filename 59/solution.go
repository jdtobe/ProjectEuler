package main

import (
	bufio "bufio"
	fmt "fmt"
	log "log"
	os "os"
	strconv "strconv"
)

func main() {
	f, err := os.Open("cipher1.txt")
	if err != nil {
		log.Fatal(err)
	}

	var s []rune
	r := bufio.NewReader(f)
	for {
		b, berr := r.ReadBytes(',')
		for l := len(b); b[l-1] < '0' || b[l-1] > '9'; l = len(b) {
			b = b[:l-1]
		}

		i, ierr := strconv.Atoi(string(b))
		fmt.Println(ierr)
		if ierr == nil {
			s = append(s, rune(i))
		}

		if berr != nil {
			break
		}
	}

	fmt.Println(s)

	var aa, bb, cc rune
	max := 0
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			for c := 'a'; c <= 'z'; c++ {
				ts := string(doXor(s, []rune{a, b, c}))

				t := 0
				for _, tc := range ts {
					if tc == ' ' || (tc >= 'a' && tc <= 'z') || (tc >= 'A' && tc <= 'Z') {
						t++
					}
				}

				if t > max {
					max = t
					aa = a
					bb = b
					cc = c
				}
			}
		}
	}

	fmt.Println(string(doXor(s, []rune{aa, bb, cc})))
	fmt.Printf("%c,%c,%c\n", aa, bb, cc)

	t := 0
	for _, c := range doXor(s, []rune{aa, bb, cc}) {
		t += int(c)
	}

	fmt.Println(t)
}

func doXor(in, p []rune) []rune {
	pl := len(p)
	tl := len(in)
	r := []rune{}
	for t := 0; t < tl; t++ {
		r = append(r, in[t]^p[t%pl])
	}

	return r
}
