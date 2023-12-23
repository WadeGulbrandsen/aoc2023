package main

import (
	"container/heap"
	"strings"

	"github.com/WadeGulbrandsen/aoc2023/internal/heaps"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/rs/zerolog/log"
)

const Day = 20

type factory map[string]module

func (f factory) PushButton(to_watch string) (map[pulse]int, map[string]bool) {
	pulses := make(map[pulse]int)
	pq := &heaps.PriorityQueue[message]{}
	heap.Init(pq)
	heap.Push(pq, &heaps.Item[message]{Value: message{input: "button", destination: "broadcaster", pulse: low}, Priority: 1})
	i := 0
	seen_high := make(map[string]bool)
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*heaps.Item[message]).Value
		if current.pulse == high && current.destination == to_watch {
			seen_high[current.input] = true
			log.Info().Msgf("%v -%v-> %v", current.input, current.pulse, current.destination)
		}
		pulses[current.pulse]++
		if mod, ok := f[current.destination]; ok {
			for _, next := range mod.handle_message(current) {
				i--
				heap.Push(pq, &heaps.Item[message]{Value: next, Priority: i})
			}
		}
	}
	return pulses, seen_high
}

func parseModule(s string) (module, bool) {
	before, after, found := strings.Cut(s, " -> ")
	if !found {
		return nil, false
	}
	destinations := strings.Split(after, ", ")
	switch {
	case before == "broadcaster":
		return &broadcast{name: before, destinations: destinations}, true
	case before[0] == '%':
		return &flip_flop{name: before[1:], destinations: destinations, state: false}, true
	case before[0] == '&':
		return &conjunction{name: before[1:], destinations: destinations, inputs: make(map[string]pulse)}, true
	}
	return nil, false
}

func parseModules(data *[]string) (factory, string) {
	factory := make(factory)
	lookup_inputs := make(map[string][]string)
	for _, line := range *data {
		if mod, found := parseModule(line); found {
			factory[mod.get_name()] = mod
			for _, destination := range mod.get_destinations() {
				lookup_inputs[destination] = append(lookup_inputs[destination], mod.get_name())
			}
		}
	}
	for name, mod := range factory {
		if mod.get_type() == con {
			mod.add_inputs(lookup_inputs[name])
		}
	}

	if rx_parent := lookup_inputs["rx"]; len(rx_parent) != 0 {
		return factory, rx_parent[0]
	}
	return factory, ""
}

func Problem1(data *[]string) int {
	factory, _ := parseModules(data)
	lows, highs := 0, 0
	for i := 0; i < 1000; i++ {
		pulses, _ := factory.PushButton("")
		lows += pulses[low]
		highs += pulses[high]
	}
	return lows * highs
}

func Problem2(data *[]string) int {
	factory, rx_parent := parseModules(data)
	if rx_parent == "" {
		return 0
	}
	need := len(*factory[rx_parent].get_inputs())
	results := make(map[string]int)
	for i := 1; i < 10000; i++ {
		_, seen_high := factory.PushButton(rx_parent)
		for n := range seen_high {
			if _, ok := results[n]; !ok {
				results[n] = i
				if len(results) == need {
					vals := utils.GetMapValues(results)
					lcm := vals[0]
					for _, v := range vals[1:] {
						lcm = utils.LCM(lcm, v)
					}
					return lcm
				}
			}
		}
	}
	return 0
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
