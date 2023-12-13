package main

import (
	"slices"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 12

type groupType string

const (
	undefined = ""
	working   = "."
	broken    = "#"
	unknown   = "?"
	mixed     = "?#"
)

type arrangements struct {
	sizes  []int
	groups []group
	left   *arrangements
	right  *arrangements
}

func (a *arrangements) solve() {
	if len(a.groups) > 1 {
		a.split()
	}
}

func (a *arrangements) split() {
	s := 0
LEFT:
	for i, g := range a.groups {
		switch g.kind {
		case broken:
			s++
		case unknown, mixed:
			a.left = &arrangements{sizes: a.sizes[0:s], groups: a.groups[0:i]}
			a.sizes = a.sizes[s:]
			a.groups = a.groups[i:]
			break LEFT
		}
	}
	s = len(a.sizes)
RIGHT:
	for i := len(a.groups); i >= 0; i-- {
		switch g := a.groups[i-1]; g.kind {
		case broken:
			s--
		case unknown, mixed:
			a.right = &arrangements{sizes: a.sizes[s:], groups: a.groups[i:]}
			a.sizes = a.sizes[0:s]
			a.groups = a.groups[0:i]
			break RIGHT
		}
	}
}

type group struct {
	kind     groupType
	contents []rune
}

func stringToGroups(s string) []group {
	var groups []group
	var current []rune
	var kind groupType
	new_group := func(r rune) ([]rune, groupType) {
		var k groupType
		switch r {
		case '.':
			k = working
		case '#':
			k = broken
		case '?':
			k = unknown
		}
		return []rune{r}, k
	}
	for _, r := range s {
		switch {
		case len(current) == 0:
			current, kind = new_group(r)
		case strings.ContainsRune(string(kind), r):
			current = append(current, r)
		case (kind == broken || kind == unknown) && strings.ContainsRune(mixed, r):
			current = append(current, r)
			kind = mixed
		default:
			groups = append(groups, group{kind: kind, contents: current})
			current, kind = new_group(r)
		}
	}
	groups = append(groups, group{kind: kind, contents: current})
	return groups
}

func getLen[T string | []any](x T) int {
	return len(x)
}

func isValidArrangement(arrangement *string, sizes *[]int) bool {
	as := strings.Split(*arrangement, ".")
	arrangement_sizes := internal.Map(&as, getLen)
	arrangement_sizes = slices.DeleteFunc(arrangement_sizes, func(i int) bool { return i == 0 })
	// slices.Sort(arrangement_sizes)
	return slices.Equal(arrangement_sizes, *sizes)
}

func allBroken(s string) bool {
	runes := []rune(s)
	return internal.All(&runes, func(r rune) bool { return r == '#' })
}

func findArrangements(s string) []string {
	var results []string
	before, after, found := strings.Cut(s, " ")
	if !found {
		return results
	}
	sizes := internal.GetIntsFromString(after, ",")
	groups := stringToGroups(before)
	arrangements := arrangements{sizes: sizes, groups: groups}
	arrangements.solve()
	return results
}

func countArrangements(s string) int {
	arrangements := findArrangements(s)
	return len(arrangements)
}

func Problem1(data *[]string) int {
	results := internal.Map(data, countArrangements)
	return internal.Sum(&results)
}

func Problem2(data *[]string) int {
	return 0
}

func main() {
	internal.RunSolutions(Day, Problem1, Problem2, "sample.txt", "sample.txt", -1)
}
