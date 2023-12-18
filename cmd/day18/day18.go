package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 18

type digInstruction struct {
	direction internal.GridDirection
	distance  int
	color     string
}

type digPlan []digInstruction

func (d digPlan) Dig() internal.Grid {
	min_x, min_y, max_x, max_y := 0, 0, 0, 0
	current := internal.GridPoint{X: 0, Y: 0}
	holes := map[internal.GridPoint]rune{current: '#'}
	lefts := make(map[internal.GridPoint]rune)
	rights := make(map[internal.GridPoint]rune)
	for _, inst := range d {
		for i := 0; i < inst.distance; i++ {
			current = current.Move(inst.direction)
			min_x = min(min_x, current.X)
			max_x = max(max_x, current.X)
			min_y = min(min_y, current.Y)
			max_y = max(max_y, current.Y)
			holes[current] = '#'
			lefts[current.Move(inst.direction.TurnL())] = 'L'
			rights[current.Move(inst.direction.TurnR())] = 'R'
		}
	}
	minPoint := internal.GridPoint{X: min_x, Y: min_y}
	maxPoint := internal.GridPoint{X: max_x + 1, Y: max_y + 1}
	grid := internal.Grid{MinPoint: minPoint, MaxPoint: maxPoint, Points: holes}
	// fmt.Printf("Dug perimiter\n\n%v\n\n", grid)
	inner := &lefts
	for p := range lefts {
		if !grid.InBounds(p) {
			inner = &rights
			break
		}
	}
	for p, r := range *inner {
		if _, ok := grid.Points[p]; !ok {
			grid.Points[p] = r
		}
	}
	return grid
}

func strToDigInstruction(s string) (digInstruction, bool) {
	parts := strings.Split(s, " ")
	inst := digInstruction{}
	if len(parts) != 3 {
		return inst, false
	}
	inst.color = parts[2]
	switch parts[0] {
	case "U":
		inst.direction = internal.N
	case "D":
		inst.direction = internal.S
	case "L":
		inst.direction = internal.W
	case "R":
		inst.direction = internal.E
	default:
		return inst, false
	}
	if n, e := strconv.Atoi(parts[1]); e == nil {
		inst.distance = n
		return inst, true
	}
	return inst, false
}

func GetDigPlan(data *[]string) digPlan {
	var plan digPlan
	for _, s := range *data {
		if di, found := strToDigInstruction(s); found {
			plan = append(plan, di)
		}
	}
	return plan
}

func Problem1(data *[]string) int {
	plan := GetDigPlan(data)
	grid := plan.Dig()
	fmt.Println(grid)
	return len(grid.Points)
}

func Problem2(data *[]string) int {
	return 0
}

func main() {
	internal.RunSolutions(Day, Problem1, Problem2, "sample.txt", "sample.txt", -1)
}
