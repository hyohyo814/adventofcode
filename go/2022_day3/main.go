package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		items - 1 of 2 compartments
		A given rucksack always has the same number of items in each of its two comps
	*/

	result := 0
	priority := priorityCreate()

	input, err := readInput()
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	for i, v := range input {
		strarr := strings.Split(v, "")
		left := strarr[:len(strarr)/2]
		right := strarr[len(strarr)/2:]

		leftguide := make(map[string]int)
		rightguide := make(map[string]int)

		var target string
		for i := range left {
			if _, ok := leftguide[left[i]]; !ok {
				leftguide[left[i]] = 1
			}
			if _, ok := rightguide[right[i]]; !ok {
				rightguide[right[i]] = 1
			}

			if _, ok := leftguide[right[i]]; ok {
				target = right[i]
				continue
			}
			if _, ok := rightguide[left[i]]; ok {
				target = left[i]
				continue
			}
		}
		fmt.Printf("The type found in both compartments for rucksack %v is: %v\n", i, target)
		lookup := []rune(target)
		if mm, ok := priority[lookup[0]]; ok {
			result += mm
		}
	}
	fmt.Printf("The sum of priorities: %v\n", result)
}

// Function to read and format input data from text file
func readInput() ([]string, error) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	for i, v := range lines {
		if v == "" {
			lines = lines[:i]
		}
	}

	return lines, nil
}

func priorityCreate() map[rune]int {
	alphabetMap := make(map[rune]int)

	// Populate the map with lowercase letters (a to z)
	for i, letter := 'a', 1; i <= 'z'; i, letter = i+1, letter+1 {
		alphabetMap[i] = letter
	}

	// Populate the map with uppercase letters (A to Z)
	for i, letter := 'A', 27; i <= 'Z'; i, letter = i+1, letter+1 {
		alphabetMap[i] = letter
	}

	return alphabetMap
}
