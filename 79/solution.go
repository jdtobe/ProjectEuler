package main

import (
	bufio "bufio"
	fmt "fmt"
	io "io"
	log "log"
	os "os"
)

const KEYFILE = "keylog.txt"

type number struct {
	start *digit
}

func (n *number) String() string {
	if n.start == nil {
		return "{}"
	}
	s := "{"

	d := n.start
	for d != nil {
		s += fmt.Sprint(d.value)

		if d.next != nil {
			s += " "
		}

		d = d.next
	}

	return s + "}"
}

func (n *number) add(v1, v2, v3 int) bool {
	if n.start == nil {
		n.start = &digit{}
		n.start.add(v1)
	}

	if !n.containsOne(v1, v2, v3) {
		return false
	}

	n.order(v1, v2)
	n.order(v2, v3)

	return true
}

func (n *number) containsOne(v ...int) bool {
	for _, vi := range v {
		if n.find(vi) != nil {
			return true
		}
	}

	return false
}

func (n *number) find(v int) *digit {
	d := n.start
	for d != nil {
		if d.contains(v) {
			return d
		}

		d = d.next
	}

	return nil
}

func (n *number) order(v1, v2 int) {
	d1 := n.find(v1)
	d2 := n.find(v2)

	if d1 == d2 {
		d1 = &digit{}

		d1.add(v1)
		d2.del(v1)

		d1.prev = d2.prev
		d1.next = d2
		d2.prev = d1

		if d1.prev == nil {
			n.start = d1
		} else {
			d1.prev.next = d1
		}

		return
	}

	if d1 != nil {
		if d2 == nil {
			if d1.next == nil {
				d1.next = &digit{}
				d1.next.prev = d1
			}

			d1.next.add(v2)
		}
	}

	if d2 != nil {
		if d1 == nil {
			if d2.prev == nil {
				d2.prev = &digit{}
				d2.prev.next = d2
			}

			d2.prev.add(v1)
		}
	}
}

type digit struct {
	value []int
	prev  *digit
	next  *digit
}

func (d *digit) add(v int) {
	if !d.contains(v) {
		d.value = append(d.value, v)
	}
}

func (d *digit) del(v int) {
	for i, vi := range d.value {
		if vi == v {
			d.value = append(d.value[:i], d.value[i+1:]...)
		}
	}
}

func (d *digit) contains(v int) bool {
	for _, vi := range d.value {
		if vi == v {
			return true
		}
	}

	return false
}

func main() {
	lines := readLines(KEYFILE)

	res := doLines(lines)

	fmt.Println(res)
}

func readLines(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fb := bufio.NewReader(f)

	lines := [][]int{}
	for {
		line, _, err := fb.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		l := string(line)
		if len(l) != 3 {
			continue
		}

		a := int(l[0] - '0')
		b := int(l[1] - '0')
		c := int(l[2] - '0')

		lines = append(lines, []int{a, b, c})
	}

	return lines
}

func doLines(lines [][]int) []int {
	// nothing to do
	if len(lines) == 0 {
		return []int{}
	}

	// only one answer
	if len(lines) == 1 {
		return lines[0]
	}

	// initialize number object with first line numbers
	n := &number{}

	// loop through lines until list exhausted
	for done := false; !done; {
		done = true
		for i := 0; i < len(lines); i++ {
			if !n.add(lines[i][0], lines[i][1], lines[i][2]) {
				done = false
			}
		}
	}

	/*
		this doesn't work.
		there is a bug in the code where it doesn't verify the order correctly.

		for len(lines) > 0 {
			for i := len(lines) - 1; i >= 0; i-- {
				fmt.Print(lines[i])

				if n.add(lines[i][0], lines[i][1], lines[i][2]) {
					lines = append(lines[:i], lines[i+1:]...)
				}

				fmt.Println("", n)
				fmt.Scanln()
			}
		}
	*/

	r := []int{}
	for d := n.start; d != nil; d = d.next {
		r = append(r, d.value[0])
	}

	return r
}
