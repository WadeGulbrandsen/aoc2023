package main

import (
	"fmt"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 0

func Problem1(filename string) int {
	defer internal.Un(internal.Trace(fmt.Sprintf("Day %v Problem1 with %v", Day, filename)))
	return 0
}

func Problem2(filename string) int {
	defer internal.Un(internal.Trace(fmt.Sprintf("Day %v Problem1 with %v", Day, filename)))
	return 0
}

func main() {
	filename := "input.txt"
	fmt.Println("Advent of Code 2023")
	fmt.Printf("\nThe answer for Day %v, Problem 1 is: %v\n\n", Day, Problem1(filename))
	fmt.Printf("\nThe answer for Day %v, Problem 2 is: %v\n\n", Day, Problem2(filename))
}
