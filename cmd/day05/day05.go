package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/WadeGulbrandsen/aoc2023/internal"
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

var backSteps = []string{
	"humidity-to-location",
	"temperature-to-humidity",
	"light-to-temperature",
	"water-to-light",
	"fertilizer-to-water",
	"soil-to-fertilizer",
	"seed-to-soil",
}

type Range = internal.Range
type RangeList = internal.RangeList

type RangeMap struct {
	dest, source, length int
}

func (r *RangeMap) GetRanges() (Range, Range) {
	return Range{Start: r.source, End: r.source + r.length - 1},
		Range{Start: r.dest, End: r.dest + r.length - 1}
}

func (r *RangeMap) SourceRangeMapper() ProcessingMap {
	s, _ := r.GetRanges()
	return ProcessingMap{input: s, offset: r.dest - r.source}
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

func (a *Almanac) GetMapperForStep(step string) ProcessingMapList {
	var p ProcessingMapList
	for _, m := range a.maps[step] {
		p = append(p, m.SourceRangeMapper())
	}
	p.Sort()
	return p
}

func (a *Almanac) Step(v int, step string) int {
	for _, r := range a.maps[step] {
		if next, found := r.Lookup(v); found {
			return next
		}
	}
	return v
}

func (a *Almanac) Backstep(v int, step string) int {
	for _, r := range a.maps[step] {
		if prev, found := r.RLookup(v); found {
			return prev
		}
	}
	return v
}

func (a *Almanac) LocationToSeed(loc int) int {
	v := loc
	for _, step := range backSteps {
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
	defer internal.Un(internal.Trace("Problem1"))
	almanac := Almanac{maps: make(map[string][]RangeMap)}
	fileToAlmanac(filename, &almanac)
	seeds := almanac.seeds
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

type ProcessingMap struct {
	input  Range
	offset int
}

func (p ProcessingMap) Process(r Range) (Range, Range, Range) {
	before, contained, after := p.input.SplitOtherRange(r)
	output := Range{}
	if !contained.IsEmpty() {
		output.Start = contained.Start + p.offset
		output.End = contained.End + p.offset
	}
	return before, output, after
}

func cmpPM(a, b ProcessingMap) int {
	if n := internal.CompareRanges(a.input, b.input); n != 0 {
		return n
	}
	return cmp.Compare(a.offset, b.offset)
}

type ProcessingMapList []ProcessingMap

func (p ProcessingMapList) Sort() {
	slices.SortFunc(p, cmpPM)
}

func (p ProcessingMapList) Process(r RangeList) RangeList {
	var next RangeList
	for _, rl := range r {
		remaining := rl
		for _, pm := range p {
			before, processed, after := pm.Process(remaining)
			remaining = after
			next = append(next, before, processed)
			if remaining.IsEmpty() {
				break
			}
		}
		if !remaining.IsEmpty() {
			next = append(next, remaining)
		}
	}
	next = next.FilterEmpty()
	next = slices.Compact(next)
	return next
}

func getLocationRangesFromSeedRanges(a *Almanac, r RangeList, step int) RangeList {
	r.Sort()
	if step < 0 || step >= len(steps) {
		return r
	}
	m := a.GetMapperForStep(steps[step])
	next := m.Process(r)
	return getLocationRangesFromSeedRanges(a, next, step+1)
}

func Problem2(filename string) int {
	defer internal.Un(internal.Trace("Problem2 (Now with ranges!)"))
	almanac := Almanac{maps: make(map[string][]RangeMap)}
	fileToAlmanac(filename, &almanac)
	seeds := almanac.seeds
	var ranges RangeList
	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, Range{Start: seeds[i], End: seeds[i] + seeds[i+1] - 1})
	}
	fmt.Println(ranges)
	locations := getLocationRangesFromSeedRanges(&almanac, ranges, 0)
	if len(locations) > 0 {
		return locations[0].Start
	}
	return 0
}

func main() {
	filename := "input.txt"
	fmt.Println("Advent of Code 2023")
	fmt.Printf("\nThe answer for Day 05, Problem 1 is: %v\n\n", Problem1(filename))
	fmt.Printf("\nThe answer for Day 05, Problem 2 is: %v\n\n", Problem2(filename))
}
