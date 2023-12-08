// tt: total time
// th: time holding the button
// tm: time moving
// dt: distance traveled
//
// equation 1: th + tm = tt
//             th = tt - tm
//
// equation 2: th * tm = dt
//             th = dt/tm
//
// combined:   tt - tm = dt/tm
//             (tt - tm) * tm = dt
//             (tt*tm) - (tm*tm) = dt
//             (tt*tm) = (tm*tm) + dt
//             0 = (tm*tm) - (tt*tm) + dt
//
// Quadratic equation!
// a = 1
// b = -tt
// c = dt

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/WadeGulbrandsen/aoc2023/internal"
)

const Day = 6

func Quadratic(a, b, c float64) (float64, float64, error) {
	discriminant := (b * b) - (4 * a * c)
	rooted := math.Sqrt(discriminant)
	if math.IsNaN(rooted) {
		return math.NaN(), math.NaN(), fmt.Errorf("discriinant is less than zero: %v^2 - 4(%v)(%v) = %v", b, a, c, discriminant)
	}
	x1 := (-b + rooted) / (2 * a)
	x2 := (-b - rooted) / (2 * a)
	return x1, x2, nil
}

type Race struct {
	time, distance int
}

func getIntsFromString(s string) []int {
	var ints []int
	for _, x := range strings.Split(s, " ") {
		if v, err := strconv.Atoi(strings.TrimSpace(x)); err == nil {
			ints = append(ints, v)
		}
	}
	return ints
}

func getRacesFromStrings(data *[]string) []Race {
	var r []Race
	if len(*data) != 2 {
		panic(fmt.Errorf("there should be 2 lines but found %v", len(*data)))
	}
	times, distances := (*data)[0], (*data)[1]
	if t, d := getIntsFromString(times), getIntsFromString(distances); t != nil && d != nil && len(t) == len(d) {
		for i, time := range t {
			r = append(r, Race{time, d[i]})
		}
	}
	return r
}

func Problem1(data *[]string) int {
	races := getRacesFromStrings(data)
	var results []int
	for _, r := range races {
		if x1, x2, err := Quadratic(1, -float64(r.time), float64(r.distance)); err == nil {
			h := math.Ceil(max(x1, x2)) - 1
			l := math.Floor(min(x1, x2)) + 1
			diff := h - l
			results = append(results, int(diff)+1)
		}
	}
	product := 1
	for _, r := range results {
		product *= r
	}
	return product
}

func onlyDigits(s string) int {
	keepDigits := func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}
	digits := strings.Map(keepDigits, s)
	if v, err := strconv.Atoi(strings.TrimSpace(digits)); err == nil {
		return v
	}
	return -1
}

func Problem2(data *[]string) int {
	if len(*data) != 2 {
		panic(fmt.Errorf("there should be 2 lines but found %v", len(*data)))
	}
	times, distances := (*data)[0], (*data)[1]
	b, c := onlyDigits(times), onlyDigits(distances)
	if x1, x2, err := Quadratic(1, -float64(b), float64(c)); err == nil {
		h := math.Ceil(max(x1, x2)) - 1
		l := math.Floor(min(x1, x2)) + 1
		diff := h - l
		return int(diff) + 1
	}
	return 0
}

func main() {
	internal.CmdSolutionRunner(Day, Problem1, Problem2)
}
