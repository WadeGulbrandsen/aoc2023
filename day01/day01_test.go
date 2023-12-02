package day01_test

import (
	"testing"

	"github.com/WadeGulbrandsen/aoc2023/day01"
)

func TestDay01Problem01(t *testing.T) {
	if a, r := 142, day01.Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}
func TestDay01Problem02(t *testing.T) {
	if a, r := 281, day01.Problem2("sample2.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}
