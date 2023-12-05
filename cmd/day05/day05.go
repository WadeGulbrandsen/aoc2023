package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

var steps = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

type RangeMap struct {
	dest, source, length int
}

func (r *RangeMap) Lookup(v int) (int, bool) {
	if v >= r.source && v < r.source+r.length {
		return r.dest + v - r.source, true
	}
	return -1, false
}

func (r *RangeMap) RLookup(v int) (int, bool) {
	if v >= r.dest && v < r.dest+r.length {
		return r.source + v - r.dest, true
	}
	return -1, false
}

type Almanac struct {
	lock  sync.RWMutex
	seeds []int
	maps  map[string][]RangeMap
}

func (a *Almanac) Step(v int, step string) int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	for _, r := range a.maps[step] {
		if next, found := r.Lookup(v); found {
			return next
		}
	}
	return v
}

func (a *Almanac) Backstep(v int, step string) int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	for _, r := range a.maps[step] {
		if prev, found := r.RLookup(v); found {
			return prev
		}
	}
	return v
}

func (a *Almanac) LocationToSeed(loc int) int {
	v := loc
	for i := 0; i < len(steps); i++ {
		step := steps[len(steps)-i-1]
		v = a.Backstep(v, step)
	}
	return v
}

func (a *Almanac) SeedToLocation(seed int) int {
	v := seed
	for _, step := range steps {
		v = a.Step(v, step)
	}
	return v
}

func getIntsFromString(s string, sep string) []int {
	var ints []int
	for _, n := range strings.Split(s, sep) {
		if v, err := strconv.Atoi(strings.TrimSpace(n)); err == nil {
			ints = append(ints, v)
		}
	}
	return ints
}

func getRangeMaps(s string) []RangeMap {
	var r []RangeMap
	ch := make(chan RangeMap)
	lines := strings.Split(s, "\n")
	for _, l := range lines {
		go func(str string) {
			ints := getIntsFromString(strings.TrimSpace(str), " ")
			ch <- RangeMap{ints[0], ints[1], ints[2]}
		}(l)
	}
	for i := 0; i < len(lines); i++ {
		r = append(r, <-ch)
	}
	return r
}

func fileToAlmanac(filename string, a *Almanac) {
	var wg sync.WaitGroup
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	sections := strings.Split(string(data), "\n\n")
	for _, s := range sections {
		title, body, found := strings.Cut(s, ":")
		if found {
			title, _ = strings.CutSuffix(strings.TrimSpace(title), " map")
			body = strings.TrimSpace(body)
			switch {
			case title == "seeds":
				a.lock.Lock()
				a.seeds = getIntsFromString(body, " ")
				a.lock.Unlock()
			case slices.Contains(steps, title):
				wg.Add(1)
				go func(t string, b string) {
					defer wg.Done()
					a.lock.Lock()
					a.maps[t] = getRangeMaps(b)
					a.lock.Unlock()
				}(title, body)
			}
		}
	}
	wg.Wait()
}

func Problem1(filename string) int {
	almanac := Almanac{maps: make(map[string][]RangeMap)}
	fileToAlmanac(filename, &almanac)
	almanac.lock.RLock()
	seeds := almanac.seeds
	almanac.lock.RUnlock()
	location := math.MaxInt
	ch := make(chan int)
	for _, seed := range seeds {
		go func(s int) {
			ch <- almanac.SeedToLocation(s)
		}(seed)
	}
	for i := 0; i < len(seeds); i++ {
		location = min(location, <-ch)
	}
	return location
}

func Problem2(filename string) int {
	// seen := make(map[int]bool)
	almanac := Almanac{maps: make(map[string][]RangeMap)}
	fileToAlmanac(filename, &almanac)
	almanac.lock.RLock()
	seeds := almanac.seeds
	almanac.lock.RUnlock()
	ranges := make(map[int]int)
	for i := 0; i < len(seeds); i += 2 {
		ranges[seeds[i]] = seeds[i] + seeds[i+1] - 1
	}
	for i := 0; i <= 157211394; i++ {
		if i%1000000 == 0 {
			fmt.Printf("Location %v\n", i)
		}
		seed := almanac.LocationToSeed(i)
		for low, high := range ranges {
			if seed >= low && seed <= high {
				fmt.Printf("Seed %v to location %v\n", seed, i)
				return i
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("Advent of Code 2023")
	fmt.Printf("\nThe answer for Day 05, Problem 1 is: %v\n\n", Problem1("input.txt"))
	fmt.Printf("\nThe answer for Day 05, Problem 2 is: %v\n\n", Problem2("input.txt"))
}
