package main

import (
	"fmt"

	"github.com/WadeGulbrandsen/aoc2023/day01"
	"github.com/WadeGulbrandsen/aoc2023/day02"
	"github.com/WadeGulbrandsen/aoc2023/day03"
)

func main() {
	fmt.Println("Advent of Code 2023")
	fmt.Printf("\nThe answer for Day 01, Problem 1 is: %v\n\n", day01.Problem1("day01/input.txt"))
	fmt.Printf("\nThe answer for Day 01, Problem 2 is: %v\n\n", day01.Problem2("day01/input.txt"))
	fmt.Printf("\nThe answer for Day 02, Problem 1 is: %v\n\n", day02.Problem1("day02/input.txt"))
	fmt.Printf("\nThe answer for Day 02, Problem 2 is: %v\n\n", day02.Problem2("day02/input.txt"))
	fmt.Printf("\nThe answer for Day 03, Problem 1 is: %v\n\n", day03.Problem1("day03/input.txt"))
	fmt.Printf("\nThe answer for Day 03, Problem 2 is: %v\n\n", day03.Problem2("day03/input.txt"))
}
