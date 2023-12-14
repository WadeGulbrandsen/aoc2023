package internal

type GridPoint struct {
	X, Y int
}

func (p *GridPoint) Distance(o *GridPoint) int {
	dx := AbsDiff[int](p.X, o.X)
	dy := AbsDiff[int](p.Y, o.Y)
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

func (g *Grid) Rows() []string {
	lines := make([]string, g.Size.Y)
	for y := 0; y < g.Size.Y; y++ {
		line := make([]rune, g.Size.X)
		for x := 0; x < g.Size.X; x++ {
			if c := g.At(GridPoint{X: x, Y: y}); c != 0 {
				line[x] = c
			} else {
				line[x] = '.'
			}
		}
		lines[y] = string(line)
	}
	return lines
}

func (g *Grid) Columns() []string {
	lines := make([]string, g.Size.X)
	for x := 0; x < g.Size.X; x++ {
		line := make([]rune, g.Size.Y)
		for y := 0; y < g.Size.Y; y++ {
			if c := g.At(GridPoint{X: x, Y: y}); c != 0 {
				line[y] = c
			} else {
				line[y] = '.'
			}
		}
		lines[x] = string(line)
	}
	return lines
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
