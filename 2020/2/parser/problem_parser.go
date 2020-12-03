package parser

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type InputDataHandler func(num1, num2 int, char rune, word string)

// line examples:
//	1-2 a: aaabbb
//	2-4 b: abababab
//	2-9 c: asdevcsdadwqvr
type InputDataParser struct {
	inputDataHandler InputDataHandler
}

func NewInputDataParser(inputDataHandler InputDataHandler) InputDataParser {
	return InputDataParser{
		inputDataHandler: inputDataHandler,
	}
}

func (m InputDataParser) Parse(file *os.File) error {
	reader := bufio.NewReader(file)

	const regex = "^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$"
	lineParsingRegex := regexp.MustCompile(regex)

	parseLine := func(line string) error {
		groups := lineParsingRegex.FindStringSubmatch(line)
		if len(groups) != 5 {
			return errors.New("Not enough groups found")
		}

		num1, err := strconv.Atoi(groups[1])
		if err != nil {
			return err
		}

		num2, err := strconv.Atoi(groups[2])
		if err != nil {
			return err
		}

		char := rune(groups[3][0])
		word := groups[4]

		m.inputDataHandler(
			num1,
			num2,
			char,
			word,
		)

		return nil
	}

	if err := m.forEachLine(reader, parseLine); err != nil {
		return err
	}

	return nil
}

type lineFunc func(line string) error

func (m InputDataParser) forEachLine(reader *bufio.Reader, fn lineFunc) error {
	for {
		line, err := reader.ReadString('\n')

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}

		line = strings.TrimRight(line, "\n")
		if err := fn(line); err != nil {
			return err
		}
	}
}
