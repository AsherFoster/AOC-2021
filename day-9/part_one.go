package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func is_less_than_adjacents(grid [][]int, x, y int) bool {
	h := grid[y][x]

	if y > 0 && grid[y-1][x] <= h {
		return false
	}
	if x > 0 && grid[y][x-1] <= h {
		return false
	}
	if x < len(grid[y])-1 && grid[y][x+1] <= h {
		return false
	}
	if y < len(grid)-1 && grid[y+1][x] <= h {
		return false
	}
	return true
}

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := [][]int{}

	// Parse the input
	for scanner.Scan() {
		row := []int{}

		for _, char := range scanner.Text() {
			h, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, h)
		}

		grid = append(grid, row)
	}

	local_minimums := [][2]int{}
	for y, row := range grid {
		for x := range row {
			if is_less_than_adjacents(grid, x, y) {
				local_minimums = append(local_minimums, [2]int{x, y})
			}
		}
	}

	risk := 0
	for _, min := range local_minimums {
		fmt.Printf("Minumim at {%2d, %2d} - height %d\n", min[0], min[1], grid[min[1]][min[0]])
		risk += 1 + grid[min[1]][min[0]]
	}

	fmt.Printf("Total risk is %d", risk)
}
