package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal"
	"github.com/rs/zerolog/log"
)

const Day = 19

var ruleMatch = regexp.MustCompile(`([xmas])([<>])(\d+):(\w+)`)

type rule struct {
	category    rune
	op          func(int, int) bool
	val         int
	destination string
}

func noop(x, y int) bool {
	return true
}

func lt(x, y int) bool {
	return x < y
}

func gt(x, y int) bool {
	return x > y
}

func parseRule(s string) rule {
	matches := ruleMatch.FindStringSubmatch(s)
	op := noop
	if matches == nil {
		return rule{op: op, destination: s}
	}
	switch matches[2] {
	case "<":
		op = lt
	case ">":
		op = gt
	}
	val, err := strconv.Atoi(matches[3])
	if err != nil {
		panic(err)
	}
	return rule{category: rune(matches[1][0]), op: op, val: val, destination: matches[4]}
}

type workflow []rule

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

type part map[rune]int

func parsePart(s string) part {
	part := make(part)
	for _, c := range strings.Split(strings.Trim(s, "{}"), ",") {
		if n, v, found := strings.Cut(c, "="); found {
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			part[rune(n[0])] = val
		}
	}
	return part
}

func followWorkflows(p part, current string, workflows *map[string]workflow) bool {
	switch current {
	case "A":
		return true
	case "R":
		return false
	}
	for _, rule := range (*workflows)[current] {
		if rule.op(p[rule.category], rule.val) {
			return followWorkflows(p, rule.destination, workflows)
		}
	}
	return false
}

func Problem1(data *[]string) int {
	w, p, found := internal.Cut(data, "")
	if !found {
		return 0
	}
	workflows := make(map[string]workflow)
	for _, wf := range w {
		name, workflow := parseWorkflow(wf)
		workflows[name] = workflow
	}
	var parts []part
	for _, part := range p {
		parts = append(parts, parsePart(part))
	}
	log.Debug().Msgf("Found %v workflows and %v parts", len(workflows), len(parts))
	var accepted []part
	for _, part := range parts {
		if followWorkflows(part, "in", &workflows) {
			accepted = append(accepted, part)
		}
	}
	sum := 0
	for _, a := range accepted {
		vals := internal.GetMapValues(a)
		sum += internal.Sum(&vals)
	}
	return sum
}

func Problem2(data *[]string) int {
	return 0
}

func main() {
	internal.RunSolutions(Day, Problem1, Problem2, "input.txt", "sample.txt", -1)
}
