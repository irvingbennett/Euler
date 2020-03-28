package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// RANK is the order of card values
var RANK = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

// Card is just one card
type Card struct {
	rank string
	suit string
}

// Less returns true if card i is higher than card j
func (c Card) Less(i, j Card) bool {
	return RANK[i.rank] > RANK[j.rank]
}

// Cards is a deck of cards, or a hand
type Cards []Card

// Lens is the length of Cards
func (c Cards) Len() int           { return len(c) }
func (c Cards) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Cards) Less(i, j int) bool { return c[i].Less(c[i], c[j]) }

// Count is the count of similar cards in a hand
type Count struct {
	count int
	rank  string
}

// Hand is just a hand of cards, it has a score and a value
type Hand struct {
	cards  Cards
	rank   map[string]int
	counts []Count
	hand   string
	score  int
	value  int
}

// HANDS is a list of hand types
var HANDS = []string{"High card", "One pair", "Two pair",
	"Three of a kind", "Straight", "Flush", "Full house",
	"Four of a kind", "Straight flush",
}

var score map[string][]string

// Score returns the score of a hand of cards
func Score(h []string) (hand Hand) {
	rank := "AKQJT98765432"
	lowStraight := "A5432"
	royal := "AKQJT"

	var cards Cards
	for _, card := range h {
		c := []rune(card)
		if len(c) == 3 {
			fmt.Println("--------------------------------")
			fmt.Println("*****", card, len(card), len(c), "*****")
			fmt.Println("--------------------------------")
			cards = append(cards, Card{"T", string(c[2])})
		} else {
			cards = append(cards, Card{string(card[0]), string(c[1])})
		}

	}
	sort.Sort(cards)
	hand.cards = cards
	r := map[string]int{}
	for _, card := range cards {
		r[card.rank]++
	}
	hand.rank = r
	counts := []Count{}
	for rank, count := range r {
		counts = append(counts, Count{count, rank})
	}
	hand.counts = counts
	// Royal Flush
	royalFlush := false

	// Flush check
	flush := true

	// Straight check
	straight := false
	lowstraight := false
	s := []string{}
	suit := hand.cards[0].suit
	for _, card := range hand.cards {
		if card.suit != suit {
			flush = false
		}
		s = append(s, card.rank)
	}
	sequence := strings.Join(s, "")

	if strings.Contains(rank, sequence) || strings.Contains(lowStraight, sequence) {
		if strings.Contains(lowStraight, sequence) {
			fmt.Println("Low straight")
			lowstraight = true
		}
		straight = true
		if strings.Contains(royal, sequence) {
			royalFlush = true
		}
	}

	// Four of a kind check
	four := false
	fullHouse := false
	three := false
	twoPair := false
	onePair := false
	pairs := 0
	for x := range counts {
		if counts[x].count == 4 {
			four = true
		}
		if counts[x].count == 3 {
			three = true
		}
		if counts[x].count == 2 {
			pairs++
			onePair = true
		}
	}
	if pairs == 2 {
		fmt.Println("Two pairs")
		twoPair = true
	}
	if len(counts) == 2 {
		if counts[0].count == 2 || counts[0].count == 3 {
			fullHouse = true
		}
	}

	// SCORE is a list of best hands
	var SCORE = map[string]int{
		"Royal flush":     649737, //  A, K, Q, J, 10, all the same suit.
		"Straight flush":  72193,  // contains five cards of sequential rank, all of the same suit
		"Four of a kind":  4164,   // four equal cards
		"Full house":      693,    // three of a kind and a pair
		"Flush":           508,    // five cards of the same suit
		"Straight":        253,    // five cards in sequence
		"Three of a kind": 46,
		"Two pair":        20,
		"One pair":        2,
		"High card":       1,
	}
	switch {
	case flush && royalFlush:
		// fmt.Println("*** Royal Flush ***")
		hand.hand = "*** Royal Flush ***"
		hand.score = SCORE["Royal flush"]
		hand.value = SCORE["Royal flush"]
	case flush && straight:
		fmt.Println("Straight Flush")
		hand.hand = "Straight Flush"
		hand.score = SCORE["Straight flush"]
		hand.value = RANK[hand.cards[0].rank]
		if lowstraight {
			hand.value = RANK[hand.cards[1].rank]
		}
	case flush && three:
		// fmt.Println("Flush")
		hand.hand = "Flush"
		hand.score = SCORE["Flush"] + (SCORE["Three of a kind"] * 2)
		hand.value = RANK[hand.cards[0].rank]

	case flush && twoPair:
		fmt.Println("Flush & Two Pair")
		hand.hand = "Flush"
		hand.score = SCORE["Flush"] + (SCORE["Two pair"] * 2)
		hand.value = RANK[hand.cards[0].rank]
		highRank := 0
		for _, count := range counts {
			if count.count == 2 {
				if highRank < RANK[count.rank] {
					highRank = RANK[count.rank]
				}
				if hand.value < RANK[count.rank] {
					hand.value = RANK[count.rank]
				}
			}
		}
		hand.score += highRank

	case flush && onePair:
		// fmt.Println("Flush")
		hand.hand = "Flush"
		hand.score = SCORE["Flush"] + (SCORE["One pair"] * 2)
		hand.value = RANK[hand.cards[0].rank]

	case four:
		// fmt.Println("Four of a kind")
		hand.hand = "Four of a kind"
		hand.score = SCORE["Four of a kind"]
		for _, count := range counts {
			if count.count == 4 {
				hand.value = RANK[count.rank]
			}
		}

	case fullHouse:
		// fmt.Println("Full House")
		hand.hand = "Full House"
		hand.score = SCORE["Full house"]
		hand.value = RANK[hand.cards[0].rank]
		for _, count := range counts {
			if count.count == 3 {
				hand.value = RANK[count.rank]
			}
		}
	case flush:
		// fmt.Println("Flush")
		hand.hand = "Flush"
		hand.score = SCORE["Flush"]
		hand.value = RANK[hand.cards[0].rank]
	case straight:
		// fmt.Println("Straight")
		hand.hand = "Straight"
		hand.score = SCORE["Straight"]
		hand.value = RANK[hand.cards[0].rank]
		if lowstraight {
			hand.value = RANK[hand.cards[1].rank]
		}
	case three:
		// fmt.Println("Three of a kind")
		hand.hand = "Three of a kind"
		hand.score = SCORE["Three of a kind"]
		for _, count := range counts {
			if count.count == 3 {
				hand.value = RANK[count.rank]
			}
		}
	case twoPair:
		// fmt.Println("Two Pairs")
		hand.hand = "Two pair"
		hand.score = SCORE["Two pair"]
		highRank := 0
		for _, count := range counts {
			if count.count == 2 {
				if highRank < RANK[count.rank] {
					highRank = RANK[count.rank]
				}
				if hand.value < RANK[count.rank] {
					hand.value = RANK[count.rank]
				}
			}
		}
		hand.score += highRank
	case onePair:
		// fmt.Println("One Pair")
		hand.hand = "One pair"
		hand.score = SCORE["One pair"] + 13
		for _, count := range counts {
			if count.count == 2 {
				if hand.value < RANK[count.rank] {
					hand.value = RANK[count.rank]
				}
			}
		}
	default:
		// fmt.Println("High Card")
		hand.hand = "High card"
		hand.score = SCORE["High card"]
		hand.value = RANK[hand.cards[0].rank]
	}

	// fmt.Println(hand)

	return
}

