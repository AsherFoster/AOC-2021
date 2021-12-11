package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func score_autocompleted_token(tok string) int {
	scores := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	return scores[tok]
}

func autocomplete(code string) []string {
	opening_tokens, closing_tokens := get_tokens()
	stack := Stack{}

	for _, token := range strings.Split(code, "") {
		if arr_includes(opening_tokens, token) {
			stack.Push(token)
		} else if arr_includes(closing_tokens, token) {
			opening := stack.Pop()
			// This closing token doesn't match the opening one!
			if get_closing_for(opening) != token {
				panic("this code shoudln't contain invalid chars")
			}
		} else {
			panic("unregocnised token")
		}
	}

	autocomplete_tokens := []string{}

	for len(stack.items) > 0 {
		tok := stack.Pop()
		autocomplete_tokens = append(autocomplete_tokens, get_closing_for(tok))
	}

	return autocomplete_tokens
}

func part_two(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	// Parse the input
	for scanner.Scan() {
		line := scanner.Text()
		_, err := validate(line)
		if err == nil {
			lines = append(lines, line)
		}
	}

	scores := []int{}
	for _, line := range lines {
		tokens := autocomplete(line)

		score := 0
		for _, tok := range tokens {
			score *= 5
			score += score_autocompleted_token(tok)
		}
		// fmt.Printf("Autocompleted %s - score %d\n", tokens, score)
		scores = append(scores, score)
	}

	sort.Ints(scores)
	fmt.Printf("Middle score is %v\n", scores[len(scores)/2])

}
