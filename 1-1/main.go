package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Awaiting input")

	var last, incr int

	for n, i := range getInput() {
		current, _ := strconv.Atoi(i)

		if n == 0 {
			last = current
			continue
		}

		if current > last {
			incr++
		}
		last = current
	}

	fmt.Printf("Total increases in depth: %d", incr)
}

func getInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}

	for {
		// reads user input until \n by default
		scanner.Scan()
		line := scanner.Text()

		// check if line was empty
		if len(line) == 0 {
			break
		}

		input = append(input, line)
	}
	return (input)
}
