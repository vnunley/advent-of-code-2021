package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Awaiting input")

	var lastSum, incr int

	input := getInput()

	for i := 0; i < len(input)-2; i++ {
		curSum := input[i] + input[i+1] + input[i+2]

		if i == 0 {
			lastSum = curSum
			continue
		}

		if curSum > lastSum {
			incr++
		}

		lastSum = curSum
	}

	fmt.Printf("Total increases in depth: %d", incr)
}

func getInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	input := []int{}

	for {
		// reads user input until \n by default
		scanner.Scan()
		line := scanner.Text()

		// check if line was empty
		if len(line) == 0 {
			break
		}

		lineInt, _ := strconv.Atoi(line)

		input = append(input, lineInt)
	}
	return (input)
}
