package main

import "fmt"

func arr_includes(arr []int, val int) bool {
	for _, test := range arr {
		if test == val {
			return true
		}
	}
	return false
}

type Grid struct {
	// for "simplicity" the "grid" is represented as a 1d array
	values []int
}

func (g *Grid) increment_all() {
	for i := range g.values {
		g.values[i]++
	}
}

func (g *Grid) Print() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%-2d\n", g.values[i*10:i*10+10])
	}
}

func (g *Grid) get_neighbours(i int) []int {
	x := i % 10
	neighbours := []int{i - 10, i + 10}
	// If this isn't on the leftmost edge, add positions to the left
	if x > 0 {
		neighbours = append(neighbours, i-11, i-1, i+9)
	}
	// If this isn't on the rightmost edge, add positions to the right
	if x < 9 {
		neighbours = append(neighbours, i-9, i+1, i+11)
	}

	// Filter to positions that are in the "grid"
	valid_neighbours := []int{}
	for _, i := range neighbours {
		if i >= 0 && i < len(g.values) {
			valid_neighbours = append(valid_neighbours, i)
		}
	}

	return valid_neighbours
}

// Tick the grid, returning the amount of flashes that occured this tick
func (g *Grid) tick() int {
	// increment all
	g.increment_all()

	flashers := []int{}
	for {
		did_any_flash := false
		// see if any octopuses flash
		for i, energy_level := range g.values {
			if energy_level > 9 && !arr_includes(flashers, i) {
				did_any_flash = true
				flashers = append(flashers, i)
				// increment their neighbours
				neighbours := g.get_neighbours(i)
				for _, pos := range neighbours {
					g.values[pos]++
				}
			}
		}
		// when octopuses stop flashing, we're done
		if !did_any_flash {
			break
		}
	}

	for _, i := range flashers {
		g.values[i] = 0
	}

	// repeat until no octopi flash
	return len(flashers)
}
