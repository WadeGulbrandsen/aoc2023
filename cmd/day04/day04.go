package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

func getScore(s string, ch chan int) {
	card, err := cardFromString(s)
	if err != nil {
		fmt.Println(err)
		ch <- 0
		return
	}
	score, message := card.Score()
	fmt.Println(message)
	ch <- score
}

func problemSolver(filename string, fn func(string, chan int)) int {
	fmt.Printf("Opening %v\n", filename)
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	ch := make(chan int)

	l := 0
	for fileScanner.Scan() {
		go fn(fileScanner.Text(), ch)
		l++
	}
	readFile.Close()

	sum := 0
	for i := 0; i < l; i++ {
		sum += <-ch
	}

	return sum
}

func Problem1(filename string) int {
	return problemSolver(filename, getScore)
}

func Problem2(filename string) int {
	copies := make(map[int]int)
	fmt.Printf("Opening %v\n", filename)
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if card, err := cardFromString(fileScanner.Text()); err == nil {
			copies[card.id]++
			wins := card.Play()
			for i := 1; i <= wins; i++ {
				copies[card.id+i] += copies[card.id]
			}
		}
	}
	readFile.Close()

	sum := 0
	for _, v := range copies {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Advent of Code 2023")
	fmt.Printf("\nThe answer for Day 04, Problem 1 is: %v\n\n", Problem1("input.txt"))
	fmt.Printf("\nThe answer for Day 04, Problem 2 is: %v\n\n", Problem2("input.txt"))
}
