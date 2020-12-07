package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	treeSymbol      = rune('#')
	clearPathSymbol = rune('.')
)

func main() {
	file, err := os.Open("fallline.txt")
	failOnError(err)

	slope := NewSlopeFileHandler(file, rune('#'))
	pathMapper := NewPathMapper(slope)

	result := 1
	for _, p := range []FallPath{
		{
			Right: 1,
			Down:  1,
		},
		{
			Right: 3,
			Down:  1,
		},
		{
			Right: 5,
			Down:  1,
		},
		{
			Right: 7,
			Down:  1,
		},
		{
			Right: 1,
			Down:  2,
		},
	} {
		treesHit := pathMapper.CalculateTreesHit(p)
		result *= treesHit
		slope.Reset()
		pathMapper.Reset()
		fmt.Printf("Trees hit: %v;\tResult: %v\n", treesHit, result)
	}
}

func failOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getLineLengthAndDiscardLine(reader *bufio.Reader) int {
	line, err := reader.ReadString('\n')
	failOnError(err)

	line = strings.TrimRight(line, "\n")

	return len(line)
}
