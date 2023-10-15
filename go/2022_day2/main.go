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
	record := make(map[int32]int32, 2)
	// initialize player points
	record[0] = 0
	record[1] = 0

	// guide map for points
	// rock +1, paper +2, scissors +3
	guide := map[string]map[string][]int32{
		"A": {
			"X": []int32{4, 4},
			"Y": []int32{1, 8},
			"Z": []int32{7, 3},
		},
		"B": {
			"X": []int32{8, 1},
			"Y": []int32{5, 5},
			"Z": []int32{2, 9},
		},
		"C": {
			"X": []int32{3, 7},
			"Y": []int32{9, 2},
			"Z": []int32{6, 6},
		},
	}
	fmt.Println(guide)

	input, err := readInput()
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	for _, v := range input {
		_, ok := guide[v[0]]
		if !ok {
			fmt.Println("invalid input found")
			return
		} else {
			mm, ok := guide[v[0]][v[1]]
			if !ok {
				fmt.Println("invalid input found")
				return
			} else {
				record[0] += mm[0]
				record[1] += mm[1]
			}
		}
	}

	fmt.Printf("Final score: Elf {%v}, You {%v}\n", record[0], record[1])
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
