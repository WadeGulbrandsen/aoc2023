package main

import (
	"fmt"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

func Problem1(filename string) int {
	defer internal.Un(internal.Trace("Problem1"))
	return 0
}

func Problem2(filename string) int {
	defer internal.Un(internal.Trace("Problem2"))
	return 0
}

func main() {
	filename := "input.txt"
	fmt.Println("Advent of Code 2023")
	fmt.Printf("\nThe answer for Day 06, Problem 1 is: %v\n\n", Problem1(filename))
	fmt.Printf("\nThe answer for Day 06, Problem 2 is: %v\n\n", Problem2(filename))
}
