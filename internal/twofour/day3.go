package twofour

import (
	"fmt"
	"regexp"

	helpers "github.com/miken22/advent-of-go-code.git/pkg/utils"
)

func Solve_day_three(memory string) {

	// Part 1

	total := 0
	total += findTotal(memory)

	fmt.Printf("total - %d\n", total)

	// Part 2

	total = 0
	total += findConditionalTotal(memory)

	fmt.Printf("total - %d\n", total)

}

func findConditionalTotal(memory string) int {
	total := 0
	tokens := "mul\\(\\d*,\\d*\\)|don't\\(\\)|do\\(\\)"
	matches := findMatches(memory, tokens)

	memEnabled := true

	for _, match := range matches {

		if match == DONT_FUNC() {
			memEnabled = false
			// log.Printf("disabling at index %d", index)
			// continue
		} else if match == DO_FUNC() {
			memEnabled = true
			// log.Printf("enabling at index %d", index)
		} else if memEnabled {
			total += findTotal(match)
		}

	}

	return total
}

func DO_FUNC() string {
	return "do()"
}

func DONT_FUNC() string {
	return "don't()"
}

func findTotal(line string) int {
	tokens := "mul\\(\\d*,\\d*\\)"
	matches := findMatches(line, tokens)
	total := 0

	for _, match := range matches {
		digits := findMatches(match, "\\d*")
		x1 := helpers.ToInteger(digits[0])
		x2 := helpers.ToInteger(digits[1])
		total += x1 * x2
	}
	return total
}

func findMatches(line, pattern string) []string {

	r, _ := regexp.Compile(pattern)

	return r.FindAllString(line, -1)

}
