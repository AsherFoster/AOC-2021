package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}
type Line [2]Coord

func normalise(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	} else {
		return a
	}
}

func parse_int(raw string) int {
	i, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
func parse_coord(raw string) Coord {
	parts := strings.Split(raw, ",")
	return Coord{
		x: parse_int(parts[0]),
		y: parse_int(parts[1]),
	}
}

func part_one_and_two(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []Line{}

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")
		lines = append(lines, Line{
			parse_coord(coords[0]),
			parse_coord(coords[1]),
		})
	}

	// Somehow count intersects. Possible ideas:
	// For each line, see if it intersects with previous lines - this would be O(n!) I think, not optimal.
	// - Idk how to to count the amount of overlapping cells without building a graph anyway
	// - We don't want to count cells that have already had an overlap
	// Build up a graph, plotting each line and incrementing the cell count. Then, count then amount of points on the graph >2
	// - It's what the example says

	// First idea is simplest to write, probably bad perf tho
	max_x := 0
	max_y := 0

	for _, line := range lines {
		for _, c := range line {
			if c.x > max_x {
				max_x = c.x
			}
			if c.y > max_y {
				max_y = c.y
			}
		}
	}

	fmt.Printf("Constructing chart %dx%d\n", max_x, max_y)
	chart := make([][]int, max_y+1)
	for i := range chart {
		chart[i] = make([]int, max_x+1)
	}

	for _, line := range lines {
		// if line[0].x == line[1].x || line[0].y == line[1].y { // Filter for part one, to only include straight lines
		// given x1, y1 and x2, y2 - how do we plot the points in between??
		x := line[0].x
		y := line[0].y
		chart[y][x]++

		for x != line[1].x || y != line[1].y {
			// Increment the numbers in the direction they're heading
			x += normalise(line[1].x - line[0].x)
			y += normalise(line[1].y - line[0].y)

			// fmt.Printf("%v - Incrementing %v/%v - %d\n", line, x, y, chart[y][x]+1)
			chart[y][x]++
		}
	}
	// }

	count := 0
	for _, row := range chart {
		for _, cell := range row {
			if cell > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}
