package main

import (
	"cmp"
	"slices"
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/mowshon/iterium"
)

const Day = 7

var cardValues = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardValuesWithJokers = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 0,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type HandStrength int

const (
	HighCard HandStrength = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	cards    string
	bid      int
	strength HandStrength
}

func cmpHands(a, b Hand) int {
	if n := cmp.Compare(a.strength, b.strength); n != 0 {
		return n
	}
	for i, c := range a.cards {
		if n := cmp.Compare(cardValues[c], cardValues[rune(b.cards[i])]); n != 0 {
			return n
		}
	}
	return cmp.Compare(a.bid, b.bid)
}

func cmpHandsWithJokers(a, b Hand) int {
	if n := cmp.Compare(a.strength, b.strength); n != 0 {
		return n
	}
	for i, c := range a.cards {
		if n := cmp.Compare(cardValuesWithJokers[c], cardValuesWithJokers[rune(b.cards[i])]); n != 0 {
			return n
		}
	}
	return cmp.Compare(a.bid, b.bid)
}

type HandList []Hand

func (h HandList) Sort() {
	slices.SortFunc(h, cmpHands)
}

func stringToHandStrengthWithJokers(s string) HandStrength {
	if !strings.ContainsRune(s, 'J') {
		return stringToHandStrength(s)
	}
	jokers := strings.Count(s, "J")
	if jokers == 5 {
		return FiveOfAKind
	}
	no_jokers := strings.Map(func(r rune) rune {
		if r == 'J' {
			return -1
		}
		return r
	}, s)
	var unique []rune
	for _, r := range no_jokers {
		if !slices.Contains(unique, r) {
			unique = append(unique, r)
		}
	}
	combos := iterium.CombinationsWithReplacement(unique, jokers)
	var results []HandStrength
	for combo := range combos.Chan() {
		newhand := no_jokers + string(combo)
		result := stringToHandStrengthWithJokers(newhand)
		results = append(results, result)
	}
	return slices.Max(results)
}

func stringToHandStrength(s string) HandStrength {
	counts := make(map[rune]int)
	for _, card := range s {
		counts[card]++
	}
	switch len(counts) {
	case 4:
		return OnePair
	case 3:
		cv := utils.GetMapValues(counts)
		if slices.Max(cv) == 3 {
			return ThreeOfAKind
		}
		return TwoPair
	case 2:
		cv := utils.GetMapValues(counts)
		if slices.Max(cv) == 4 {
			return FourOfAKind
		}
		return FullHouse
	case 1:
		return FiveOfAKind
	default:
		return HighCard
	}
}

func stringToHand(s string, fn func(string) HandStrength) Hand {
	if b, a, f := strings.Cut(s, " "); f {
		if i, err := strconv.Atoi(strings.TrimSpace(a)); err == nil {
			if cards := strings.TrimSpace(b); len(cards) == 5 {
				return Hand{cards: cards, bid: i, strength: fn(cards)}
			}
		}
	}
	return Hand{}
}

func linesToHandList(data *[]string, fn func(string) HandStrength) HandList {
	ch := make(chan Hand)
	for _, line := range *data {
		go func(s string, c chan Hand) {
			c <- stringToHand(s, fn)
		}(line, ch)
	}

	var hl HandList
	for i := 0; i < len(*data); i++ {
		h := <-ch
		if len(h.cards) == 5 {
			hl = append(hl, h)
		}
	}
	return hl
}

func Problem1(data *[]string) int {
	hl := linesToHandList(data, stringToHandStrength)
	hl.Sort()
	sum := 0
	for i, h := range hl {
		sum += (i + 1) * h.bid
	}
	return sum
}

func Problem2(data *[]string) int {
	hl := linesToHandList(data, stringToHandStrengthWithJokers)
	slices.SortFunc(hl, cmpHandsWithJokers)
	sum := 0
	for i, h := range hl {
		sum += (i + 1) * h.bid
	}
	return sum
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
