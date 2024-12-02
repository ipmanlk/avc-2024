package day1

import (
	"bufio"
	"fmt"
	"ipmanlk/avc_2024/utils"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("Day 1")

	list1, list2, err := readInput("inputs/day1")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	solvePart1(list1, list2)
	solvePart2(list1, list2)
}

func solvePart1(list1, list2 []int) {
	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		distance := utils.AbsInt(list1[i] - list2[i])
		totalDistance += distance
	}

	fmt.Println("Part 1: Total Distance:", totalDistance)
}

func solvePart2(list1, list2 []int) {
	// Build a frequency map for list2
	list2Counts := make(map[int]int)
	for _, num := range list2 {
		list2Counts[num]++
	}

	similarityScore := 0
	for _, num := range list1 {
		if count, ok := list2Counts[num]; ok {
			similarityScore += num * count
		}
	}

	fmt.Println("Part 2: Similarity Score:", similarityScore)
}

func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			return nil, nil, fmt.Errorf("invalid input line: %s", line)
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse number: %w", err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse number: %w", err)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return list1, list2, nil
}