func main() {
	p1, p2 := 0, 0
	fmt.Println("Poker Hands")
	n, err := ioutil.ReadFile("poker.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(n))
	hands := strings.Split(string(n), "\n")
	for x := range hands {
		if len(hands[x]) > 1 {
			winner := check(hands[x])
			if winner == 1 {
				p1++
			}
			if winner == 2 {
				p2++
			}
			// fmt.Println(check(hands[x]))
		}
	}
	fmt.Println("Player one wins", p1, "Player two wins", p2)
}

func check(h string) int {
	hands := strings.Split(h, " ")
	// fmt.Println(hands)

	p1 := hands[0:5]
	p2 := hands[5:]
	h1 := Score(p1)
	h2 := Score(p2)
	// fmt.Println(h1, h2)

	return winner(h1, h2)
}

func winner(h1, h2 Hand) int {
	if h1.score > h2.score {
		return 1
	}
	if h1.score < h2.score {
		return 2
	}
	// fmt.Println(h1, h2)
	if h1.value > h2.value {
		return 1
	}
	if h1.value < h2.value {
		return 2
	}

	fmt.Println(h1)
	fmt.Println(h2)

	if RANK[h1.cards[0].rank] > RANK[h2.cards[0].rank] {
		return 1
	}
	if RANK[h1.cards[0].rank] < RANK[h2.cards[0].rank] {
		return 2
	}
	fmt.Println(h1)
	fmt.Println(h2)
	return 0
}
