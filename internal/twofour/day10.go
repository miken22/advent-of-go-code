package twofour

import "log"

type trailLocation struct {
	row    int
	column int
}

var peak byte = '9'

func Solve_day_ten(input []string) {

	// Part 1
	score := scoreTrails(input)

	log.Printf("count: %d\n", score)

	// Part 2

	ratings := rateTrails(input)

	log.Printf("ratings: %d\n", ratings)
}

func rateTrails(input []string) int {
	trailHeads := findTrailHeads(input)
	rating := 0
	for _, trailHead := range trailHeads {
		step := byte('1')
		neighbours := getNeighbours(input, trailHead)
		rating += rateNeighbours(input, neighbours, step)
	}
	return rating
}

func rateNeighbours(input []string, neighbours []trailLocation, step byte) int {
	rating := 0
	for _, neighbour := range neighbours {
		if input[neighbour.row][neighbour.column] == step &&
			step == peak {
			rating++
			continue
		}
		newNeighbours := getNeighbours(input, neighbour)
		rating += rateNeighbours(input, newNeighbours, step+1)
	}
	return rating
}

func scoreTrails(input []string) int {
	trailHeads := findTrailHeads(input)
	score := 0
	for _, trailHead := range trailHeads {
		step := byte('1')
		uniqueEnds := map[trailLocation]bool{}
		neighbours := getNeighbours(input, trailHead)
		score += scoreNeighbours(input, neighbours, uniqueEnds, step)
	}
	return score
}

func scoreNeighbours(input []string, neighbours []trailLocation, uniqueEnds map[trailLocation]bool, step byte) int {
	count := 0

	for _, neighbour := range neighbours {
		if input[neighbour.row][neighbour.column] == step &&
			step == peak {
			if _, ok := uniqueEnds[neighbour]; !ok {
				count++
				uniqueEnds[neighbour] = true
			}
			continue
		}
		newNeighbours := getNeighbours(input, neighbour)
		count += scoreNeighbours(input, newNeighbours, uniqueEnds, step+1)
	}

	return count
}

func getNeighbours(input []string, node trailLocation) []trailLocation {
	deltas := []int{-1, 1}
	neighbours := []trailLocation{}
	for _, dx := range deltas {
		column := node.column + dx
		if column < 0 || column >= len(input[0]) {
			continue
		}
		if byte(input[node.row][column])-byte(input[node.row][node.column]) != 1 {
			continue
		}
		neighbours = append(neighbours, trailLocation{row: node.row, column: column})
	}
	for _, dy := range deltas {
		row := node.row + dy
		if row < 0 || row >= len(input) {
			continue
		}
		if byte(input[row][node.column])-byte(input[node.row][node.column]) != 1 {
			continue
		}
		neighbours = append(neighbours, trailLocation{row: row, column: node.column})
	}
	return neighbours
}

func findTrailHeads(input []string) []trailLocation {
	trailHeads := []trailLocation{}
	for rowIndex, row := range input {
		for colIndex, char := range row {
			if char == '0' {
				trailHeads = append(trailHeads, trailLocation{row: rowIndex, column: colIndex})
			}
		}
	}
	return trailHeads
}
