package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// const fish_timer = 6
// const new_fish_delay = 2

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Take a population, and shift it along by one day
func tick_smarter(population []int) []int {
	new_population := make([]int, fish_timer+new_fish_delay)

	copy(new_population, population[1:])
	new_population[8] = population[0]
	new_population[6] += population[0]

	return new_population
}

func part_two(path string) {
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.TrimSpace(string(raw)), ",")

	population := make([]int, fish_timer+new_fish_delay)

	for _, raw := range input {
		age, err := strconv.Atoi(raw)
		if err != nil {
			log.Fatal(err)
		}
		population[age]++
	}

	fmt.Printf("Initial Population: %v\n", population)

	for day := 1; day <= tick_count; day++ {
		population = tick_smarter(population)

		fmt.Printf("Day %-2d - School of %-4d\n", day, sum(population))
	}
}
