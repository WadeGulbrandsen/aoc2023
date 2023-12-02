package day02_test

import (
	"testing"

	"github.com/WadeGulbrandsen/aoc2023/day02"
)

func TestDay02Problem01(t *testing.T) {
	if a, r := 8, day02.Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}

func TestDay02Problem02(t *testing.T) {
	if a, r := 2286, day02.Problem2("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}
