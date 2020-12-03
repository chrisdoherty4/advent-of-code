package requirement

import (
	"sort"
)

type PositionalChar struct {
	char      rune
	positions []int
}

func NewPositionalChar(char rune, positions []int) PositionalChar {
	// copy the slice so we don't upset anything else using it externally.
	positionsCopy := make([]int, len(positions), len(positions))
	copy(positionsCopy, positions)
	sort.Ints(positionsCopy)

	return PositionalChar{
		char:      char,
		positions: positionsCopy,
	}
}

func (m PositionalChar) Validate(stringToValidate string) bool {
	return m.hasOnePositionalChar(stringToValidate)
}

func (m *PositionalChar) hasOnePositionalChar(stringToValidate string) bool {
	var positionalCharMatches int
	for _, position := range m.positions {
		// has enough chars to check position
		if position > len(stringToValidate) {
			continue
		}

		// positional char matches
		if rune(stringToValidate[position-1]) == m.char {
			positionalCharMatches++
		}
	}

	return positionalCharMatches == 1 //cba putting the magic number somewhere
}
