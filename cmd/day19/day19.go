package main

import (
	"maps"
	"regexp"
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/functional"
	"github.com/WadeGulbrandsen/aoc2023/internal/span"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/rs/zerolog/log"
)

const Day = 19

type workflow []rule

type part map[string]int

type RangeMap map[string]span.Span

func (sm *RangeMap) Product() int {
	if len(*sm) == 0 {
		return 0
	}
	product := 1
	for _, s := range *sm {
		product *= span.Span(s).Len()
	}
	return product
}

type rule struct {
	category    string
	op          string
	val         int
	destination string
}

var ruleMatch = regexp.MustCompile(`([xmas])([<>])(\d+):(\w+)`)

func parseRule(s string) rule {
	matches := ruleMatch.FindStringSubmatch(s)
	if matches == nil {
		return rule{category: "x", op: ">", destination: s, val: 0}
	}
	val, err := strconv.Atoi(matches[3])
	if err != nil {
		panic(err)
	}
	return rule{category: matches[1], op: matches[2], val: val, destination: matches[4]}
}

type sorter struct {
	workflows map[string]workflow
}

func (s *sorter) AcceptPart(p part, wf string) bool {
	workflow, ok := s.workflows[wf]
	if !ok {
		return wf == "A"
	}
	for _, rule := range workflow {
		result := false
		switch rule.op {
		case "<":
			result = p[rule.category] < rule.val
		default:
			result = p[rule.category] > rule.val
		}
		if result {
			return s.AcceptPart(p, rule.destination)
		}
	}
	return false
}

func (s *sorter) AcceptedSpans(sm RangeMap, wf string) []RangeMap {
	workflow, ok := s.workflows[wf]
	if !ok {
		if wf == "A" {
			return []RangeMap{sm}
		}
		return nil
	}
	var accepted []RangeMap
	current := maps.Clone(sm)
	for _, rule := range workflow {
		next := maps.Clone(current)
		var n, k span.Span
		if rule.op == "<" {
			n, k = current[rule.category].SplitAt(rule.val)
		} else {
			k, n = current[rule.category].SplitAt(rule.val + 1)
		}
		next[rule.category] = n
		current[rule.category] = k
		accepted = append(accepted, s.AcceptedSpans(next, rule.destination)...)
	}
	return accepted
}

func parseWorkflow(s string) (string, workflow) {
	var wf workflow
	name, rules, found := strings.Cut(s, "{")
	if !found {
		return "", wf
	}
	rules = strings.Trim(rules, "}")
	for _, rule := range strings.Split(rules, ",") {
		wf = append(wf, parseRule(rule))
	}
	return name, wf
}

func parsePart(s string) part {
	part := make(part)
	for _, c := range strings.Split(strings.Trim(s, "{}"), ",") {
		if n, v, found := strings.Cut(c, "="); found {
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			part[n] = val
		}
	}
	return part
}

func parseInput(data *[]string) (sorter, []part) {
	w, p, found := functional.Cut(data, "")
	if !found {
		return sorter{}, nil
	}
	workflows := make(map[string]workflow)
	for _, wf := range w {
		name, workflow := parseWorkflow(wf)
		workflows[name] = workflow
	}
	s := sorter{workflows: workflows}
	var parts []part
	for _, part := range p {
		parts = append(parts, parsePart(part))
	}
	log.Debug().Msgf("Found %v workflows and %v parts", len(workflows), len(parts))
	return s, parts
}

func Problem1(data *[]string) int {
	s, parts := parseInput(data)
	var accepted []part
	for _, part := range parts {
		if s.AcceptPart(part, "in") {
			accepted = append(accepted, part)
		}
	}
	sum := 0
	for _, a := range accepted {
		vals := utils.GetMapValues(a)
		sum += functional.Sum(&vals)
	}
	return sum
}

func Problem2(data *[]string) int {
	s, _ := parseInput(data)
	sm := RangeMap{
		"x": {Start: 1, End: 4000},
		"m": {Start: 1, End: 4000},
		"a": {Start: 1, End: 4000},
		"s": {Start: 1, End: 4000},
	}
	accepted := s.AcceptedSpans(sm, "in")
	sum := 0
	for _, a := range accepted {
		sum += a.Product()
	}
	return sum
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
