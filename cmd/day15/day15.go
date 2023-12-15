package main

import (
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal"
	"goki.dev/ordmap"
)

const Day = 15

func hash(s string) int {
	current := 0
	for _, r := range s {
		current += int(r)
		current = (current * 17) % 256
	}
	return current
}

func Problem1(data *[]string) int {
	strings := strings.Split(strings.Join(*data, ""), ",")
	return internal.SumSolver(&strings, hash)
}

func Problem2(data *[]string) int {
	var boxes [256]*ordmap.Map[string, int]
	for _, s := range strings.Split(strings.Join(*data, ""), ",") {
		if label, val, found := strings.Cut(s, "="); found {
			if v, err := strconv.Atoi(val); err == nil {
				box := hash(label)
				if boxes[box] == nil {
					boxes[box] = ordmap.New[string, int]()
				}
				boxes[box].Add(label, v)
			}
		} else if label, found := strings.CutSuffix(s, "-"); found {
			box := hash(label)
			if boxes[box] != nil {
				boxes[box].DeleteKey(label)
			}
		}
	}
	sum := 0
	for i, b := range boxes {
		if b != nil {
			for j, lense := range b.Order {
				sum += (i + 1) * (j + 1) * lense.Val
			}
		}
	}
	return sum
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
