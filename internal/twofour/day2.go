package twofour

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/miken22/advent-of-go-code.git/pkg/utils"
)

func Solve_day_two(input []string) {

	// part 1
	valid := 0

	for _, line := range input {

		report := strings.Split(line, " ")

		if validReport(report) {
			valid++
		}
	}

	fmt.Printf("Part 1 valid reports: %d\n", valid)

	// part 2

	valid = 0

	for _, line := range input {

		report := strings.Split(line, " ")

		if !validReport(report) {
			for index := range report {
				dampenedReport := utils.Remove(report, index)
				if validReport(dampenedReport) {
					valid++
					break
				}
			}

		} else {
			valid++
		}
	}

	fmt.Printf("Part 2 valid reports: %d\n", valid)
}

func validReport(report []string) bool {

	increasing := false
	for index := range report {

		if index == 0 {

			x1, _ := strconv.Atoi(report[index])
			x2, _ := strconv.Atoi(report[index+1])

			diff := x1 - x2

			if diff == 0 {
				return false
			}

			if diff < 0 {
				increasing = true
			} else {
				increasing = false
			}

			continue
		}

		x1, _ := strconv.Atoi(report[index-1])
		x2, _ := strconv.Atoi(report[index])
		diff := x1 - x2

		if increasing {

			if diff < -3 || diff >= 0 {
				return false
			}

		} else {

			if diff > 3 || diff <= 0 {
				return false
			}

		}

	}

	return true
}
