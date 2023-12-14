package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 12

type countParams struct {
	data  string
	sizes string
}

type safeCache struct {
	lock  sync.RWMutex
	cache map[countParams]int
}

func (c *safeCache) Get(p *countParams) (int, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	v, ok := c.cache[*p]
	return v, ok
}

func (c *safeCache) Set(p *countParams, v int) {
	c.lock.Lock()
	c.cache[*p] = v
	c.lock.Unlock()
}

var cache safeCache = safeCache{cache: make(map[countParams]int)}

func countArrangements(data string, sizes []int) int {
	params := countParams{data: data, sizes: fmt.Sprint(sizes)}
	if v, ok := cache.Get(&params); ok {
		return v
	}
	total := internal.Sum(&sizes)
	minimum := strings.Count(data, "#")
	maximum := len(data) - strings.Count(data, ".")
	result := 0
	switch {
	case minimum > total || maximum < total:
		result = 0
	case total == 0:
		result = 1
	case data[0] == '.':
		result = countArrangements(data[1:], sizes)
	case data[0] == '#':
		size := sizes[0]
		if !strings.ContainsRune(data[:size], '.') {
			if size == len(data) {
				result = 1
				break
			}
			if data[size] != '#' {
				result = countArrangements(data[size+1:], sizes[1:])
				break
			}
		}
		result = 0
	default:
		result = countArrangements(data[1:], sizes) + countArrangements("#"+data[1:], sizes)
	}
	cache.Set(&params, result)
	return result
}

func parse(s string, scale int) (string, []int) {
	before, after, found := strings.Cut(s, " ")
	if !found {
		return "", nil
	}
	ints := internal.GetIntsFromString(after, ",")
	data := before
	sizes := ints
	for i := 1; i < scale; i++ {
		data += "?" + before
		sizes = append(sizes, ints...)
	}
	return data, sizes
}

func getCount(s string) int {
	return countArrangements(parse(s, 1))
}

func expandAndCount(s string) int {
	return countArrangements(parse(s, 5))
}

func Problem1(data *[]string) int {
	return internal.SumSolver(data, getCount)
}

func Problem2(data *[]string) int {
	return internal.SumSolver(data, expandAndCount)
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
