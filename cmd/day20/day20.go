package main

import (
	"container/heap"
	"strings"

	priorityqueue "github.com/WadeGulbrandsen/aoc2023/internal/priority_queue"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
	"github.com/rs/zerolog/log"
)

const Day = 20

type factory map[string]module

func (f factory) PushButton() (map[pulse]int, bool) {
	pulses := make(map[pulse]int)
	pq := &priorityqueue.PriorityQueue[message]{}
	heap.Init(pq)
	heap.Push(pq, &priorityqueue.Item[message]{Value: message{input: "button", destination: "broadcaster", pulse: low}, Priority: 1})
	i := 0
	rx_low := false
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*priorityqueue.Item[message]).Value
		if current.pulse == low && current.destination == "rx" {
			log.Info().Msgf("%v -%v-> %v", current.input, current.pulse, current.destination)
			rx_low = true
		}
		pulses[current.pulse]++
		if mod, ok := f[current.destination]; ok {
			for _, next := range mod.handle_message(current) {
				i--
				heap.Push(pq, &priorityqueue.Item[message]{Value: next, Priority: i})
			}
		}
	}
	return pulses, rx_low
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

func parseModules(data *[]string) factory {
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
	return factory
}

func Problem1(data *[]string) int {
	factory := parseModules(data)
	lows, highs := 0, 0
	for i := 0; i < 1000; i++ {
		pulses, _ := factory.PushButton()
		lows += pulses[low]
		highs += pulses[high]
	}
	return lows * highs
}

func Problem2(data *[]string) int {
	factory := parseModules(data)
	for i := 1; i < 1000; i++ {
		_, seen := factory.PushButton()
		if seen {
			return i
		}
	}
	return 0
}

func main() {
	utils.RunSolutions(Day, Problem1, Problem2, "input.txt", "input.txt", -1)
}
