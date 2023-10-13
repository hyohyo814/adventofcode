package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cals, err := readInput()
	if err != nil {
		fmt.Println("input error")
		return
	}

	var calsInt []int

	j := 0
	buffer := 0
	for _, v := range cals {
		if v != "" {
			calVal, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("err converting str")
			}

			buffer += calVal
		} else if v == "" {
			calsInt = append(calsInt, buffer)
			j++
			buffer = 0
		}
	}


	maxVal, maxIndex := 0, 0
	for i, v := range calsInt {
		if v > maxVal {
			// convert from zero index
			maxIndex = i + 1
			maxVal = v
		}
	}

	fmt.Printf("The elf with the most calories held is elf %v with %v calories\n", maxIndex, maxVal)
}

// Function to read and format input data from text file
func readInput() ([]string, error) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	return lines, nil
}

