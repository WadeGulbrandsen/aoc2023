package grid

import (
	"slices"

	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

type GridDirection rune

const (
	N GridDirection = 'N'
	E GridDirection = 'E'
	S GridDirection = 'S'
	W GridDirection = 'W'
)

func (d GridDirection) TurnL() GridDirection {
	switch d {
	case N:
		return W
	case E:
		return N
	case S:
		return E
	case W:
		return S
	}
	return d
}

func (d GridDirection) TurnR() GridDirection {
	switch d {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	}
	return d
}

func (d GridDirection) Reverse() GridDirection {
	switch d {
	case N:
		return S
	case E:
		return W
	case S:
		return N
	case W:
		return E
	}
	return d
}

type GridVector struct {
	Point     GridPoint
	Direction GridDirection
}

func (v *GridVector) TurnL() GridVector {
	return GridVector{v.Point, v.Direction.TurnL()}
}

func (v *GridVector) TurnR() GridVector {
	return GridVector{v.Point, v.Direction.TurnR()}
}

func (v *GridVector) Reverse() GridVector {
	return GridVector{v.Point, v.Direction.Reverse()}
}

func (v *GridVector) Next() GridVector {
	return GridVector{v.Point.Move(v.Direction, 1), v.Direction}
}

func (v *GridVector) MoveL() GridVector {
	return GridVector{v.Point.Move(v.Direction.TurnL(), 1), v.Direction.TurnL()}
}

func (v *GridVector) MoveR() GridVector {
	return GridVector{v.Point.Move(v.Direction.TurnR(), 1), v.Direction.TurnR()}
}

type GridPoint struct {
	X, Y int
}

func (p *GridPoint) Move(direction GridDirection, distance int) GridPoint {
	switch direction {
	case N:
		return GridPoint{p.X, p.Y - distance}
	case E:
		return GridPoint{p.X + distance, p.Y}
	case S:
		return GridPoint{p.X, p.Y + distance}
	case W:
		return GridPoint{p.X - distance, p.Y}
	}
	return *p
}

func (p *GridPoint) Distance(o *GridPoint) int {
	dx := utils.AbsDiff[int](p.X, o.X)
	dy := utils.AbsDiff[int](p.Y, o.Y)
	return dx + dy
}

func (p *GridPoint) N() GridPoint {
	return GridPoint{p.X, p.Y - 1}
}

func (p *GridPoint) S() GridPoint {
	return GridPoint{p.X, p.Y + 1}
}

func (p *GridPoint) W() GridPoint {
	return GridPoint{p.X - 1, p.Y}
}

func (p *GridPoint) E() GridPoint {
	return GridPoint{p.X + 1, p.Y}
}

type Grid struct {
	MinPoint GridPoint
	MaxPoint GridPoint
	Points   map[GridPoint]rune
}

func (g *Grid) InBounds(p GridPoint) bool {
	return p.X >= g.MinPoint.X && p.X < g.MaxPoint.X && p.Y >= g.MinPoint.Y && p.Y < g.MaxPoint.Y
}

func (g *Grid) At(p GridPoint) rune {
	return g.Points[p]
}

func (g *Grid) Fill(p GridPoint, r rune) {
	to_replace := g.At(p)
	to_check := []GridPoint{p}
	checked := make(map[GridPoint]bool)
	for len(to_check) > 0 {
		point := to_check[0]
		to_check = to_check[1:]
		if checked[point] {
			continue
		}
		checked[point] = true
		if g.At(point) == to_replace {
			g.Points[point] = r
			for _, next := range [4]GridPoint{point.N(), point.E(), point.S(), point.W()} {
				if !checked[next] && g.At(next) == to_replace && g.InBounds(next) && !slices.Contains(to_check, next) {
					to_check = append(to_check, next)
				}
			}
		}
	}
}

func (g Grid) String() string {
	s := ""
	for y := g.MinPoint.Y; y < g.MaxPoint.Y; y++ {
		for x := g.MinPoint.X; x < g.MaxPoint.X; x++ {
			r := g.At(GridPoint{x, y})
			if r == 0 {
				r = '.'
			}
			s += string(r)
		}
		s += "\n"
	}
	return s
}

func (g *Grid) Rows() []string {
	var lines []string
	for y := g.MinPoint.Y; y < g.MaxPoint.Y; y++ {
		line := ""
		for x := g.MinPoint.X; x < g.MaxPoint.X; x++ {
			if c := g.At(GridPoint{X: x, Y: y}); c != 0 {
				line += string(c)
			} else {
				line += "."
			}
		}
		lines = append(lines, line)
	}
	return lines
}

func (g *Grid) Columns() []string {
	var lines []string
	for x := g.MinPoint.X; x < g.MaxPoint.X; x++ {
		line := ""
		for y := g.MinPoint.Y; y < g.MaxPoint.Y; y++ {
			if c := g.At(GridPoint{X: x, Y: y}); c != 0 {
				line += string(c)
			} else {
				line += "."
			}
		}
		lines = append(lines, line)
	}
	return lines
}

func MakeGridFromLines(lines *[]string) Grid {
	g := Grid{Points: make(map[GridPoint]rune)}
	g.MaxPoint.Y = len(*lines)
	if g.MaxPoint.Y > 0 {
		g.MaxPoint.X = len((*lines)[0])
		for y, l := range *lines {
			for x, r := range l {
				if r != '.' {
					g.Points[GridPoint{x, y}] = r
				}
			}
		}
	}
	return g
}

func ShoelaceArea(path []GridPoint) int {
	a, b, perimiter := 0, 0, 0
	for i, p := range path[1:] {
		a += path[i].X * p.Y
		b += path[i].Y * p.X
		perimiter += p.Distance(&path[i])
	}
	area := utils.AbsDiff(a, b) / 2
	return 1 + area + perimiter/2
}
