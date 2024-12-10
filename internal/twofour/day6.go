package twofour

import (
	"fmt"
	"slices"
	"strings"
)

var (
	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

func Solve_day_six(input []string) {

	row, column, direction := get_start_position(input)

	originalInput := make([]string, len(input))
	copy(originalInput, input)

	// account for starting position
	findRoute(input, row, column, direction)

	steps := 0

	for row := range input {
		steps += strings.Count(input[row], "X")
	}

	fmt.Printf("number of steps: %d\n", steps)

	copy(input, originalInput)

	numLoops := findLoops(input, row, column)

	fmt.Printf("number of loops: %d\n", numLoops)

}

var loopFound bool = false
var visitedNodes = map[location][]int{}

// Sloooooooooooooooooooooooooooooooooooooooooooooooow
func findLoops(input []string, row, column int) int {

	startPos := row
	startCol := column
	numLoops := 0

	for obsRow, rowStr := range input {

		for obsCol := range rowStr {

			row = startPos
			column = startCol
			loopFound = false
			left = false
			direction := 0

			for {

				steps := 0

				steps += walkDirection2(input, row, column, obsRow, obsCol, direction)

				if left {
					clear(visitedNodes)
					break
				}

				if loopFound {
					numLoops++
					clear(visitedNodes)
					break
				}

				switch direction {
				case UP:
					row -= steps
				case DOWN:
					row += steps
				case LEFT:
					column -= steps
				case RIGHT:
					column += steps
				}

				direction = rotate(direction)
			}

		}

	}

	return numLoops

}

type location struct {
	row       int
	column    int
	direction int
}

func walkDirection2(input []string, row, column, obsRow, obsCol, direction int) int {
	steps := 0

	dx := 0
	dy := 0
	switch direction {
	case UP:
		dy = -1
	case DOWN:
		dy = 1
	case LEFT:
		dx = -1
	case RIGHT:
		dx = 1
	}

	for {

		if !(row+dy >= 0 && row+dy < len(input) && column+dx >= 0 && column+dx < len(input[0])) {
			left = true
			return steps
		}

		row += dy
		column += dx
		char := input[row][column]

		if char == '#' || (row == obsRow && column == obsCol) {
			// hit obstacle, need to rotate
			return steps
		}

		loc := location{row: row, column: column, direction: direction}

		if _, ok := visitedNodes[loc]; !ok {
			visitedNodes[loc] = []int{direction}
		} else if ok {
			if slices.Contains(visitedNodes[loc], direction) {
				loopFound = true
				return steps
			}
		}

		steps++
	}

}

var left bool = false

func findRoute(input []string, row, column, direction int) {

	for {

		steps := 0

		steps += walkDirection(input, row, column, direction)

		if left {
			break
		}

		switch direction {
		case UP:
			row -= steps
		case DOWN:
			row += steps
		case LEFT:
			column -= steps
		case RIGHT:
			column += steps
		}

		direction = rotate(direction)

	}
}

func rotate(direction int) int {
	return (direction + 1) % 4
}

func walkDirection(input []string, row, column, direction int) int {
	steps := 0

	dx := 0
	dy := 0
	switch direction {
	case UP:
		dy = -1
	case DOWN:
		dy = 1
	case LEFT:
		dx = -1
	case RIGHT:
		dx = 1
	}

	for {
		if !(row+dy >= 0 && row+dy < len(input) && column+dx >= 0 && column+dx < len(input[0])) {
			left = true
			return steps
		}

		row += dy
		column += dx
		char := input[row][column]
		if char == '#' {
			return steps
		}

		t := []rune(input[row])
		t[column] = 'X'
		input[row] = string(t)

		steps++
	}

}

func get_start_position(input []string) (int, int, int) {

	for index, row := range input {
		if col := strings.Index(row, "^"); col != -1 {
			return index, col, UP
		} else if col = strings.Index(row, "v"); col != -1 {
			return index, col, DOWN
		} else if col = strings.Index(row, ">"); col != -1 {
			return index, col, RIGHT
		} else if col = strings.Index(row, "<"); col != -1 {
			return index, col, LEFT
		}
	}

	return 0, 0, 0
}
