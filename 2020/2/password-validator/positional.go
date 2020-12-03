package main

import (
	"aoc-2/policy"
	"aoc-2/policy/requirement"
)

type PositionalCharCalculator struct {
	valid   int
	invalid int
}

func NewPositionalCharCalculator() PositionalCharCalculator {
	return PositionalCharCalculator{}
}

func (m *PositionalCharCalculator) Calculate(num1, num2 int, char rune, word string) {
	positionalPolicy := policy.NewStringPolicy()
	positionalCharRequirement := requirement.NewPositionalChar(
		char,
		[]int{num1, num2},
	)
	positionalPolicy.AddRequirement(positionalCharRequirement)

	switch positionalPolicy.IsValid(word) {
	case true:
		m.valid++
	case false:
		m.invalid++
	}
}

func (m PositionalCharCalculator) Valid() int {
	return m.valid
}

func (m PositionalCharCalculator) Invalid() int {
	return m.invalid
}

func (m PositionalCharCalculator) Total() int {
	return m.valid + m.invalid
}
