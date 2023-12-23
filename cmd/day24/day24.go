package main

import (
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/functional"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/mowshon/iterium"
)

const Day = 24

type hailstone struct {
	px, py, pz, vx, vy, vz float64
}

func (h hailstone) intersection_time(o hailstone) (float64, float64) {
	x1, x2, x3, x4 := h.px, h.px+h.vx, o.px, o.px+o.vx
	y1, y2, y3, y4 := h.py, h.py+h.vy, o.py, o.py+o.vy
	denominator := ((x1 - x2) * (y3 - y4)) - ((y1 - y2) * (x3 - x4))
	if denominator == 0 {
		return -1, -1
	}
	t := (((x1 - x3) * (y3 - y4)) - ((y1 - y3) * (x3 - x4))) / denominator
	u := (((x1 - x3) * (y1 - y2)) - ((y1 - y3) * (x1 - x2))) / denominator
	return t, u
}

type testarea struct {
	low_x, low_y, high_x, high_y float64
}

func (ta *testarea) inside(x, y float64) bool {
	result := x >= ta.low_x && x <= ta.high_x && y >= ta.low_y && y <= ta.high_y
	return result
}

func parseHailstone(s string) hailstone {
	before, after, found := strings.Cut(s, " @ ")
	if !found {
		return hailstone{}
	}
	ps, vs := utils.GetFloatsFromString(before, ", "), utils.GetFloatsFromString(after, ", ")
	if len(ps) != 3 || len(vs) != 3 {
		return hailstone{}
	}
	return hailstone{ps[0], ps[1], ps[2], vs[0], vs[1], vs[2]}
}

func Problem1(data *[]string) int {
	hailstorm := functional.Map(data, parseHailstone)
	var bounds testarea
	if len(hailstorm) == 5 {
		bounds = testarea{7, 7, 27, 27}
	} else {
		bounds = testarea{200000000000000, 200000000000000, 400000000000000, 400000000000000}
	}
	combos := iterium.Combinations(hailstorm, 2)
	collisions := 0
	for combo := range combos.Chan() {
		if t, u := combo[0].intersection_time(combo[1]); t >= 0 && u >= 0 {
			x := combo[0].px + (t * combo[0].vx)
			y := combo[0].py + (t * combo[0].vy)
			if bounds.inside(x, y) {
				collisions++
			}
		}
	}
	return collisions
}

func Problem2(data *[]string) int {
	return 0
}

func main() {
	utils.RunSolutions(Day, Problem1, Problem2, "input.txt", "sample.txt", -1)
}
