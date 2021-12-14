package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Awaiting input")

	scanner := bufio.NewScanner(os.Stdin)
	var x, y int

	for scanner.Scan() {
		if scanner.Text() == "" { // break on empty newline
			break
		}

		input := strings.Split(scanner.Text(), " ") // get slice
		s := input[0]
		n, _ := strconv.Atoi(input[1])

		if s == "forward" {
			x += n
		} else if s == "up" {
			y -= n
		} else if s == "down" {
			y += n
		}
	}

	fmt.Printf("Product: %d", y*x)
}
