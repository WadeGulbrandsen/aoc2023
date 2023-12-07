package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

type Hand struct {
	red, green, blue int
}

type Game struct {
	id    int
	hands []Hand
}

var max_cubes = Hand{12, 13, 14}

func getGameID(s string) (int, error) {
	after, found := strings.CutPrefix(s, "Game")
	if !found {
		return 0, fmt.Errorf("could not find 'Game' in %v", s)
	}
	return strconv.Atoi(strings.TrimSpace(after))
}

func getHands(s string) []Hand {
	var h []Hand
	for _, hand := range strings.Split(s, ";") {
		r, g, b := 0, 0, 0
		for _, cube := range strings.Split(hand, ",") {
			before, after, found := strings.Cut(strings.TrimSpace(cube), " ")
			if found {
				if v, err := strconv.Atoi(strings.TrimSpace(before)); err == nil {
					switch strings.TrimSpace(after) {
					case "red":
						r = v
					case "green":
						g = v
					case "blue":
						b = v
					}
				}
			}
		}
		h = append(h, Hand{r, g, b})
	}
	return h
}

func stringToGame(s string) (Game, error) {
	before, after, found := strings.Cut(s, ":")
	if !found {
		return Game{}, fmt.Errorf("could not find ':' in %v", s)
	}
	id, err := getGameID(before)
	if err != nil {
		return Game{}, err
	}
	return Game{id, getHands(after)}, nil
}

func validateHand(h Hand) bool {
	if h.red > max_cubes.red || h.green > max_cubes.green || h.blue > max_cubes.blue {
		return false
	}
	return true
}

func validateGame(s string) int {
	g, err := stringToGame(s)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if len(g.hands) == 0 {
		fmt.Printf("Game %+v has no hands\n", g)
		return 0
	}
	for i, hand := range g.hands {
		if !validateHand(hand) {
			fmt.Printf("Game %+v is invalid because hand %v is not valid\n", g, i)
			return 0
		}
	}
	fmt.Printf("Game %+v is valid\n", g)
	return g.id
}

func minDiceNeededForGame(s string) int {
	game, err := stringToGame(s)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if len(game.hands) == 0 {
		fmt.Printf("Game %+v has no hands\n", game)
		return 0
	}
	r, g, b := 0, 0, 0
	for _, hand := range game.hands {
		r = max(r, hand.red)
		g = max(g, hand.green)
		b = max(b, hand.blue)
	}
	pow := r * g * b
	fmt.Printf("Game %+v is valid needs %v red, %v green and %v blue cubes: %v\n", game, r, g, b, pow)
	return pow
}

func Problem1(filename string) int {
	return internal.FileSumSolver(filename, validateGame)
}

func Problem2(filename string) int {
	return internal.FileSumSolver(filename, minDiceNeededForGame)
}

func main() {
	fmt.Println("Advent of Code 2023")
	fmt.Printf("\nThe answer for Day 02, Problem 1 is: %v\n\n", Problem1("input.txt"))
	fmt.Printf("\nThe answer for Day 02, Problem 2 is: %v\n\n", Problem2("input.txt"))
}
