package main

import (
	"testing"
)

func TestSolutions(t *testing.T) {
	if a, r := 4361, Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 550934, Problem1("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 467835, Problem2("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 81997870, Problem2("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
}
