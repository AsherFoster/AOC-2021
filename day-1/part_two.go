package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const window_size = 3

func get_int(scanner *bufio.Scanner) int {
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func part_two(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	last_window := make([]int, window_size)

	// Initialise window
	for i := 0; i < window_size; i++ {
		scanner.Scan()
		last_window[i] = get_int(scanner)
	}

	larger_count := 0
	for scanner.Scan() {
		i := get_int(scanner)
		new_window := append(last_window[1:], i)

		if sum(new_window) > sum(last_window) {
			larger_count++
		}
		last_window = new_window
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total, %v\n", larger_count)
}
