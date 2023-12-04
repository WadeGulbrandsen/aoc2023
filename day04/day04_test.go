package day04_test

import (
	"testing"

	"github.com/WadeGulbrandsen/aoc2023/day04"
)

func TestDay04Problem1(t *testing.T) {
	if a, r := 13, day04.Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}

func TestDay04Problem2(t *testing.T) {
	if a, r := 30, day04.Problem2("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}
