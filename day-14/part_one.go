package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func insert_random_letters_everywhere(original string, rules map[string]string) string {
	new := ""

	for i, char := range original[:len(original)-1] {
		new += string(char)
		new += rules[original[i:i+2]]
	}
	new += string(original[len(original)-1])
	return new
}

func count_letters(str string) map[string]int {
	counts := map[string]int{}

	for _, char := range str {
		counts[string(char)]++
	}
	return counts
}

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Grab the starting polymer from the first line
	scanner.Scan()
	polymer := scanner.Text()

	scanner.Scan()

	rules := map[string]string{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		rules[parts[0]] = parts[1] // AB: C
	}

	fmt.Printf("Starting with polymer %s, and %d rules\n", polymer, len(rules))
	for round := 0; round < 40; round++ {
		polymer = insert_random_letters_everywhere(polymer, rules)
		fmt.Printf("After round %-2d (%-2d long)\n", round+1, len(polymer))
		fmt.Println(count_letters(polymer))
	}
}
