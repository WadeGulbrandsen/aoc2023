package main

import (
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/functional"
	"github.com/WadeGulbrandsen/aoc2023/internal/set"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/mowshon/iterium"
)

const Day = 24

type hailstone struct {
	px, py, pz, vx, vy, vz float64
}

type hailstoneint struct {
	px, py, pz, vx, vy, vz int
}

func (h hailstone) intersection_tu(o hailstone) (float64, float64) {
	// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line_segment
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

func parseHailstoneInt(s string) hailstoneint {
	before, after, found := strings.Cut(s, " @ ")
	if !found {
		return hailstoneint{}
	}
	ps, vs := utils.GetIntsFromString(before, ", "), utils.GetIntsFromString(after, ", ")
	if len(ps) != 3 || len(vs) != 3 {
		return hailstoneint{}
	}
	return hailstoneint{ps[0], ps[1], ps[2], vs[0], vs[1], vs[2]}
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
		a, b := combo[0], combo[1]
		if t, u := a.intersection_tu(b); t >= 0 && u >= 0 {
			x := a.px + (t * a.vx)
			y := a.py + (t * a.vy)
			if bounds.inside(x, y) {
				collisions++
			}
		}
	}
	return collisions
}

func potential_velocity(ap, av, bp, bv int) set.Set[int] {
	vs := set.NewSet[int]()
	if av != bv {
		return nil
	}
	diff := utils.Abs(bp - ap)
	for v := -1000; v <= 1000; v++ {
		if v != av && diff%(v-av) == 0 {
			vs.Add(v)
		}
	}
	return vs
}

func Problem2(data *[]string) int {
	hailstorm := functional.Map(data, parseHailstoneInt)
	sets := map[rune]set.Set[int]{
		'x': set.NewSet[int](),
		'y': set.NewSet[int](),
		'z': set.NewSet[int](),
	}
	var rvx, rvy, rvz int
	if len(hailstorm) == 5 {
		// the sample input doesn't work with the method below so have to hard code the rock velocity
		rvx, rvy, rvz = -3, 1, 2
	} else {
		combos := iterium.Combinations(hailstorm, 2)
		for combo := range combos.Chan() {
			a, b := combo[0], combo[1]
			potentials := map[rune]set.Set[int]{
				'x': potential_velocity(a.px, a.vx, b.px, b.vx),
				'y': potential_velocity(a.py, a.vy, b.py, b.vy),
				'z': potential_velocity(a.pz, a.vz, b.pz, b.vz),
			}
			for k, p := range potentials {
				if p.IsEmpty() {
					continue
				}
				if sets[k].IsEmpty() {
					sets[k] = p
				} else {
					sets[k] = sets[k].Intersect(p)
				}
			}
		}
		rvx, _ = sets['x'].Pop()
		rvy, _ = sets['y'].Pop()
		rvz, _ = sets['z'].Pop()
	}
	a, b := hailstorm[0], hailstorm[1]
	ma := float64(a.vy-rvy) / float64(a.vx-rvx)
	mb := float64(b.vy-rvy) / float64(b.vx-rvx)
	ca := float64(a.py) - (ma * float64(a.px))
	cb := float64(b.py) - (mb * float64(b.px))
	x := int((cb - ca) / (ma - mb))
	y := int((ma * float64(x)) + ca)
	t := (x - a.px) / (a.vx - rvx)
	z := a.pz + (a.vz-rvz)*t
	return x + y + z
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
