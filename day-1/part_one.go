package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	last, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	larger_count := 0
	for scanner.Scan() {
		l := scanner.Text()
		i, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}

		if i > last {
			larger_count++
		}

		last = i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total, %v\n", larger_count)
}
