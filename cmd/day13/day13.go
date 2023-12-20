package main

import (
	"github.com/WadeGulbrandsen/aoc2023/internal/functional"
	"github.com/WadeGulbrandsen/aoc2023/internal/grid"
	"github.com/WadeGulbrandsen/aoc2023/internal/solve"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 13

func stringDiff(a string, b string) int {
	ar, br := []rune(a), []rune(b)
	s, l := min(len(ar), len(br)), max(len(ar), len(br))
	diffs := l - s
	for x := 0; x < s; x++ {
		if ar[x] != br[x] {
			diffs++
		}
	}
	return diffs
}

func findReflection(p []string) int {
MAIN:
	for i := 1; i < len(p); i++ {
		c := min(i, len(p)-i)
		for j := 0; j < c; j++ {
			if p[i-(1+j)] != p[i+j] {
				continue MAIN
			}
		}
		return i
	}
	return 0
}

func findReflections(p []string) int {
	g := grid.MakeGridFromLines(&p)
	rows := g.Rows()
	if result := findReflection(rows); result != 0 {
		return result * 100
	}
	cols := g.Columns()
	return findReflection(cols)
}

func findSmudge(p []string) int {
MAIN:
	for i := 1; i < len(p); i++ {
		c := min(i, len(p)-i)
		d := 0
		for j := 0; j < c; j++ {
			d += stringDiff(p[i-(1+j)], p[i+j])
			if d > 1 {
				continue MAIN
			}
		}
		if d == 1 {
			return i
		}
	}
	return 0
}

func fixSmudges(p []string) int {
	g := grid.MakeGridFromLines(&p)
	rows := g.Rows()
	if result := findSmudge(rows); result != 0 {
		return result * 100
	}
	cols := g.Columns()
	return findSmudge(cols)
}

func Problem1(data *[]string) int {
	patterns := functional.Split(data, "")
	return solve.SumSolver(&patterns, findReflections)
}

func Problem2(data *[]string) int {
	patterns := functional.Split(data, "")
	return solve.SumSolver(&patterns, fixSmudges)
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
