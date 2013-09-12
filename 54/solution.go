package main

import (
	bufio "bufio"
	fmt "fmt"
	io "io"
	log "log"
	os "os"
	sort "sort"
)

const GAMEFILE = "poker.txt"

type game struct {
	hands []hand
}

type hand struct {
	cards []card
}

func (h hand) Len() int {
	return len(h.cards)
}

func (h hand) Less(i, j int) bool {
	if h.cards[i].value < h.cards[j].value {
		return true
	}

	/*
		if h.cards[i].suit < h.cards[j].suit {
			return true
		}
	*/

	return false
}

func (h hand) Swap(i, j int) {
	h.cards[i], h.cards[j] = h.cards[j], h.cards[i]
}

func (h hand) String() string {
	return fmt.Sprint(h.cards[0], h.cards[1], h.cards[2], h.cards[3], h.cards[4])
}

type card struct {
	value int
	suit  int
}

func (c card) String() string {
	cs := ""
	for i, v := range cardSuit {
		if v == c.suit {
			cs = string(i)

			break
		}
	}

	cv := ""
	for i, v := range cardValue {
		if v == c.value {
			cv = string(i)

			break
		}
	}

	return fmt.Sprintf("%v%v", cs, cv)
}

var cardValue = map[uint8]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardSuit = map[uint8]int{
	'C': 1,
	'D': 2,
	'H': 3,
	'S': 4,
}

const (
	HIGH_CARD = 20 * iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
	ROYAL_FLUSH
)

func main() {
	games := loadFile(GAMEFILE)

	p1, p2 := 0, 0
	for _, g := range games {
		s1 := scoreHand(g.hands[0])
		s2 := scoreHand(g.hands[1])

		if s1 > s2 {
			p1++
		}

		if s1 < s2 {
			p2++
		}
	}

	fmt.Println("Player 1 wins:", p1)
	fmt.Println("Player 2 wins:", p2)
}

func loadFile(filename string) []game {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	games := []game{}

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		if len(line) == 0 {
			continue
		}

		h1 := parseHand(string(line[:14]))
		h2 := parseHand(string(line[15:]))
		games = append(games, game{[]hand{h1, h2}})
	}

	return games
}

func parseHand(in string) hand {
	h := hand{}

	for i := 0; i < 5; i++ {
		c := card{
			cardValue[in[i*3]],
			cardSuit[in[i*3+1]],
		}
		h.cards = append(h.cards, c)
	}

	return h
}

func scoreHand(h hand) int {
	sort.Sort(h)

	high_card := 0
	for i := 0; i < 5; i++ {
		if high_card < h.cards[i].value {
			high_card = h.cards[i].value
		}
	}

	straight := 0
	for i := 0; i < 5-1; i++ {
		if h.cards[i].value+1 != h.cards[i+1].value {
			break
		}

		if i == 3 {
			straight = h.cards[i+1].value
		}
	}

	flush := true
	for i := 0; i < 5-1; i++ {
		if h.cards[i].suit != h.cards[i+1].suit {
			flush = false

			break
		}
	}

	counts := map[int]int{}
	for _, v := range cardValue {
		for i := 0; i < 5; i++ {
			if v == h.cards[i].value {
				counts[v]++
			}
		}
	}

	four_of_a_kind := 0
	three_of_a_kind := 0
	one_pair := 0
	two_pair := 0

	for v, c := range counts {
		if c == 4 {
			four_of_a_kind = v
		}

		if c == 3 {
			three_of_a_kind = v
		}

		if c == 2 {
			if one_pair > 0 {
				two_pair = v
				if one_pair > two_pair {
					one_pair, two_pair = two_pair, one_pair
				}
			} else {
				one_pair = v
			}
		}
	}

	full_house := 0
	if one_pair > 0 && three_of_a_kind > 0 {
		full_house = three_of_a_kind
	}

	if flush && h.cards[0].value == 10 {
		return ROYAL_FLUSH
	}

	if straight > 0 && flush {
		return STRAIGHT_FLUSH + straight
	}

	if four_of_a_kind > 0 {
		return FOUR_OF_A_KIND + four_of_a_kind
	}

	if full_house > 0 {
		return FULL_HOUSE + full_house
	}

	if flush {
		return FLUSH + high_card
	}

	if straight > 0 {
		return STRAIGHT + straight
	}

	if three_of_a_kind > 0 {
		return THREE_OF_A_KIND + three_of_a_kind
	}

	if two_pair > 0 {
		return TWO_PAIR + two_pair
	}

	if one_pair > 0 {
		return ONE_PAIR + one_pair
	}

	return high_card
}
