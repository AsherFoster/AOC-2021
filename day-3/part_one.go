package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sums := [line_width]int{} // Keeps track of which side is more common - +ve for 1s, -ve for 0s

	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {
			if char == '1' {
				sums[i]++
			} else {
				sums[i]--
			}
		}
	}

	var result uint = 0

	for i := 0; i < line_width; i++ {
		var bit uint = 0
		if sums[i] > 0 { // If 1s were more common, this digit is a 1
			bit = 1
		}

		result = result | (bit << (line_width - i - 1))
		fmt.Printf("Iteration #%-2d -- %-5t -- Result is %b\n", i, sums[i] > 0, result)
	}

	fmt.Printf("Result is %v and inverse is %v, product: %v", result, mask^result, result*(mask^result))
}
