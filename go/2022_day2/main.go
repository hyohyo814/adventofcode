package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	// opponent: A rock, B paper, C scissors
	// you: X rock, Y paper, Z scissors

	// record keeping map for players
	record := make(map[int]int, 2)
	// initialize player points
	record[0] = 0
	record[1] = 0

	// guide map for points
	// rock +1, paper +2, scissors +3
	guide := map[string]map[string][]int{
		"A": {
			"X": []int{4, 4},
			"Y": []int{1, 8},
			"Z": []int{7, 3},
		},
		"B": {
			"X": []int{8, 1},
			"Y": []int{5, 5},
			"Z": []int{2, 9},
		},
		"C": {
			"X": []int{3, 7},
			"Y": []int{9, 2},
			"Z": []int{6, 6},
		},
	}
	fmt.Println(guide)

	input, err := readInput()
	if err != nil {
		fmt.Errorf("error reading input")
		return
	}

	for _, v := range input {
		_, ok := guide[v[0]]
		if !ok {
			fmt.Errorf("invalid input")
			return
		}

	}
}

// Function to read and format input data from text file
func readInput() ([][]string, error) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	formattedPairs := make([][]string, 0, len(lines))
	for _, v := range lines {
		pairs := strings.Split(v, " ")
		if len(pairs) != 2 {
			continue
		}
		formattedPairs = append(formattedPairs, []string{pairs[0], pairs[1]})
	}

	return formattedPairs, nil
}

