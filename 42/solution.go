package main

import (
	strings "strings"
	ioutil "io/ioutil"
	fmt "fmt"
	log "log"
)

const WORDSFILE = "words.txt"


func main() {
	triangles := map[int]bool{}
	for i:=1.0; i<=26; i++ {
		t := int((i/2)*(i+1))
		triangles[t] = true
	}
//	fmt.Println("Triangles:", triangles)

	count := 0
	words := readWords()
	for _,word := range words {
		t := 0
		for _,l := range word {
			t = t + int(l) - 'A' + 1
		}

		if v,ok := triangles[t]; ok && v {
			count++
		}
	}
	fmt.Println("Count:", count)
//	fmt.Println("Words:", words)
//	fmt.Println("Num Words:", len(words))
}

func readWords() (r []string) {
	f,err := ioutil.ReadFile(WORDSFILE)
	if err != nil {
		log.Fatal(err)
	}

	w := strings.Split(string(f), ",")

	for _,word := range w {
		l := len(word)
		if l > 2 {
			r = append(r, word[1:l-1])
		}
	}

	return
}
