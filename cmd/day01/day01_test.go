package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Answers
const ans1sample = 142
const ans1input = 54708
const ans2sample = 281
const ans2input = 54087

// filenames
const file1sample = "sample.txt"
const file1input = "input.txt"
const file2sample = "sample2.txt"
const file2input = file1input

func TestSolutions(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	testCases := []struct {
		problem  int
		answer   int
		filename string
		fn       func(*[]string) int
	}{
		{1, ans1sample, file1sample, Problem1},
		{1, ans1input, file1input, Problem1},
		{2, ans2sample, file2sample, Problem2},
		{2, ans2input, file2input, Problem2},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Day %v Problem %v with %v", Day, tc.problem, tc.filename), func(t *testing.T) {
			data := utils.FileToLines(tc.filename)
			if r := tc.fn(&data); tc.answer != r {
				t.Fatalf("The correct answer is %v but received %v", tc.answer, r)
			}
		})
	}
}

func TestGetNumber(t *testing.T) {
	t.Run("Good input", func(t *testing.T) {
		if a, r := 25, getNumber("fdasf2lkjkljklj5ooonaa"); a != r {
			t.Fatalf("Expected %v got %v", a, r)
		}
	})
	t.Run("Bad input", func(t *testing.T) {
		if a, r := 0, getNumber("vjheib"); a != r {
			t.Fatalf("Expected %v got %v", a, r)
		}
	})
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
