package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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

func getNumber(s string, ch chan int) {
	first := strings.IndexFunc(s, unicode.IsDigit)
	last := strings.LastIndexFunc(s, unicode.IsDigit)
	if first < 0 || last < 0 {
		fmt.Printf("No digits in %v", s)
		return
	}
	sn := fmt.Sprintf("%s%s", string(s[first]), string(s[last]))
	v, err := strconv.Atoi(sn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(" %v found in %v\n", v, s)
	ch <- v
}

func getNumberWithWords(s string, ch chan int) {
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
	ch <- value
}

func problemSolver(filename string, fn func(string, chan int)) int {
	fmt.Printf("Opening %v\n", filename)
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	ch := make(chan int)

	l := 0
	for fileScanner.Scan() {
		go fn(fileScanner.Text(), ch)
		l++
	}
	readFile.Close()

	sum := 0
	for i := 0; i < l; i++ {
		sum += <-ch
	}

	return sum
}

func Problem1(filename string) int {
	return problemSolver(filename, getNumber)
}

func Problem2(filename string) int {
	return problemSolver(filename, getNumberWithWords)
}

func main() {
	fmt.Println("Advent of Code 2023")
	fmt.Printf("\nThe answer for Day 01, Problem 1 is: %v\n\n", Problem1("input.txt"))
	fmt.Printf("\nThe answer for Day 01, Problem 2 is: %v\n\n", Problem2("input.txt"))
}
