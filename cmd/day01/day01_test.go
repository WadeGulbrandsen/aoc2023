package main

import "testing"

func TestDay01Problem01(t *testing.T) {
	if a, r := 142, Problem1("sample.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 54708, Problem1("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
}

func TestDay01Problem02(t *testing.T) {
	if a, r := 281, Problem2("sample2.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
	if a, r := 54087, Problem2("input.txt"); a != r {
		t.Fatalf("The correct answer is %v but received %v", a, r)
	}
}
