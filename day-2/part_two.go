package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FancyPosition struct {
	depth int
	x     int
	aim   int
}

func (p FancyPosition) move(input string) FancyPosition {
	command, param := parse_command(input)

	switch command {
	case "forward":
		return FancyPosition{p.depth + (p.aim * param), p.x + param, p.aim}
	case "up":
		return FancyPosition{p.depth, p.x, p.aim - param}
	case "down":
		return FancyPosition{p.depth, p.x, p.aim + param}
	default:
		log.Fatalf("Unrecognised command %s\n", command)
		return FancyPosition{}
	}
}

func part_two(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	pos := FancyPosition{}
	for scanner.Scan() {
		pos = pos.move(scanner.Text())
	}

	fmt.Println(pos, pos.depth*pos.x)
}
