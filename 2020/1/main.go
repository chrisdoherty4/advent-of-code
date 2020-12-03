package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrInsufficientNumbers = errors.New("Insufficient numbers")
	ErrNoAddends           = errors.New("No addends could be found")
)

func find2AddendsOf2020(numbers []int) (int, int, error) {
	if len(numbers) < 2 {
		return 0, 0, ErrInsufficientNumbers
	}

	const targetSum = 2020
	addendDeltaCache := make(map[int]int)

	for i, number := range numbers {
		delta := targetSum - number
		if numbersIndex, ok := addendDeltaCache[delta]; ok {
			return number, numbers[numbersIndex], nil
		}

		addendDeltaCache[number] = i
	}

	return 0, 0, ErrNoAddends
}

func find3AddendsOf2020(numbers []int) (int, int, int, error) {
	if len(numbers) < 3 {
		return 0, 0, 0, ErrInsufficientNumbers
	}

	sort.Ints(numbers)
	const targetSum = 2020

	for i := 0; i < len(numbers)-2; i++ {
		left := i + 1
		right := len(numbers) - 1

		for left < right {
			sum := numbers[i] + numbers[left] + numbers[right]
			switch {
			case sum == targetSum:
				return numbers[i], numbers[left], numbers[right], nil
			case sum > targetSum:
				right--
			default:
				left++
			}
		}
	}

	return 0, 0, 0, ErrNoAddends
}

func product(numbers []int) int {
	product := 1
	for _, number := range numbers {
		product *= number
	}

	return product
}

func readNumbers(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	numbers := []int{}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		switch {
		case err == io.EOF:
			return numbers, nil
		case err != nil:
			return nil, err
		}

		line = strings.TrimRight(line, "\n")
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expecting CLI param 1 to be a path")
	}

	// Expect the build to occur in the current working directory
	numbers, err := readNumbers(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// addend1, addend2, err := find2AddendsOf2020(numbers)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// product := product(addend1, addend2)
	// fmt.Println(product)

	addend1, addend2, addend3, err := find3AddendsOf2020(numbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(product([]int{addend1, addend2, addend3}))
}
