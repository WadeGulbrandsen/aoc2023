package day03_test

import (
	"testing"

	"github.com/WadeGulbrandsen/aoc2023/day03"
)

func TestDay03Problem01(t *testing.T) {
	if a, r := 4361, day03.Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}
func TestDay03Problem02(t *testing.T) {
	if a, r := 467835, day03.Problem2("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}
