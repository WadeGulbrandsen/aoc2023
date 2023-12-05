package main

import (
	"os"
	"testing"
)

func TestSolutions(t *testing.T) {
	t.Run("Problem1 with sample.txt", func(t *testing.T) {
		if a, r := 35, Problem1("sample.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem1 with input.txt", func(t *testing.T) {
		if a, r := 157211394, Problem1("input.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem2 with sample.txt", func(t *testing.T) {
		if a, r := 46, Problem2("sample.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem2 with input.txt", func(t *testing.T) {
		if a, r := 50855035, Problem2("input.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
