package main

import (
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/rs/zerolog/log"
	"github.com/twmb/algoimpl/go/graph"
)

const Day = 25

func graphInput(data *[]string) *graph.Graph {
	g := graph.New(graph.Undirected)
	wires := make(map[string]graph.Node, 0)
	for _, s := range *data {
		before, after, _ := strings.Cut(s, ":")
		src := strings.TrimSpace(before)
		dests := strings.Split(strings.TrimSpace(after), " ")
		if _, ok := wires[src]; !ok {
			wires[src] = g.MakeNode()
		}
		for _, d := range dests {
			if _, ok := wires[d]; !ok {
				wires[d] = g.MakeNode()
			}
			g.MakeEdge(wires[src], wires[d])
		}
	}
	for k, n := range wires {
		*n.Value = k
	}
	return g
}

func Problem1(data *[]string) int {
	g := graphInput(data)
	var cuts []graph.Edge
	for len(cuts) != 3 {
		cuts = g.RandMinimumCut(16, 16)
	}
	for _, c := range cuts {
		log.Debug().Msgf("Cutting %v to %v", *c.Start.Value, *c.End.Value)
		g.RemoveEdge(c.Start, c.End)
	}
	groups := g.StronglyConnectedComponents()
	if len(groups) != 2 {
		return 0
	}
	return len(groups[0]) * len(groups[1])
}

func Problem2(data *[]string) int {
	return 0
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
