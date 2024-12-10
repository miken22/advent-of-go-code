package twofour

import "fmt"

func Solve_day_four(input []string) {

	// Part 1

	count := countXmas(input)

	fmt.Printf("Count: %d", count)

	// Part 2

	count = countSam(input)

	fmt.Printf("Count: %d", count)
}

func countSam(input []string) int {
	count := 0

	for rowIndex := 1; rowIndex < len(input)-1; rowIndex++ {
		for colIndex := 1; colIndex < len(input[0])-1; colIndex++ {

			if input[rowIndex][colIndex] == 'A' {
				count += searchSam(input, rowIndex, colIndex)
			}
		}
	}

	return count
}

func searchSam(input []string, rowIndex, colIndex int) int {
	count := 0

	if input[rowIndex-1][colIndex-1] == 'M' && input[rowIndex-1][colIndex+1] == 'S' &&
		input[rowIndex+1][colIndex-1] == 'M' && input[rowIndex+1][colIndex+1] == 'S' {
		count++
	}

	if input[rowIndex-1][colIndex-1] == 'M' && input[rowIndex-1][colIndex+1] == 'M' &&
		input[rowIndex+1][colIndex-1] == 'S' && input[rowIndex+1][colIndex+1] == 'S' {
		count++
	}

	if input[rowIndex-1][colIndex-1] == 'S' && input[rowIndex-1][colIndex+1] == 'S' &&
		input[rowIndex+1][colIndex-1] == 'M' && input[rowIndex+1][colIndex+1] == 'M' {
		count++
	}

	if input[rowIndex-1][colIndex-1] == 'S' && input[rowIndex-1][colIndex+1] == 'M' &&
		input[rowIndex+1][colIndex-1] == 'S' && input[rowIndex+1][colIndex+1] == 'M' {
		count++
	}

	return count
}

func countXmas(input []string) int {
	count := 0

	for rowIndex := 0; rowIndex < len(input); rowIndex++ {
		for colIndex := 0; colIndex < len(input[0]); colIndex++ {

			if input[rowIndex][colIndex] == 'X' {
				count += search(input, rowIndex, colIndex)
			}
		}
	}

	return count
}

func search(input []string, rowIndex, colIndex int) int {
	count := 0

	letters := []byte{'X', 'M', 'A', 'S'}

	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {

			// try to move in that direction
			depth := 1

			row := rowIndex + dy
			column := colIndex + dx

			for {
				if row < 0 || row >= len(input) || column < 0 || column >= len(input[0]) {
					break
				}
				if input[row][column] == letters[depth] {
					depth++
					row += dy
					column += dx
				} else {
					break
				}
				if depth == len(letters) {
					count++
					break
				}
			}
		}
	}

	return count
}
