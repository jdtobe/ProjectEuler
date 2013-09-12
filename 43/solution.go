package main

import (
	fmt "fmt"
	big "math/big"
)

func main() {
	t := big.NewInt(0)

	p := perm{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for {
		if p.isDivisible() {
			fmt.Println(p)

			fmt.Printf("%v + %v = ", t, p.int())

			t.Add(t, big.NewInt(p.int()))

			fmt.Println(t)
		}

		err := p.next()
		if err != nil {
			break
		}
	}

	fmt.Println(t)
}

type perm []int64

func (pp *perm) next() error {
	p := *pp

	ll := len(p)

	k := ll - 2
	for {
		if p[k] < p[k+1] {
			break
		}

		k--

		if k < 0 {
			return fmt.Errorf("No more permutations")
		}
	}

	l := ll - 1
	for {
		if p[k] < p[l] {
			break
		}

		l--

		if l == k {
			return fmt.Errorf("Something went wrong...")
		}
	}

	p[k], p[l] = p[l], p[k]

	for i, j := k+1, ll-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}

	*pp = p

	return nil
}

func (pp *perm) isDivisible() bool {
	p := *pp

	l := len(p)

	found := true

	s := ""

	for i := 1; i <= l-3; i++ {
		c := (100 * p[i]) + (10 * p[i+1]) + p[i+2]

		switch i {
		case 1:
			found = found && (c%2 == 0)
			n := int64(2)
			s = s + fmt.Sprintf("d%d, d%d, d%d => %v / %v =  %v r%v %v\n", i, i+1, i+2, c, n, c/n, c%n, (c%n == 0))
		case 2:
			found = found && (c%3 == 0)
			n := int64(3)
			s = s + fmt.Sprintf("d%d, d%d, d%d => %v / %v =  %v r%v %v\n", i, i+1, i+2, c, n, c/n, c%n, (c%n == 0))
		case 3:
			found = found && (c%5 == 0)
			n := int64(5)
			s = s + fmt.Sprintf("d%d, d%d, d%d => %v / %v =  %v r%v %v\n", i, i+1, i+2, c, n, c/n, c%n, (c%n == 0))
		case 4:
			found = found && (c%7 == 0)
			n := int64(7)
			s = s + fmt.Sprintf("d%d, d%d, d%d => %v / %v =  %v r%v %v\n", i, i+1, i+2, c, n, c/n, c%n, (c%n == 0))
		case 5:
			found = found && (c%11 == 0)
			n := int64(11)
			s = s + fmt.Sprintf("d%d, d%d, d%d => %v / %v =  %v r%v %v\n", i, i+1, i+2, c, n, c/n, c%n, (c%n == 0))
		case 6:
			found = found && (c%13 == 0)
			n := int64(13)
			s = s + fmt.Sprintf("d%d, d%d, d%d => %v / %v =  %v r%v %v\n", i, i+1, i+2, c, n, c/n, c%n, (c%n == 0))
		case 7:
			found = found && (c%17 == 0)
			n := int64(17)
			s = s + fmt.Sprintf("d%d, d%d, d%d => %v / %v =  %v r%v %v\n", i, i+1, i+2, c, n, c/n, c%n, (c%n == 0))
		}

		/*
			m := int64(2)
			var r int64
			var v int64
			for ; m < c/2; m++ {
				v = c / m
				r = c % m

				if r == 0 {
					found = true

					break
				}
			}

			fmt.Printf("d%d, d%d, d%d => %v / %v =  %v r%v %v\n", i, i+1, i+2, c, m, v, r, f)
		*/

	}

	if found {
		fmt.Print(s)
	}

	return found
}

func (pp *perm) int() int64 {
	p := *pp

	t := int64(0)
	for _, v := range p {
		t = t*10 + v
	}

	return t
}
