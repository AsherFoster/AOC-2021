package main

import (
	"fmt"
	"time"
)

// const path = "sample.txt"
// const tick_count = 18

const path = "input.txt"
const tick_count = 256

func main() {
	// part_one(path)
	start := time.Now()
	part_two(path)

	fmt.Println(time.Since(start))
}
