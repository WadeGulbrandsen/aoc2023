package main

import (
	"os"
	"testing"
)

func TestSolutions(t *testing.T) {
	if a, r := 8, Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 2348, Problem1("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 2286, Problem2("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 76008, Problem2("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
