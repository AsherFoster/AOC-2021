package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func part_two(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := [][2][]Digit{} // yes this type is horrific, shhhh

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

	output_sum := 0
	for _, line := range lines {
		examples, outputs := line[0], line[1]

		// make a guess at which input maps to which output
		mapping := magically_determine_mapping(examples)

		// use this mapping to decode the outputs
		for i, output := range outputs {
			val := decode_with_mapping(output, mapping)
			output_sum += val * int(math.Pow10(3-i))
		}
	}

	// win
	fmt.Println(output_sum)
}
