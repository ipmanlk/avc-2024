package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Solve() {
	fmt.Println("\nDay 3")

	solvePart1()
	solvePart2()
}

func solvePart1() {
	file, _ := os.ReadFile("inputs/day3")
	content := string(file)

	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)

	matches := r.FindAllStringSubmatch(content, -1)

	total := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		total += num1 * num2
	}

	fmt.Println("Part 1: Total:", total)
}

func solvePart2() {
	file, _ := os.ReadFile("inputs/day3")
	content := string(file)

	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)
	matches := r.FindAllStringSubmatch(content, -1)

	total := 0
	multiply := true
	for _, match := range matches {
		if match[0] == "don't()" {
			multiply = false
			continue
		}

		if match[0] == "do()" {
			multiply = true
			continue
		}

		if !multiply {
			continue
		}

		if multiply {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			total += num1 * num2
		}
	}

	fmt.Println("Part 2: Total:", total)
}
