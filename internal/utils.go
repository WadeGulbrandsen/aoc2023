package internal

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type StartTime struct {
	item string
	time time.Time
}

func StartTimer(s string) StartTime {
	log.Println("START:", s)
	return StartTime{item: s, time: time.Now()}
}

func EndTimer(s StartTime) {
	end := time.Now()
	log.Println("END:  ", s.item, "ElapsedTime:", end.Sub(s.time))
}

func GetMapValues[M ~map[K]V, K comparable, V any](m M) []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func GetMapKeys[M ~map[K]V, K comparable, V any](m M) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetArgs() string {
	var filename string
	flag.StringVar(&filename, "input", "input.txt", "file to be processed")
	flag.Parse()
	return filename
}

func FileToLines(filename string) []string {
	var l []string
	fmt.Printf("\nReading %v\n", filename)
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return l
	}
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	readFile.Close()
	return l
}

func timeProblem[T any](msg string, d *[]string, p func(*[]string) T) T {
	start := StartTimer(msg)
	result := p(d)
	EndTimer(start)
	return result
}

func SolutionRunner[T any](day int, p1, p2 func(*[]string) T) {
	filename := GetArgs()
	fmt.Printf("Advent of Code 2023: Day %v\n", day)
	data := FileToLines(filename)

	fmt.Printf("\nThe answer for Day %v, Problem 1 is: %v\n\n", day, timeProblem(fmt.Sprintf("Day %v: Problem 1", day), &data, p1))
	fmt.Printf("\nThe answer for Day %v, Problem 2 is: %v\n\n", day, timeProblem(fmt.Sprintf("Day %v: Problem 2", day), &data, p2))
}
