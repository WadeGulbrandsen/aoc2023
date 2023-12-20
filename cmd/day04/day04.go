package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/solve"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/rs/zerolog/log"
)

const Day = 4

type Card struct {
	id              int
	winning_numbers []int
	played_numbers  []int
}

func cardFromString(s string) (Card, error) {
	card := Card{}
	card_info, nums, found := strings.Cut(strings.TrimSpace(s), ": ")
	if !found {
		return card, fmt.Errorf("could not find ': ' in %v", s)
	}
	if c, f := strings.CutPrefix(card_info, "Card "); f {
		if id, err := strconv.Atoi(strings.TrimSpace(c)); err == nil {
			card.id = id
		} else {
			return card, err
		}
	} else {
		return card, fmt.Errorf("could not find 'Card ' in %v", s)
	}
	if winners, played, found := strings.Cut(nums, " | "); found {
		for _, n := range strings.Split(strings.TrimSpace(winners), " ") {
			if v, err := strconv.Atoi(strings.TrimSpace(n)); err == nil {
				card.winning_numbers = append(card.winning_numbers, v)
			}
		}
		for _, n := range strings.Split(strings.TrimSpace(played), " ") {
			if v, err := strconv.Atoi(strings.TrimSpace(n)); err == nil {
				card.played_numbers = append(card.played_numbers, v)
			}
		}
	}
	return card, nil
}

func (c *Card) Play() int {
	wins := 0
	for _, n := range c.played_numbers {
		if slices.Contains(c.winning_numbers, n) {
			wins++
		}
	}
	return wins
}

func (c *Card) Score() (int, string) {
	wins := c.Play()
	if wins > 0 {
		score := 1 << (wins - 1)
		return score, fmt.Sprintf("%v win(s) on %+v: %v", wins, *c, score)
	}
	return 0, fmt.Sprintf("No wins on %+v: 0", *c)
}

func getScore(s string) int {
	card, err := cardFromString(s)
	if err != nil {
		log.Err(err).Msg("card conversion")
		return 0
	}
	score, message := card.Score()
	log.Debug().Msg(message)
	return score
}

func Problem1(data *[]string) int {
	return solve.SumSolver(data, getScore)
}

func Problem2(data *[]string) int {
	copies := make(map[int]int)

	for _, s := range *data {
		if card, err := cardFromString(s); err == nil {
			copies[card.id]++
			wins := card.Play()
			for i := 1; i <= wins; i++ {
				copies[card.id+i] += copies[card.id]
			}
		}
	}

	sum := 0
	for _, v := range copies {
		sum += v
	}
	return sum
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
