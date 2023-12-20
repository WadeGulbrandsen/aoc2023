package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/WadeGulbrandsen/aoc2023/internal/solve"
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

const Day = 1

var number_names = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getNumber(s string) int {
	first := strings.IndexFunc(s, unicode.IsDigit)
	last := strings.LastIndexFunc(s, unicode.IsDigit)
	if first < 0 || last < 0 {
		fmt.Printf("No digits in %v", s)
		return 0
	}
	sn := fmt.Sprintf("%s%s", string(s[first]), string(s[last]))
	v, err := strconv.Atoi(sn)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	fmt.Printf(" %v found in %v\n", v, s)
	return v
}

func getNumberWithWords(s string) int {
	first_index := math.MaxInt
	last_index := -1
	first_int, last_int := 0, 0
	if i := strings.IndexFunc(s, unicode.IsDigit); i != -1 {
		v, err := strconv.Atoi(string(s[i]))
		if err != nil {
			fmt.Println(err)
		} else {
			first_int = v
			first_index = i
		}
	}
	if i := strings.LastIndexFunc(s, unicode.IsDigit); i != -1 {
		v, err := strconv.Atoi(string(s[i]))
		if err != nil {
			fmt.Println(err)
		} else {
			last_int = v
			last_index = i
		}
	}
	for name, v := range number_names {
		if i := strings.Index(s, name); i != -1 && i < first_index {
			first_index = i
			first_int = v
		}
		if i := strings.LastIndex(s, name); i != -1 && i > last_index {
			last_index = i
			last_int = v
		}
	}
	value := (first_int * 10) + last_int
	fmt.Printf(" %v found in %v\n", value, s)
	return value
}

func Problem1(data *[]string) int {
	return solve.SumSolver(data, getNumber)
}

func Problem2(data *[]string) int {
	return solve.SumSolver(data, getNumberWithWords)
}

func main() {
	utils.CmdSolutionRunner(Day, Problem1, Problem2)
}
