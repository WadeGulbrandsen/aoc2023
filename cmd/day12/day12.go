package main

import (
	"slices"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 12

type groupType string

const (
	undefined groupType = ""
	working   groupType = "."
	broken    groupType = "#"
	unknown   groupType = "?"
	mixed     groupType = "?#"
)

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
		case (kind == broken || kind == unknown) && strings.ContainsRune(string(mixed), r):
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

func groupSizes(groups *[]group) []int {
	var sizes []int
	for _, g := range *groups {
		switch g.kind {
		case broken, unknown, mixed:
			sizes = append(sizes, len(g.contents))
		}
	}
	return sizes
}

func findArrangements(groups []group, sizes []int, prev string) []string {
	if len(groups) == 0 {
		if len(sizes) == 0 {
			return []string{prev}
		}
		return nil
	}
	groups_head, groups_tail := groups[0], groups[1:]
	switch groups_head.kind {
	case working:
		return findArrangements(groups_tail, sizes, prev+string(groups_head.contents))
	case broken:
		sizes_head, sizes_tail := sizes[0], sizes[1:]
		if len(groups_head.contents) == sizes_head {
			return findArrangements(groups_tail, sizes_tail, prev+string(groups_head.contents))
		} else {
			return nil
		}
	case unknown, mixed:
		sizes_head, sizes_tail := sizes[0], sizes[1:]
		space := len(groups_head.contents)
		if space == sizes_head {
			return findArrangements(groups_tail, sizes_tail, prev+strings.Repeat("#", sizes_head))
		}
		if space < sizes_head {
			return nil
		}
		var results []string
		groups_tail_sizes := groupSizes(&groups_tail)
		if groups_head.kind == unknown {
			for i := 0; i <= space-sizes_head; i++ {
				results = append(results, findArrangements(groups_tail, sizes_tail, prev+strings.Repeat(".", i)+strings.Repeat("#", sizes_head)+strings.Repeat(".", space-sizes_head-i))...)
				if len(sizes_tail) > len(groups_tail_sizes) && internal.Sum(&sizes_tail) < internal.Sum(&groups_tail_sizes) && space-sizes_head-i-1 >= sizes_tail[0] {
					results = append(results, findArrangements(groups_tail[1:], sizes_tail[1:], prev+strings.Repeat(".", i)+strings.Repeat("#", sizes_head)+"."+strings.Repeat("#", space-sizes_head-i-1))...)
				}
			}
		} else {
			for i := 0; i <= space-sizes_head; i++ {
				if i > 0 && slices.Contains(groups_head.contents[0:i+1], '#') {
					break
				}
				remaining := space - sizes_head - i
				if remaining > 0 && groups_head.contents[i+sizes_head] == '#' {
					continue
				}
				switch remaining {
				case 0:
					results = append(results, findArrangements(groups_tail, sizes_tail, prev+strings.Repeat(".", i)+strings.Repeat("#", sizes_head))...)
				case 1:
					results = append(results, findArrangements(groups_tail, sizes_tail, prev+strings.Repeat(".", i)+strings.Repeat("#", sizes_head)+".")...)
				default:
					rest_groups := stringToGroups(string(groups_head.contents[i+sizes_head+1:]))
					results = append(results, findArrangements(append(groups_tail, rest_groups...), sizes_tail, prev+strings.Repeat(".", i)+strings.Repeat("#", sizes_head)+".")...)
				}
			}
		}
		return results
	}
	return nil
}

func countArrangements(s string) int {
	before, after, found := strings.Cut(s, " ")
	if !found {
		return 0
	}
	sizes := internal.GetIntsFromString(after, ",")
	groups := stringToGroups(before)
	results := findArrangements(groups, sizes, "")
	return len(results)
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
