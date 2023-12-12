package internal

type GridPoint struct {
	X, Y int
}

func (p *GridPoint) Distance(o *GridPoint) int {
	dx := AbsDiff[int](p.X, o.X)
	dy := AbsDiff[int](p.Y, o.Y)
	return dx + dy
}

type Grid struct {
	Size   GridPoint
	Points map[GridPoint]rune
}

func (g *Grid) At(p GridPoint) rune {
	return g.Points[p]
}

func (g Grid) String() string {
	s := ""
	for y := 0; y < g.Size.Y; y++ {
		for x := 0; x < g.Size.X; x++ {
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

func MakeGridFromLines(lines *[]string) Grid {
	g := Grid{Points: make(map[GridPoint]rune)}
	g.Size.Y = len(*lines)
	if g.Size.Y > 0 {
		g.Size.X = len((*lines)[0])
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
