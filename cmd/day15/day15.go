package main

import (
	"strconv"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map"

	"github.com/WadeGulbrandsen/aoc2023/internal"
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

type myInt struct {
	val int
}

func Problem2(data *[]string) int {
	var boxes [256]orderedmap.OrderedMap
	for i := range boxes {
		boxes[i] = *orderedmap.New()
	}
	for _, s := range strings.Split(strings.Join(*data, ""), ",") {
		if label, val, found := strings.Cut(s, "="); found {
			if v, err := strconv.Atoi(val); err == nil {
				box := hash(label)
				boxes[box].Set(label, &myInt{v})
			}
		} else if label, found := strings.CutSuffix(s, "-"); found {
			box := hash(label)
			boxes[box].Delete(label)
		}
	}
	sum := 0
	for i, b := range boxes {
		j := 1
		for lense := b.Oldest(); lense != nil; lense = lense.Next() {
			sum += (i + 1) * j * int(lense.Value.(*myInt).val)
			j++
		}
	}
	return sum
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
