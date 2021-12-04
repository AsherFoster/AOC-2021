package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	depth int
	x     int
}

func parse_command(input string) (string, int) {
	parts := strings.Split(input, " ")

	param, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	return parts[0], param
}

func (p Position) move(input string) Position {
	command, param := parse_command(input)

	switch command {
	case "forward":
		return Position{p.depth, p.x + param}
	case "up":
		return Position{p.depth - param, p.x}
	case "down":
		return Position{p.depth + param, p.x}
	default:
		log.Fatalf("Unrecognised command %s\n", command)
		return Position{}
	}
}

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	pos := Position{0, 0}
	for scanner.Scan() {
		pos = pos.move(scanner.Text())
	}

	fmt.Println(pos, pos.depth*pos.x)
}
