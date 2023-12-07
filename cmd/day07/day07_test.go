package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

// Answers
const p1sample = 6440
const p1input = 250474325
const p2sample = 5905
const p2input = 248909434

func TestSolutions(t *testing.T) {
	sample := internal.FileToLines("sample.txt")
	input := internal.FileToLines("input.txt")
	t.Run(fmt.Sprintf("Day %v Problem1 with sample.txt", Day), func(t *testing.T) {
		if a, r := p1sample, Problem1(&sample); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run(fmt.Sprintf("Day %v Problem1 with input.txt", Day), func(t *testing.T) {
		if a, r := p1input, Problem1(&input); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run(fmt.Sprintf("Day %v Problem2 with sample.txt", Day), func(t *testing.T) {
		if a, r := p2sample, Problem2(&sample); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run(fmt.Sprintf("Day %v Problem2 with input.txt", Day), func(t *testing.T) {
		if a, r := p2input, Problem2(&input); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
