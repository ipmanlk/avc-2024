package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve() {
	fmt.Println("\nDay 4")

	solvePart1()
	solvePart2()
}

// =============================================================================
//	Part 1
// =============================================================================

// Each direction is represented by a pair of integers (dx, dy):
//
//	{0, 1}   -> right
//	{0, -1}  -> left
//	{1, 0}   -> down
//	{-1, 0}  -> up
//	{1, 1}   -> diagonal down-right
//	{-1, -1} -> diagonal up-left
//	{1, -1}  -> diagonal down-left
//	{-1, 1}  -> diagonal up-right
var p1Directions = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func solvePart1() {
	matrix := readInput("inputs/day4")

	searchWord := "XMAS"
	count := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			for _, dir := range p1Directions {
				if isWord(matrix, searchWord, i, j, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	fmt.Printf("Part 1: Count is %d\n", count)
}

// Checks if the word exists starting at (x, y) in the given direction
func isWord(matrix [][]rune, word string, x, y, dx, dy int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < len(word); i++ {
		// Check if the position is out of bounds
		if x < 0 || y < 0 || x >= rows || y >= cols {
			return false
		}
		// Check if the current character matches
		if matrix[x][y] != rune(word[i]) {
			return false
		}
		// Move to the next character in the specified direction
		x += dx
		y += dy
	}
	return true
}

// =============================================================================
//	Part 2
// =============================================================================

// Each direction is represented by a key and a pair of integers:
//
//	"Q": {-1, -1} -> top-left (NW)
//	"R": {-1,  1} -> top-right (NE)
//	"Z": { 1, -1} -> bottom-left (SW)
//	"C": { 1,  1} -> bottom-right (SE)
var p2Directions = map[string][2]int{
	"Q": {-1, -1},
	"R": {-1, 1},
	"Z": {1, -1},
	"C": {1, 1},
}

func solvePart2() {
	charArray := readInput("inputs/day4")
	rows := len(charArray)
	cols := len(charArray[0])

	totalCountMAS := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if countMAS(charArray, row, col) {
				totalCountMAS++
			}
		}
	}

	fmt.Printf("Part 2: Count is %d\n", totalCountMAS)
}

// Checks if a valid "MAS" pattern exists at the given position
func countMAS(charArray [][]rune, row, col int) bool {
	rows := len(charArray)
	cols := len(charArray[0])

	// Ensure the center is 'A'
	if charArray[row][col] != 'A' {
		return false
	}

	// Check top-left and bottom-right diagonals
	topLeftRow, topLeftCol, validTL := nextCheckpoint(row, col, "Q", rows, cols)
	bottomRightRow, bottomRightCol, validBR := nextCheckpoint(row, col, "C", rows, cols)
	if !validTL || !validBR {
		return false
	}
	if charArray[topLeftRow][topLeftCol] == 'M' {
		if charArray[bottomRightRow][bottomRightCol] != 'S' {
			return false
		}
	} else if charArray[topLeftRow][topLeftCol] == 'S' {
		if charArray[bottomRightRow][bottomRightCol] != 'M' {
			return false
		}
	} else {
		return false
	}

	// Check top-right and bottom-left diagonals
	topRightRow, topRightCol, validTR := nextCheckpoint(row, col, "R", rows, cols)
	bottomLeftRow, bottomLeftCol, validBL := nextCheckpoint(row, col, "Z", rows, cols)
	if !validTR || !validBL {
		return false
	}
	if charArray[topRightRow][topRightCol] == 'M' {
		if charArray[bottomLeftRow][bottomLeftCol] != 'S' {
			return false
		}
	} else if charArray[topRightRow][topRightCol] == 'S' {
		if charArray[bottomLeftRow][bottomLeftCol] != 'M' {
			return false
		}
	} else {
		return false
	}

	return true
}

// Gets the next position in the specified direction
// Returns the new row, column, and whether the position is valid
func nextCheckpoint(row, col int, direction string, rows, cols int) (int, int, bool) {
	dir, exists := p2Directions[direction]
	if !exists {
		return -1, -1, false
	}
	newRow, newCol := row+dir[0], col+dir[1]
	if outOfBounds(newRow, newCol, rows, cols) {
		return -1, -1, false
	}
	return newRow, newCol, true
}

// Checks if the given coordinates are out of bounds
func outOfBounds(row, col, rows, cols int) bool {
	return row < 0 || col < 0 || row >= rows || col >= cols
}

// =============================================================================
//	Common utilities
// =============================================================================

// Reads the input file and returns a 2D slice of characters
func readInput(filename string) [][]rune {
	file, _ := os.Open(filename)
	defer file.Close()

	var charArray [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		charArray = append(charArray, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading lines: %v", err)
	}

	return charArray
}
