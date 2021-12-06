package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fish_timer = 7
const new_fish_delay = 2

// Take a school, and model what it looks like one day later
func tick(school []int) []int {
	new_school := []int{}

	for _, fish := range school {
		if fish == 0 {
			// Add the same fish, reset, and a new fish, plus the extra delay days
			new_school = append(new_school, fish_timer, fish_timer+new_fish_delay)
		} else {
			new_school = append(new_school, fish-1)
		}
	}

	return new_school
}

func part_one(path string) {
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(raw)), ",")

	school := []int{}

	for _, raw := range input {
		fish, err := strconv.Atoi(raw)
		if err != nil {
			log.Fatal(err)
		}
		school = append(school, fish)
	}

	fmt.Printf("Initial School: %v\n", school)

	for day := 1; day <= tick_count; day++ {
		school = tick(school)

		fmt.Printf("Day %-2d - School of %-4d\n", day, len(school))
	}
}
