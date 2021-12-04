package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func nightmare_binary_search(lines []uint, mask uint, use_most_common bool) uint {
	var ones []uint
	var zeros []uint

	for _, line := range lines {
		if line&mask != 0 {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}

	winners := ones

	// This feels so wrong, but hey, it works 🤷‍♀️
	if use_most_common {
		if len(zeros) > len(ones) {
			winners = zeros
		}
	} else {
		if len(zeros) <= len(ones) {
			winners = zeros
		}
	}

	if len(winners) == 1 {
		return winners[0]
	} else {
		// Recurse, checking the next digit along.
		// This makes the assumption that we'll reach a single result eventually
		return nightmare_binary_search(winners, mask>>1, use_most_common)
	}
}

func part_two(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []uint{}

	for scanner.Scan() {
		line, err := strconv.ParseUint(scanner.Text(), 2, line_width)
		if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, uint(line))
	}

	start_mask := uint(math.Pow(2, line_width-1)) // the one in the left-most column
	oxygen_rating := nightmare_binary_search(lines, start_mask, true)
	co2_rating := nightmare_binary_search(lines, start_mask, false)

	fmt.Printf("Oxygen: %d, CO2: %d, Life support: %d\n", oxygen_rating, co2_rating, oxygen_rating*co2_rating)
}
