package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculate_fuel(x_positions []int, target int) int {
	fuel_spent := 0
	for _, x := range x_positions {
		delta := int(math.Abs(float64(x - target)))

		// fuel_spent += delta // part one
		fuel_spent += int(float64(1+delta) * float64(delta) / 2) // part two
	}
	return fuel_spent
}

func part_one(path string) {
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(raw)), ",")

	x_positions := []int{}

	for _, raw := range input {
		crab, err := strconv.Atoi(raw)
		if err != nil {
			log.Fatal(err)
		}
		x_positions = append(x_positions, crab)
	}

	// First of, calculate the min and max positions to make sure we're only testing sensible numbers
	var min_x, max_x int
	for _, x := range x_positions {
		if min_x > x {
			min_x = x
		} else if max_x < x {
			max_x = x
		}
	}

	minimum_fuel := math.MaxInt // would ideally do math.Inf but not sure how go handles float64(Inf) -> int(Inf) conversion
	minimum_fuel_at := 0
	for x := min_x; x < max_x; x++ {
		fuel_consumed := calculate_fuel(x_positions, x)
		if fuel_consumed < minimum_fuel {
			minimum_fuel_at = x
			minimum_fuel = fuel_consumed
		}
	}

	fmt.Printf("Found optimum fuel: %d used at x = %d\n", minimum_fuel, minimum_fuel_at)
}
