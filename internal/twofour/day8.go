package twofour

import (
	"log"
)

type Location struct {
	row    int
	column int
}

var EMPTY_RUNE rune = '.'

var usedLocations = [][]bool{}

func Solve_day_eight(input []string) {

	// part 1

	usedLocations = make([][]bool, len(input))
	for key := range usedLocations {
		usedLocations[key] = make([]bool, len(input[0]))
	}

	mapping := findAntennas(input)

	totalAntinodes := findAntinodes(mapping)

	log.Printf("totalAntinodes: %d", totalAntinodes)

	// part 2

	usedLocations = make([][]bool, len(input))
	for key := range usedLocations {
		usedLocations[key] = make([]bool, len(input[0]))
	}

	totalAntinodes = findUpdatedAntinodes(mapping)

	log.Printf("totalAntinodes: %d", totalAntinodes)

}

func findUpdatedAntinodes(mapping map[string][]Location) int {
	sum := 0
	for _, locations := range mapping {
		for startIndex, startLocation := range locations {
			for compareIndex := startIndex + 1; compareIndex < len(locations); compareIndex++ {
				compareLocation := locations[compareIndex]
				dx, dy := computeDifferenceOfPoints(startLocation, compareLocation)
				antinodes := findLinearAntinodes(startLocation, compareLocation, dx, dy)
				// add node themselves
				antinodes = append(antinodes, []Location{startLocation, compareLocation}...)
				sum += countLinearAntinodes(antinodes)
			}
		}
	}
	return sum
}

func countLinearAntinodes(antinodes []Location) int {
	sum := 0
	for _, antinode := range antinodes {
		if validAntinode(antinode) {
			sum++
		}
	}
	return sum
}

func findLinearAntinodes(startLocation, compareLocation Location, dx, dy int) []Location {

	antinodes := []Location{}

	xPos := startLocation.column + dx
	yPos := startLocation.row + dy

	for {
		if xPos < 0 || xPos >= len(usedLocations) || yPos < 0 || yPos >= len(usedLocations[0]) {
			break
		}
		antinodes = append(antinodes, Location{
			row:    yPos,
			column: xPos,
		})
		xPos = xPos + dx
		yPos = yPos + dy
	}

	xPos = compareLocation.column - dx
	yPos = compareLocation.row - dy
	for {
		if xPos < 0 || xPos >= len(usedLocations) || yPos < 0 || yPos >= len(usedLocations[0]) {
			break
		}
		antinodes = append(antinodes, Location{
			row:    yPos,
			column: xPos,
		})
		xPos = xPos - dx
		yPos = yPos - dy
	}

	return antinodes
}

func findAntinodes(mapping map[string][]Location) int {
	sum := 0
	for _, locations := range mapping {
		for startIndex, startLocation := range locations {
			for compareIndex := startIndex + 1; compareIndex < len(locations); compareIndex++ {
				compareLocation := locations[compareIndex]
				dx, dy := computeDifferenceOfPoints(startLocation, compareLocation)
				// try to apply in both directions
				antinodeOne := Location{
					row:    startLocation.row + dy,
					column: startLocation.column + dx,
				}
				// try to apply in both directions
				antinodeTwo := Location{
					row:    compareLocation.row - dy,
					column: compareLocation.column - dx,
				}

				sum += countValidAntinodes(antinodeOne, antinodeTwo)
			}
		}
	}

	return sum
}

func countValidAntinodes(antinodeOne, antinodeTwo Location) int {
	count := 0
	if validAntinode(antinodeOne) {
		count++
	}
	if validAntinode(antinodeTwo) {
		count++
	}
	return count
}

func validAntinode(antinode Location) bool {
	if antinode.row >= 0 && antinode.row < len(usedLocations) &&
		antinode.column >= 0 && antinode.column < len(usedLocations[0]) &&
		!usedLocations[antinode.row][antinode.column] {
		usedLocations[antinode.row][antinode.column] = true
		return true
	}
	return false
}

func computeDifferenceOfPoints(startLocation, compareLocation Location) (int, int) {
	dx := startLocation.column - compareLocation.column
	dy := startLocation.row - compareLocation.row
	return dx, dy
}

func findAntennas(input []string) map[string][]Location {

	mapping := map[string][]Location{}

	for row, rowStr := range input {
		for column, columnRune := range rowStr {
			if columnRune == EMPTY_RUNE {
				continue
			}
			columnStr := string(columnRune)
			if _, ok := mapping[columnStr]; !ok {
				mapping[columnStr] = []Location{{row: row, column: column}}
			} else {
				mapping[columnStr] = append(mapping[columnStr], Location{row: row, column: column})
			}
		}
	}

	return mapping
}
