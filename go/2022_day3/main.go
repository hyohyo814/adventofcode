package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	/*
		items - 1 of 2 compartments
		A given rucksack always has the same number of items in each of its two comps
	*/

	input, err := readInput()
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	for _, v := range input {
		strarr := strings.Split(v, "")
		fmt.Println(strarr)
		for i := range strarr {
			tmpleft := strarr[:i+1]	
			tmpright := strarr[i+1:]	
			fmt.Println("lhs: ", tmpleft)
			fmt.Println("rhs: ", tmpright)
			
			// check for matches using map between two slices
		}
	}
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
