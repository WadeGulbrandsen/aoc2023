package grid

import (
	"slices"

	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

type Direction rune

const (
	N Direction = 'N'
	E Direction = 'E'
	S Direction = 'S'
	W Direction = 'W'
)

func (d Direction) TurnL() Direction {
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

func (d Direction) TurnR() Direction {
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

func (d Direction) Reverse() Direction {
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

type Vector struct {
	Point     Point
	Direction Direction
}

func (v *Vector) TurnL() Vector {
	return Vector{v.Point, v.Direction.TurnL()}
}

func (v *Vector) TurnR() Vector {
	return Vector{v.Point, v.Direction.TurnR()}
}

func (v *Vector) Reverse() Vector {
	return Vector{v.Point, v.Direction.Reverse()}
}

func (v *Vector) Next() Vector {
	return Vector{v.Point.Move(v.Direction, 1), v.Direction}
}

func (v *Vector) MoveL() Vector {
	return Vector{v.Point.Move(v.Direction.TurnL(), 1), v.Direction.TurnL()}
}

func (v *Vector) MoveR() Vector {
	return Vector{v.Point.Move(v.Direction.TurnR(), 1), v.Direction.TurnR()}
}

type ChessColor int

const (
	White ChessColor = 0
	Black ChessColor = 1
)

type Point struct {
	X, Y int
}

func (p Point) ChessColor() ChessColor {
	return ChessColor((p.X + p.Y) % 2)
}

func (p *Point) Move(direction Direction, distance int) Point {
	switch direction {
	case N:
		return Point{p.X, p.Y - distance}
	case E:
		return Point{p.X + distance, p.Y}
	case S:
		return Point{p.X, p.Y + distance}
	case W:
		return Point{p.X - distance, p.Y}
	}
	return *p
}

func (p *Point) Distance(o *Point) int {
	dx := utils.AbsDiff[int](p.X, o.X)
	dy := utils.AbsDiff[int](p.Y, o.Y)
	return dx + dy
}

func (p *Point) N() Point {
	return Point{p.X, p.Y - 1}
}

func (p *Point) S() Point {
	return Point{p.X, p.Y + 1}
}

func (p *Point) W() Point {
	return Point{p.X - 1, p.Y}
}

func (p *Point) E() Point {
	return Point{p.X + 1, p.Y}
}

type Grid struct {
	MinPoint Point
	MaxPoint Point
	Points   map[Point]rune
}

func (g *Grid) InBounds(p Point) bool {
	return p.X >= g.MinPoint.X && p.X < g.MaxPoint.X && p.Y >= g.MinPoint.Y && p.Y < g.MaxPoint.Y
}

func (g *Grid) At(p Point) rune {
	return g.Points[p]
}

func (g *Grid) Fill(p Point, r rune) {
	to_replace := g.At(p)
	to_check := []Point{p}
	checked := make(map[Point]bool)
	for len(to_check) > 0 {
		point := to_check[0]
		to_check = to_check[1:]
		if checked[point] {
			continue
		}
		checked[point] = true
		if g.At(point) == to_replace {
			g.Points[point] = r
			for _, next := range [4]Point{point.N(), point.E(), point.S(), point.W()} {
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
			r := g.At(Point{x, y})
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
			if c := g.At(Point{X: x, Y: y}); c != 0 {
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
			if c := g.At(Point{X: x, Y: y}); c != 0 {
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
	g := Grid{Points: make(map[Point]rune)}
	g.MaxPoint.Y = len(*lines)
	if g.MaxPoint.Y > 0 {
		g.MaxPoint.X = len((*lines)[0])
		for y, l := range *lines {
			for x, r := range l {
				if r != '.' {
					g.Points[Point{x, y}] = r
				}
			}
		}
	}
	return g
}

func ShoelaceArea(path []Point) int {
	a, b, perimiter := 0, 0, 0
	for i, p := range path[1:] {
		a += path[i].X * p.Y
		b += path[i].Y * p.X
		perimiter += p.Distance(&path[i])
	}
	area := utils.AbsDiff(a, b) / 2
	return 1 + area + perimiter/2
}
