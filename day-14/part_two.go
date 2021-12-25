package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func tick_pairs(pairs map[string]int, rules map[string]string) map[string]int {
	new := map[string]int{}

	for pair, count := range pairs {
		insert := rules[pair]
		if insert != "" {
			// When a pair is split in half, we don't want it to count for the original anymore
			new[string(pair[0])+insert] += count
			new[insert+string(pair[1])] += count
		} else {
			new[pair] += count
		}
	}

	return new
}

func count_pairs_dict(pairs map[string]int) map[string]int {
	counts := map[string]int{}

	for pair, count := range pairs {
		counts[string(pair[0])] += count
		counts[string(pair[1])] += count
	}
	return counts
}

func part_two(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Grab the starting polymer from the first line
	scanner.Scan()

	// Parse the polymer string into a set of pairs
	pairs := map[string]int{}
	raw_polymer := scanner.Text()
	for i := range raw_polymer[:len(raw_polymer)-1] {
		pairs[raw_polymer[i:i+2]]++
	}

	scanner.Scan()

	rules := map[string]string{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		rules[parts[0]] = parts[1] // AB: C
	}

	for round := 0; round < 40; round++ {
		pairs = tick_pairs(pairs, rules)
		fmt.Printf("Round %d, %v\n", round+1, count_pairs_dict(pairs))
	}

	counts := count_pairs_dict(pairs)

	for pair, count := range counts {
		fmt.Printf("%s: %d\n", pair, (count+1)/2) // rely on integer division to ensure odd numbers are rounded up
	}
}
