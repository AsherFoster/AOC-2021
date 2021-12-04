package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const board_size = 5

type Board [board_size][board_size]int

func in_arr(haystack []int, needle int) bool {
	for _, test := range haystack {
		if test == needle {
			return true
		}
	}
	return false
}

func consume_draws(scanner *bufio.Scanner) []int {
	scanner.Scan()
	draws := []int{}

	for _, raw := range strings.Split(scanner.Text(), ",") {
		draw, err := strconv.Atoi(raw)
		if err != nil {
			log.Fatal(err)
		}
		draws = append(draws, draw)
	}

	return draws
}

func consume_board(scanner *bufio.Scanner) Board {
	board := Board{}

	for row := 0; row < board_size; row++ {
		scanner.Scan()
		col := 0
		for _, raw := range strings.Split(scanner.Text(), " ") {
			if raw == "" {
				continue // Skip over padding
			}
			num, err := strconv.Atoi(raw)
			if err != nil {
				log.Fatal(err)
			}
			board[row][col] = num
			col++
		}
	}

	return board
}

func (b Board) CheckWin(draws []int) bool {
	// check each row for a win
row: // you bet I just used a label
	for _, row := range b {
		for _, num := range row {
			if !in_arr(draws, num) {
				continue row
			}
		}
		return true
	}
	// check each col for a win
col:
	for col := 0; col < board_size; col++ {
		for _, row := range b {
			num := row[col]
			if !in_arr(draws, num) {
				continue col
			}
		}
		return true
	}

	return false
}

func (b Board) CalcScore(draws []int) int {
	unmarked_sum := 0
	for _, row := range b {
		for _, num := range row {
			if !in_arr(draws, num) {
				unmarked_sum += num
			}
		}
	}

	return unmarked_sum * draws[len(draws)-1]
}

func part_one(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	draws := consume_draws(scanner)

	boards := []Board{}
	for scanner.Scan() { // scanner.Scan gets called twice for each board - this is intentional, it skips over the blank line
		boards = append(boards, consume_board(scanner))
	}

	for i := range draws {
		current_draws := draws[:i+1]
		for b_number, board := range boards {
			if board.CheckWin(current_draws) &&
				// PART 2: quick hack to print out every board that wins, as long as it hadn't won already
				!board.CheckWin(current_draws[:len(current_draws)-1]) {
				// woo, it won!
				fmt.Printf("Hooray, board #%d won with %d points after %d draws\n", b_number, board.CalcScore(current_draws), i+1)
				// return
			}
		}
	}
}
