package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const max_ticks = 100

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := Grid{}
	// Parse the input
	for scanner.Scan() {
		raw_vals := strings.Split(scanner.Text(), "")
		for _, raw := range raw_vals {
			val, err := strconv.Atoi(raw)
			if err != nil {
				log.Fatal(err)
			}
			grid.values = append(grid.values, val)
		}
	}

	total_flashes := 0
	for i := 0; i < max_ticks; i++ {
		flashes := grid.tick()
		total_flashes += flashes

		fmt.Printf("\n%-2d flashes on tick #%-2d\n", flashes, i)
		grid.Print()
	}

	fmt.Printf("%d flashed occurred over %d steps", total_flashes, max_ticks)
}
