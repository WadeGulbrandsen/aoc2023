package main

import (
	"cmp"
	"container/heap"
	"slices"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/heaps"
	"github.com/WadeGulbrandsen/aoc2023/internal/set"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 22

type Point3D struct {
	X, Y, Z int
}

type Brick struct {
	ID                 int
	MinPoint, MaxPoint Point3D
	Above, Below       []*Brick
}

func (b *Brick) Overlap(a *Brick) bool {
	if a.MinPoint.Y > b.MaxPoint.Y || b.MinPoint.Y > a.MaxPoint.Y || a.MinPoint.X > b.MaxPoint.X || b.MinPoint.X > a.MaxPoint.X {
		return false
	}
	return true
}

func CmpPoint3D(a, b Point3D) int {
	if n := cmp.Compare(a.Z, b.Z); n != 0 {
		return n
	}
	if n := cmp.Compare(a.Y, b.Y); n != 0 {
		return n
	}
	return cmp.Compare(a.X, b.X)
}

func CmpBrickBottoms(a, b Brick) int {
	if n := CmpPoint3D(a.MinPoint, b.MinPoint); n != 0 {
		return n
	}
	return CmpPoint3D(a.MaxPoint, b.MaxPoint)
}

func CmpBrickTops(a, b Brick) int {
	if n := CmpPoint3D(a.MaxPoint, b.MaxPoint); n != 0 {
		return n
	}
	return CmpPoint3D(a.MinPoint, b.MinPoint)
}

func CmpBrickPointersTops(a, b *Brick) int {
	return CmpBrickTops(*a, *b)
}

func parsePoint3D(s string) Point3D {
	ints := utils.GetIntsFromString(s, ",")
	if len(ints) != 3 {
		return Point3D{}
	}
	return Point3D{X: ints[0], Y: ints[1], Z: ints[2]}
}

func parseBrick(s string, id int) Brick {
	before, after, found := strings.Cut(s, "~")
	if !found {
		return Brick{}
	}
	p1, p2, empty := parsePoint3D(before), parsePoint3D(after), Point3D{}
	if p1 == empty || p2 == empty {
		return Brick{}
	}
	return Brick{MinPoint: p1, MaxPoint: p2, ID: id}
}

func findBricksBelow(bricks *[]Brick, brick *Brick) []*Brick {
	var below []*Brick
	for i, b := range *bricks {
		if b.MaxPoint == brick.MaxPoint && b.MinPoint == brick.MinPoint {
			continue
		}
		if b.MaxPoint.Z < brick.MinPoint.Z && brick.Overlap(&b) {
			below = append(below, &(*bricks)[i])
		}
	}
	if len(below) != 0 {
		top := slices.MaxFunc(below, CmpBrickPointersTops)
		below = slices.DeleteFunc(below, func(b *Brick) bool {
			return b.MaxPoint.Z < top.MaxPoint.Z
		})
	}
	return below
}

func (b *Brick) SafeToDisintegrate() bool {
	for _, above := range b.Above {
		if len(above.Below) <= 1 {
			return false
		}
	}
	return true
}

func dropIt(bricks []Brick) ([]Brick, int) {
	slices.SortFunc(bricks, CmpBrickBottoms)
	for i := range bricks {
		if bricks[i].MinPoint.Z == 1 {
			continue
		}
		below := findBricksBelow(&bricks, &bricks[i])
		if len(below) == 0 {
			bricks[i].MaxPoint.Z -= bricks[i].MinPoint.Z - 1
			bricks[i].MinPoint.Z = 1
		} else {
			top := below[0]
			bricks[i].MaxPoint.Z -= bricks[i].MinPoint.Z - top.MaxPoint.Z - 1
			bricks[i].MinPoint.Z = top.MaxPoint.Z + 1
			for _, u := range below {
				bricks[i].Below = append(bricks[i].Below, u)
				u.Above = append(u.Above, &bricks[i])
			}
		}
	}
	return bricks, 0
}

func getBrickStack(data *[]string) []Brick {
	var bricks []Brick
	for i, s := range *data {
		bricks = append(bricks, parseBrick(s, i+1))
	}
	bricks, _ = dropIt(bricks)
	return bricks
}

func Problem1(data *[]string) int {
	bricks := getBrickStack(data)
	disintegration := 0
	for _, b := range bricks {
		if b.SafeToDisintegrate() {
			disintegration++
		}
	}
	return disintegration
}

func (b Brick) willFall(disintegrated *set.Set[int]) bool {
	for _, below := range b.Below {
		if !disintegrated.Contains(below.ID) {
			return false
		}
	}
	return true
}

func Disintegrate(start_id int, bricks *map[int]Brick) set.Set[int] {
	disintegrated := set.NewSet[int]()
	q := &heaps.PriorityQueue[int]{}
	heap.Init(q)
	heap.Push(q, &heaps.Item[int]{Value: start_id})
	seen := make(map[int]bool)
	for q.Len() > 0 {
		id := heap.Pop(q).(*heaps.Item[int]).Value
		if seen[id] {
			continue
		}
		seen[id] = true
		if id == start_id || (*bricks)[id].willFall(&disintegrated) {
			disintegrated.Add(id)
			for _, above := range (*bricks)[id].Above {
				heap.Push(q, &heaps.Item[int]{Value: above.ID, Priority: -above.MinPoint.Z})
			}
		}
	}
	return disintegrated
}

func Problem2(data *[]string) int {
	bricks := make(map[int]Brick)
	for _, b := range getBrickStack(data) {
		bricks[b.ID] = b
	}
	sum := 0
	for i, b := range bricks {
		if !b.SafeToDisintegrate() {
			results := Disintegrate(i, &bricks)
			sum += len(results) - 1
		}
	}
	return sum
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
