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

	// Input map
	// Key: binary represented by a string
	// Value: binary value converted to dec
	input := make(map[string]int64)
	var o2_rating, co2_rating string
	var acc int

	for scanner.Scan() {
		if scanner.Text() == "" { // break on empty newline
			break
		}
		line := scanner.Text()

		for i, c := range line {
			incr, _ := strconv.Atoi(string(c))
			tally[i] += incr
		}

		acc++
		input[line], _ = strconv.ParseInt(line, 2, 64)
	}

	// This next part could be done in one line in some languages
	// It must be possible to do it in Go, but I like how clean it is
	for index, element := range tally {
		element -= acc / 2
		if element > 0 {
			o2_rating = o2_rating[:index] + "1" + o2_rating[index:]
			co2_rating = co2_rating[:index] + "0" + co2_rating[index:]
		} else {
			o2_rating = o2_rating[:index] + "0" + o2_rating[index:]
			co2_rating = co2_rating[:index] + "1" + co2_rating[index:]
		}
	}

	// Recursive find algorithm
	// Because why not be creative
	if val, ok := input[o2_rating]; ok {

	} else {
		for i := 10; i >= 0; i-- { //C++ style, nice!

		}
	}
	o2_rating_converted, _ := strconv.ParseInt(o2_rating, 2, 64)
	co2_rating_converted, _ := strconv.ParseInt(co2_rating, 2, 64)
	fmt.Printf("Power Consumption: %v",
		o2_rating_converted*co2_rating_converted)
}

// filter_input is a recursive function which starting at the end
// of a string binary representation does a brute-force replacement
// of all bits past the index bit, checking if the generated binary
// representation exists
func filter_input(rating string, input map[string]int, index int) int {
	if val, ok := input[rating]; ok {
		return val
	} else {
		index--
		rating = rating[:index] + swap(string(rating[index])) + rating[index:]
		filter_input(rating)
	}
}

func swap(s string) string {
	if s == "1" {
		return "0"
	} else {
		return "1"
	}
}
