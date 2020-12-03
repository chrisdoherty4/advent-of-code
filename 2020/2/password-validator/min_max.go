package main

import (
	"aoc-2/policy"
	"aoc-2/policy/requirement"
)

type MinMaxCalculator struct {
	valid   int
	invalid int
}

func NewMinMaxCalculator() MinMaxCalculator {
	return MinMaxCalculator{}
}

func (m *MinMaxCalculator) Calculate(num1, num2 int, char rune, word string) {
	minMaxPolicy := policy.NewStringPolicy()
	minMaxCharRequirement := requirement.NewMinMaxCharCount(char, num1, num2)
	minMaxPolicy.AddRequirement(minMaxCharRequirement)

	switch minMaxPolicy.IsValid(word) {
	case true:
		m.valid++
	case false:
		m.invalid++
	}
}

func (m MinMaxCalculator) Valid() int {
	return m.valid
}

func (m MinMaxCalculator) Invalid() int {
	return m.invalid
}

func (m MinMaxCalculator) Total() int {
	return m.valid + m.invalid
}
