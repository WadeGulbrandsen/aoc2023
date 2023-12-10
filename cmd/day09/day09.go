package main

import (
	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 9

func getNextStep(step *[]int) []int {
	var next []int
	for i := 1; i < len(*step); i++ {
		next = append(next, (*step)[i]-(*step)[i-1])
	}
	return next
}

func isZero(x int) bool {
	return x == 0
}

func nextItemInHistory(s string) int {
	history := internal.GetIntsFromString(s)
	steps := [][]int{history}
	for {
		next := getNextStep(&steps[len(steps)-1])
		if internal.All(&next, isZero) {
			next = append(next, 0)
			steps = append(steps, next)
			break
		}
		steps = append(steps, next)
	}
	for i := len(steps) - 2; i >= 0; i-- {
		x, y := internal.Last(&steps[i]), internal.Last(&steps[i+1])
		steps[i] = append(steps[i], x+y)
	}
	return internal.Last(&steps[0])
}

func prevItemInHistory(s string) int {
	history := internal.GetIntsFromString(s)
	steps := [][]int{history}
	for {
		next := getNextStep(&steps[len(steps)-1])
		if internal.All(&next, isZero) {
			next = append(next, 0)
			steps = append(steps, next)
			break
		}
		steps = append(steps, next)
	}
	for i := len(steps) - 2; i >= 0; i-- {
		x, y := steps[i][0], steps[i+1][0]
		steps[i] = append([]int{x - y}, steps[i]...)
	}
	return steps[0][0]
}

func Problem1(data *[]string) int {
	return internal.SumSolver(data, nextItemInHistory)
}

func Problem2(data *[]string) int {
	return internal.SumSolver(data, prevItemInHistory)
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
