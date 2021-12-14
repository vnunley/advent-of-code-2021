package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Awaiting input")

	scanner := bufio.NewScanner(os.Stdin)

	// Tally Slice
	// Every element corrosponds to the binary bit
	tally := [12]int{}
	var gamma, epsilon string
	var acc int

	for scanner.Scan() {
		if scanner.Text() == "" { // break on empty newline
			break
		}

		for i, c := range scanner.Text() {
			incr, _ := strconv.Atoi(string(c))
			tally[i] += incr
		}

		acc++
	}

	// This next part could be done in one line in some languages
	// It must be possible to do it in Go, but I like how clean it is
	for index, element := range tally {
		element -= acc / 2
		if element > 0 {
			gamma = gamma[:index] + "1" + gamma[index:]
			epsilon = epsilon[:index] + "0" + epsilon[index:]
		} else {
			gamma = gamma[:index] + "0" + gamma[index:]
			epsilon = epsilon[:index] + "1" + epsilon[index:]
		}
	}

	gamma_converted, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_converted, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Printf("Power Consumption: %v",
		gamma_converted*epsilon_converted)
}
