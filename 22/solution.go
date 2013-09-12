 package main

import (
	strings "strings"
	ioutil "io/ioutil"
	sort "sort"
	fmt "fmt"
	os "os"
)



func main() {
	namesTxt,err := ioutil.ReadFile("names.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Reading File: %s\n", err)
		os.Exit(1)
	}

	names := strings.Split(string(namesTxt), ",")
	sort.Strings(names)

	total := 0
	for k,v := range names {
		v = v[1:len(v)-1]
		s := score(v)

		total = total + ((k+1) * s)

//		fmt.Println(v, " -> ", s)
	}

	fmt.Println("Total: ", total)
}

func score(name string) int {
	total := 0
	for _,c := range name {
		total = total + int(c - 'A' + 1)
	}

	return total
}

