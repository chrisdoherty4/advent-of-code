package main

import (
	"fmt"
	"log"
	"os"

	"aoc-2/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected param 1 to be a path")
	}

	//calculator := NewMinMaxCalculator()
	calculator := NewPositionalCharCalculator()
	inputDataParser := parser.NewInputDataParser(calculator.Calculate)

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := inputDataParser.Parse(file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Passwords: %v\n", calculator.Total())
	fmt.Printf("ValidPasswords: %v\n", calculator.Valid())
}
