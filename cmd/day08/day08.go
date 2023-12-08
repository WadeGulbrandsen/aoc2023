package main

import (
	"regexp"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal"
	"github.com/mowshon/iterium"
	"github.com/rs/zerolog/log"
)

const Day = 8

var nodeMatch = regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

type Node struct {
	name, left, right string
	ends_in_a         bool
	ends_in_z         bool
}

func stringToNode(s string) Node {
	matches := nodeMatch.FindStringSubmatch(s)
	if matches == nil {
		return Node{}
	}
	last_rune := rune(matches[1][2])
	n := Node{matches[1], matches[2], matches[3], last_rune == 'A', last_rune == 'Z'}
	return n
}

func Problem1(data *[]string) int {
	directions := strings.TrimSpace((*data)[0])
	nodes := make(map[string]Node)
	for _, s := range (*data)[2:] {
		n := stringToNode(s)
		nodes[n.name] = n
	}
	log.Debug().Str("directions", directions).Int("nodes", len(nodes)).Msg("Configuration loaded")
	c := "AAA"
	cycler := iterium.Cycle(iterium.New(strings.Split(directions, "")...))
	i := 0
	for c != "ZZZ" {
		d, _ := cycler.Next()
		n := nodes[c]
		if d == "R" {
			c = n.right
		} else {
			c = n.left
		}
		i++
	}
	return i
}

type cyclePosition struct {
	node string
	step int
}

func trackCycles(node string, nodes *map[string]Node, directions string) int {
	cycler := iterium.Cycle(iterium.New(strings.Split(directions, "")...))
	tracker := make(map[cyclePosition]int)
	c := node
	i := 0
	for {
		n := (*nodes)[c]
		cp := cyclePosition{c, i % len(directions)}
		if last_seen := (tracker)[cp]; last_seen != 0 {
			return i - last_seen
		}
		tracker[cp] = i
		d, _ := cycler.Next()
		if d == "R" {
			c = n.right
		} else {
			c = n.left
		}
		i++
	}
}

func Problem2(data *[]string) int {
	directions := strings.TrimSpace((*data)[0])
	nodes := make(map[string]Node)
	var current []string
	for _, s := range (*data)[2:] {
		n := stringToNode(s)
		nodes[n.name] = n
		if n.ends_in_a {
			current = append(current, n.name)
		}
	}
	cycle_sizes := make(map[string]int)
	for _, c := range current {
		cycle_sizes[c] = trackCycles(c, &nodes, directions)
	}
	lcm := 1
	for _, v := range cycle_sizes {
		lcm = internal.LCM(lcm, v)
	}
	return lcm
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
