package main

import (
	"os"
	"testing"
)

func TestDay01(t *testing.T) {
	t.Run("Problem1 with sample.txt", func(t *testing.T) {
		if a, r := 142, Problem1("sample.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem1 with input.txt", func(t *testing.T) {
		if a, r := 54708, Problem1("input.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem2 with sample2.txt", func(t *testing.T) {
		if a, r := 281, Problem2("sample2.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
	t.Run("Problem2 with input.txt", func(t *testing.T) {
		if a, r := 54087, Problem2("input.txt"); a != r {
			t.Fatalf("The correct answer is %v but received %v", a, r)
		}
	})
}

func TestGetNumber(t *testing.T) {
	t.Run("Good input", func(t *testing.T) {
		if a, r := 25, getNumber("fdasf2lkjkljklj5ooonaa"); a != r {
			t.Fatalf("Expected %v got %v", a, r)
		}
	})
	t.Run("Bad input", func(t *testing.T) {
		if a, r := 0, getNumber("vjheib"); a != r {
			t.Fatalf("Expected %v got %v", a, r)
		}
	})
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
