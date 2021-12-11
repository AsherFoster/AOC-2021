package main

type Stack struct {
	items []string
}

func (s *Stack) Pop() string {
	popped := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]

	return popped
}
func (s *Stack) Push(val string) {
	s.items = append(s.items, val)
}
func arr_includes(arr []string, val string) bool {
	for _, test := range arr {
		if test == val {
			return true
		}
	}
	return false
}

func get_tokens() ([]string, []string) {
	return []string{"(", "[", "{", "<"}, []string{")", "]", "}", ">"}
}
func get_closing_for(opening string) string {
	opening_tokens, closing_tokens := get_tokens()

	for i, tok := range opening_tokens {
		if tok == opening {
			return closing_tokens[i]
		}
	}

	panic("unrecognised opening token" + opening)
}
