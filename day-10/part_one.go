package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func score_bad_token(tok string) int {
	scores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	return scores[tok]
}

func validate(code string) (string, error) {
	opening_tokens, closing_tokens := get_tokens()
	stack := Stack{}

	for _, token := range strings.Split(code, "") {
		if arr_includes(opening_tokens, token) {
			stack.Push(token)
		} else if arr_includes(closing_tokens, token) {
			opening := stack.Pop()
			expected := get_closing_for(opening)
			// This closing token doesn't match the opening one!
			if expected != token {
				return token, fmt.Errorf("expected closing %v, found %v instead", expected, token)
			}
		} else {
			return token, fmt.Errorf("unregocnised token %v", token)
		}
	}

	return "", nil
}

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	// Parse the input
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	bad_toks := []string{}
	for i, line := range lines {
		token, err := validate(line)
		if err != nil {
			bad_toks = append(bad_toks, token)
			fmt.Printf("Error on line %d: %v\n", i, err)
		}
	}

	score := 0
	for _, tok := range bad_toks {
		score += score_bad_token(tok)
	}

	fmt.Printf("High score! Invalid parser score %d", score)
}
