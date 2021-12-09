package main

import "fmt"

/*
	Because I'm "sensible", I'll remap the seven segment displays to
	 0000
	1    2
	1    2
	 3333
	4    5
	4    5
	 6666
*/

const segment_count = 7

// This could totally be a bitmap, but that's a little bit less readable - plus we don't need peak efficiency here
type Digit [segment_count]bool

func get_true_digits() []Digit {
	// Ok so like... I hate this... But I think it's kinda needed?
	// There has to be some sort of "picture of a number" -> "number" conversion somewhere
	return []Digit{
		{true, true, true, false, true, true, true},
		{false, false, true, false, false, true, false},
		{true, false, true, true, true, false, true},
		{true, false, true, true, false, true, true},
		{false, true, true, true, false, true, false},
		{true, true, false, true, false, true, true},
		{true, true, false, true, true, true, true},
		{true, false, true, false, false, true, false},
		{true, true, true, true, true, true, true},
		{true, true, true, true, false, true, true},
	}
}

func analyse_frequencies(digits []Digit) [segment_count]int {
	freqs := [segment_count]int{}
	for i := 0; i < segment_count; i++ {
		for _, digit := range digits {
			if digit[i] {
				freqs[i]++
			}
		}
	}
	return freqs
}

func parse_digit(raw string) Digit {
	digit := Digit{}
	for _, char := range raw {
		// Use the ASCII value of the letter to map it to a number: b -> 98; (98 - 97) -> 1;
		digit[char-97] = true
	}

	return digit
}

func count_segments(digit Digit) int {
	count := 0
	for _, segment := range digit {
		if segment {
			count++
		}
	}
	return count
}

func decode_with_mapping(encoded Digit, mappings [segment_count]int) int {
	decoded := Digit{}

	for key, val := range encoded {
		decoded[mappings[key]] = val
	}

	for i, test_digit := range get_true_digits() {
		if decoded == test_digit {
			return i
		}
	}

	panic("Decoding failed!")
}

// Determnine a mapping between encoded segment and true segtment
func magically_determine_mapping(examples []Digit) [segment_count]int {
	mapping := [segment_count]int{-1, -1, -1, -1, -1, -1, -1} // -1 as null

	// We can figure out segment 0 by finding the segment that's on in 4, but not in 1
	// 2 segments -> 1
	// 3 segments -> 7
	// 4 segments -> 4
	// 7 segments -> 8
	var encoded_one, encoded_four, encoded_seven Digit

	for _, digit := range examples {
		count := count_segments(digit)
		if count == 2 {
			encoded_one = digit
		} else if count == 3 {
			encoded_seven = digit
		} else if count == 4 {
			encoded_four = digit
		}
	}
	for i, on := range encoded_seven {
		if on && !encoded_one[i] {
			mapping[i] = 0
			break
		}
	}

	// Certain segments are on in N amount of digits. By analysing how often segments are on, we can decoded 3 segments
	encoded_frequencies := analyse_frequencies(examples)
	for i, freq := range encoded_frequencies {
		// Using the "true frequency map"
		// [8 6 8 7 4 9 7]
		if freq == 8 {
			// Segment 0 and 2 are both on in 8 digits. But, since we already know which is Segment 0, the other must be Segment 2
			if mapping[i] != 0 {
				mapping[i] = 2
			}
		} else if freq == 6 {
			mapping[i] = 1
		} else if freq == 4 {
			mapping[i] = 4
		} else if freq == 9 {
			mapping[i] = 5
		}
	}

	// Finally, we need to figure out which of Segment 3 and 6 is which
	// The unmapped segment in the digit 4 must be Segment 3
	for i, on := range encoded_four {
		if on && mapping[i] == -1 {
			mapping[i] = 3
		}
	}

	// ...leaving the last one to be Segment 6
	for i, target := range mapping {
		if target == -1 {
			mapping[i] = 6
		}
	}

	fmt.Printf("Decoded to map %d\n", mapping)
	return mapping
}
