package main

import (
	"os"
	"testing"
)

func TestSolutions(t *testing.T) {
	t.Run("Problem1 with sample.txt", func(t *testing.T) {
		if a, r := 4361, Problem1("sample.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem1 with input.txt", func(t *testing.T) {
		if a, r := 550934, Problem1("input.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem2 with sample.txt", func(t *testing.T) {
		if a, r := 467835, Problem2("sample.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem2 with input.txt", func(t *testing.T) {
		if a, r := 81997870, Problem2("input.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
