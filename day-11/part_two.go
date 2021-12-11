package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part_two(path string) {
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
	for i := 0; i < 10000; i++ {
		flashes := grid.tick()
		if flashes == 100 {
			fmt.Printf("Tick %d, all flashed!", i)
			grid.Print()
			break
		}
	}
}
