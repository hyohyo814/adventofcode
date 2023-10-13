package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var calsInt []int
	cals, err := readInput()
	if err != nil {
		fmt.Println("input error")
		return
	}

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

	calsMap := make(map[int]int)
	for i, v := range calsInt {
		calsMap[i] = v
	}
	// get top 3 elves
	top := sortMap(calsMap)
	total := 0
	for _, v := range top {
		total += calsMap[v]
	}
	fmt.Println("total calories held by the top 3 elves are: ", total)
}

func sortMap(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys[:3]
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

