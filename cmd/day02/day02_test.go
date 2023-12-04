package main

import (
	"testing"
)

func TestDay02Problem01(t *testing.T) {
	if a, r := 8, Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
	if a, r := 2348, Problem1("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}

func TestDay02Problem02(t *testing.T) {
	if a, r := 2286, Problem2("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
	if a, r := 76008, Problem2("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but recieved %v", a, r)
	}
}
