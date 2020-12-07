package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func NewSlopeFileHandler(
	file *os.File,
	treeSymbol rune,
) *SlopeFileHandler {
	reader := bufio.NewReader(file)
	// assume they've given us something sensible
	line, _ := readLineWithoutNewlineCharOrFail(reader)

	return &SlopeFileHandler{
		file:        file,
		reader:      reader,
		currentLine: line,
		lineLength:  len(line),
		treeSymbol:  treeSymbol,
	}
}

type SlopeFileHandler struct {
	file        *os.File
	reader      *bufio.Reader
	currentLine string
	lineIndex   int
	lineLength  int
	treeSymbol  rune
	eof         bool
}

func (s *SlopeFileHandler) Right(count int) {
	if s.eof {
		return
	}

	s.lineIndex = (s.lineIndex + count) % s.lineLength
}

func (s *SlopeFileHandler) Down(count int) {
	if s.eof {
		return
	}

	for ; count > 0; count-- {
		s.currentLine, s.eof = readLineWithoutNewlineCharOrFail(s.reader)
		if s.eof {
			return
		}
	}
}

func (s *SlopeFileHandler) AtBottom() bool {
	return s.eof
}

func (s *SlopeFileHandler) HaveHitTree() bool {
	if s.eof {
		return false
	}

	return rune(s.currentLine[s.lineIndex]) == s.treeSymbol
}

func (s *SlopeFileHandler) Reset() {
	s.file.Seek(0, 0)
	s.reader.Reset(s.file)
	s.currentLine, _ = readLineWithoutNewlineCharOrFail(s.reader)
	s.lineIndex = 0
	s.eof = false
}

func readLineWithoutNewlineCharOrFail(
	reader *bufio.Reader,
) (string, bool) {
	line, err := reader.ReadString('\n')

	switch {
	case err == io.EOF:
		return "", true
	case err != nil:
		// Just bomb out, we don't care
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.TrimRight(line, "\n"), false
}
