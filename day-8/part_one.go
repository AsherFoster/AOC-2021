package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Given a shuffled digit, make a guess of what it is (returning -1 if unknown)
func guess_digit(digit Digit) int {
	illuminated := count_segments(digit)
	// fmt.Printf("Digit %v has %d segments illuminated\n", digit, illuminated)
	if illuminated == 2 {
		return 1
	} else if illuminated == 3 {
		return 7
	} else if illuminated == 4 {
		return 4
	} else if illuminated == 7 {
		return 8
	}
	return -1
}

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := [][2][]Digit{} // yes this type is horrific, shhhh

	// Parse the each
	for scanner.Scan() {
		line := [2][]Digit{}
		raw_digits := strings.Split(scanner.Text(), " ")
		delimeter_hit := false

		for _, digit := range raw_digits {
			if digit == "|" {
				delimeter_hit = true
				continue
			}
			segments := parse_digit(digit)
			if delimeter_hit {
				line[1] = append(line[1], segments)
			} else {
				line[0] = append(line[0], segments)
			}
		}

		lines = append(lines, line)
	}

	easy_occurences := 0
	for _, line := range lines {
		_, outputs := line[0], line[1]

		// make a guess at which input maps to which output
		for _, output := range outputs {
			guess := guess_digit(output)
			if guess != -1 {
				easy_occurences++
			}
		}
		// use this mapping to decode the outputs

	}
	// win
	fmt.Println(easy_occurences)
}
