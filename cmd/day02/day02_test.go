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
const ans1sample = 8
const ans1input = 2348
const ans2sample = 2286
const ans2input = 76008

// filenames
const file1sample = "sample.txt"
const file1input = "input.txt"
const file2sample = file1sample
const file2input = file1input

func TestSolutions(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
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

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
