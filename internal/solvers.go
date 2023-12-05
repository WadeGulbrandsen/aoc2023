package internal

import (
	"bufio"
	"fmt"
	"os"
)

func SumSolver(filename string, fn func(string, chan int)) int {
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
