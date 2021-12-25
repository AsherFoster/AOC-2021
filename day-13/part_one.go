package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var instruction_regex = regexp.MustCompile("fold along (x|y)=([0-9]+)")

type Grid struct {
	x_size int
	y_size int
	values [][]bool
}

func (g Grid) Print() {
	for _, row := range g.values {
		for _, cell := range row {
			if cell {
				fmt.Print("█")
				// fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (g *Grid) EnsureSize(x_size, y_size int) {
	if x_size > g.x_size {
		g.x_size = x_size
		// resize every row

		for y, row := range g.values {
			// new_row := row[:g.x_size]
			new_row := make([]bool, g.x_size)
			copy(new_row, row)

			g.values[y] = new_row
		}
	}

	if y_size > g.y_size {
		g.y_size = y_size

		new_grid := make([][]bool, g.y_size)
		copy(new_grid, g.values)

		// add extra rows
		for row := len(g.values); row < g.y_size; row++ {
			new_grid[row] = make([]bool, g.x_size)
		}

		g.values = new_grid
	}
}

func (grid Grid) Fold(axis string, fold_point int) Grid {
	new_grid := Grid{}

	if axis == "y" {
		new_grid.EnsureSize(grid.x_size, fold_point)

		for y, row := range new_grid.values {

			// Every row has the top half row
			copy(new_grid.values[y], grid.values[y])

			// And overlap the row from the bottom half
			if fold_point*2-y < grid.y_size {
				for x := range row {
					row[x] = row[x] || grid.values[(fold_point*2)-y][x]
				}
			}
		}
	} else {
		new_grid.EnsureSize(fold_point, grid.y_size)

		// add the left half
		for y, row := range new_grid.values {
			copy(new_grid.values[y], grid.values[y])

			// reverse add the right half
			for x := range row {
				if fold_point*2-x < grid.x_size {
					row[x] = row[x] || grid.values[y][(fold_point*2)-x]
				}
			}
		}
	}

	return new_grid
}

func (grid Grid) Count() int {
	count := 0
	for _, row := range grid.values {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}
	return count
}

func parse_int_but_mean_it(raw string) int {
	val, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := Grid{}
	// Parse the input. We use a grid.ensure_size function to ensure the grid is large enough to hold this coord
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // once we reach a blank line, it's instruction time!
			break
		}
		parts := strings.Split(line, ",")
		x, y := parse_int_but_mean_it(parts[0]), parse_int_but_mean_it(parts[1])

		grid.EnsureSize(x+1, y+1)
		grid.values[y][x] = true
	}

	fmt.Printf("Start grid - %dx%d. %d cells\n", grid.x_size, grid.y_size, grid.Count())
	// grid.Print()

	for scanner.Scan() {
		segments := instruction_regex.FindStringSubmatch(scanner.Text())
		axis, pos := segments[1], parse_int_but_mean_it(segments[2])

		fmt.Printf("\nFolding along %v=%v\n", axis, pos)
		grid = grid.Fold(axis, pos)

		fmt.Printf("Grid is now %dx%d. %d cells\n", grid.x_size, grid.y_size, grid.Count())
		// grid.Print()
		// break // PART ONE - break immediately
	}
	grid.Print()
}
