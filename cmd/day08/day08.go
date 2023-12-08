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
	log.Debug().Int("directions", len(directions)).Int("nodes", len(nodes)).Msg("Configuration loaded")
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

func findFirstZ(start string, nodes *map[string]Node, directions string) int {
	cycler := iterium.Cycle(iterium.New(strings.Split(directions, "")...))
	next := (*nodes)[start]
	i := 0
	for !next.ends_in_z {
		d, _ := cycler.Next()
		if d == "R" {
			next = (*nodes)[next.right]
		} else {
			next = (*nodes)[next.left]
		}
		i++
	}
	return i
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

	lcm := findFirstZ(current[0], &nodes, directions)
	for _, v := range current[1:] {
		lcm = internal.LCM(lcm, findFirstZ(v, &nodes, directions))
	}
	return lcm
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
