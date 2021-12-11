package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Either go doesn't include this, or I don't
func arr_includes(arr [][2]int, el [2]int) bool {
	for _, test := range arr {
		if test == el {
			return true
		}
	}
	return false
}

//   0
// 1 X 2
//   3
func get_explorable(grid [][]int, x, y int) [][2]int {
	explorable := [][2]int{}
	if y > 0 && grid[y-1][x] < 9 {
		explorable = append(explorable, [2]int{x, y - 1})
	}
	if x > 0 && grid[y][x-1] < 9 {
		explorable = append(explorable, [2]int{x - 1, y})
	}
	if x < len(grid[y])-1 && grid[y][x+1] < 9 {
		explorable = append(explorable, [2]int{x + 1, y})
	}
	if y < len(grid)-1 && grid[y+1][x] < 9 {
		explorable = append(explorable, [2]int{x, y + 1})
	}
	return explorable
}

func get_basin_size(grid [][]int, start_loc [2]int) int {
	// So I think this needs to do a flood-fill esque algorithm to either
	// Flood until reaching a perimeter of 9s
	// Flood until the value goes down again

	locations := [][2]int{start_loc}
	// len(locations) will keep increasing until we run out of locations to explore
	// therefore, this will keep looping as the queue increases, until we catch up to the end
	for i := 0; i < len(locations); i++ {
		next_loc := locations[i]

		// get potential locations from a position (non-9 adjacents)
		for _, new_loc := range get_explorable(grid, next_loc[0], next_loc[1]) {
			// if this location is new, add it to the queue
			if !arr_includes(locations, new_loc) {
				locations = append(locations, new_loc)
			}
		}
	}

	fmt.Printf("Basin {%d, %d} has locations %v\n", start_loc[0], start_loc[1], locations)
	return len(locations)
}

func part_two(path string) {
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

	basins := []int{}
	for y, row := range grid {
		for x := range row {
			if is_less_than_adjacents(grid, x, y) {
				basins = append(basins, get_basin_size(grid, [2]int{x, y}))
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	fmt.Printf("Basin product is %d\n", basins[0]*basins[1]*basins[2])
}
