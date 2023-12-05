package main

import (
	"os"
	"testing"
)

func TestSolutions(t *testing.T) {
	if a, r := 13, Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 27454, Problem1("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 30, Problem2("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 6857330, Problem2("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
