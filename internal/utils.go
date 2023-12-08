package internal

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type StartTime struct {
	item string
	time time.Time
}

func StartTimer(s string) StartTime {
	log.Trace().Msgf("START: %v", s)
	return StartTime{item: s, time: time.Now()}
}

func EndTimer(s StartTime) {
	end := time.Now()
	elapsed := end.Sub(s.time)
	log.Trace().Str("Elapsed", elapsed.String()).Msgf("END:   %v", s.item)
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

func GetArgs() (string, int) {
	var filename string
	var loglevel int
	flag.StringVar(&filename, "input", "input.txt", "file to be processed")
	flag.IntVar(&loglevel, "level", 1, "will log the selected level and higher \n-1 Trace\n 0 Debug\n 1 Info\n 2 Warn\n 3 Error\n 4 Fatal\n 5 Panic")
	flag.Parse()
	return filename, loglevel
}

func FileToLines(filename string) []string {
	var l []string
	log.Info().Msgf("Reading %v", filename)
	readFile, err := os.Open(filename)
	if err != nil {
		log.Err(err).Send()
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

func RunSolutions[T any](day int, p1, p2 func(*[]string) T, f1 string, f2 string, loglevel int) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if loglevel <= 5 && loglevel >= -1 {
		zerolog.SetGlobalLevel(zerolog.Level(loglevel))
	}
	fmt.Printf("Advent of Code 2023: Day %v\n", day)
	data := FileToLines(f1)
	fmt.Printf("The answer for Day %v, Problem 1 is: %v\n", day, timeProblem(fmt.Sprintf("Day %v: Problem 1", day), &data, p1))
	if f1 != f2 {
		data = FileToLines(f2)
	}
	fmt.Printf("The answer for Day %v, Problem 2 is: %v\n", day, timeProblem(fmt.Sprintf("Day %v: Problem 2", day), &data, p2))
}

func CmdSolutionRunner[T any](day int, p1, p2 func(*[]string) T) {
	filename, loglevel := GetArgs()
	RunSolutions(day, p1, p2, filename, filename, loglevel)
}
