package internal

import "testing"

func stringLength(s string, ch chan int) {
	ch <- len(s)
}

func TestSolvers(t *testing.T) {
	t.Run("No file", func(t *testing.T) {
		if r := SumSolver("nofile.txt", stringLength); r != 0 {
			t.Fatalf("Non existing file should return 0 but got %v", r)
		}
	})
	t.Run("Valid file", func(t *testing.T) {
		if r := SumSolver("testdata.txt", stringLength); r != 17 {
			t.Fatalf("The sum should be 17 but got %v", r)
		}
	})
}
