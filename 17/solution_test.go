package main

import (
	"testing"
)

func TestOneThroughFive(test *testing.T) {
	ch, done, ret := makeCounter()

	totalLength := 19
	words := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	for i,w := range words {
		s := num2str(i)

		if s != w {
			test.Errorf("Result should be '%v' but was '%v'.", w, s)
		}

		ch <- num2str(i)
	}

	done<-true

	r := <-ret
	if totalLength != r {
		test.Errorf("Result should be %v but was %v.", totalLength, r)
	}
}

func TestLengthOf115(test *testing.T) {
	ch, done, ret := makeCounter()

	totalLength := 20

	ch <- num2str(115)

	done<-true

	r := <-ret
	if totalLength != r {
		test.Errorf("Result should be %v but was %v.", totalLength, r)
	}
}

func TestLengthOf342(test *testing.T) {
	ch, done, ret := makeCounter()

	totalLength := 23

	ch <- num2str(342)

	done<-true

	r := <-ret
	if totalLength != r {
		test.Errorf("Result should be %v but was %v.", totalLength, r)
	}
}
