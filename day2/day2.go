package day2

import (
	"bufio"
	"fmt"
	"ipmanlk/avc_2024/utils"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	fmt.Println("\nDay 2")

	reports := loadReports("inputs/day2")
	solvePart1(reports)
	solvePart2(reports)
}

func solvePart1(reports [][]int) {
	safeReports := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}

	fmt.Printf("Part 1: Safe Reports: %d\n", safeReports)
}

func solvePart2(reports [][]int) {
	safeReports := 0

	for _, report := range reports {
		if isReportSafe(report) || checkWithProblemDampener(report) {
			safeReports++
		}
	}

	fmt.Printf("Part 2: Safe Reports: %d\n", safeReports)
}

func isReportSafe(report []int) bool {
	increasing := report[0] < report[1]

	for i := 0; i < len(report)-1; i++ {
		if (increasing && report[i] >= report[i+1]) || (!increasing && report[i] <= report[i+1]) {
			return false
		}

		diff := utils.AbsInt(report[i] - report[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func checkWithProblemDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		newNums := make([]int, len(report)-1)
		copy(newNums, report[:i])
		copy(newNums[i:], report[i+1:])
		if isReportSafe(newNums) {
			return true
		}
	}
	return false
}

func loadReports(filePath string) [][]int {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reports := make([][]int, 0)
	for scanner.Scan() {
		report := parseReport(scanner.Text())
		reports = append(reports, report)
	}

	return reports
}

func parseReport(line string) []int {
	parts := strings.Fields(line)
	nums := make([]int, len(parts))

	for i, part := range parts {
		nums[i], _ = strconv.Atoi(part)
	}

	return nums
}
