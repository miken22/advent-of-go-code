package twofour

import (
	"log"
	"strconv"
	"strings"

	"github.com/miken22/advent-of-go-code.git/pkg/utils"
)

var partOneOperands = []string{"*", "+"}
var partTwoOperands = []string{"||", "*", "+"}

func Solve_day_seven(input []string) {

	sum := 0
	// for _, line := range input {

	// 	split := strings.Split(line, ":")

	// 	result, err := strconv.Atoi(split[0])

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	numbers := utils.StringToIntSlice(split[1], " ")

	// 	sum += solve_recursive(numbers, result, false)
	// }

	// log.Printf("Part One Sum: %d", sum)

	sum = 0
	for _, line := range input {

		split := strings.Split(line, ":")

		result, err := strconv.Atoi(split[0])

		if err != nil {
			panic(err)
		}

		numbers := utils.StringToIntSlice(split[1], " ")

		total := solve_recursive(numbers, result, true)
		sum += total
	}

	log.Printf("Part Two Sum: %d", sum)

}

func solve_recursive(numbers []int, result int, partTwo bool) int {

	// set up recursive call
	if check(numbers, 1, result, numbers[0], partTwo) {
		return result
	}
	return 0
}

func check(numbers []int, index, result, currentTotal int, partTwo bool) bool {

	if currentTotal == result && index == len(numbers) {
		log.Printf("%d = %v\n", result, numbers)
		// we found it
		return true
	}

	if index >= len(numbers) {
		// too far, no result
		return false
	}

	holder := currentTotal
	operands := []string{}
	if partTwo {
		operands = partTwoOperands
	} else {
		operands = partOneOperands
	}

	for opIndex := 0; opIndex < len(operands); opIndex++ {
		switch operands[opIndex] {
		case "+":
			currentTotal += numbers[index]
		case "*":
			currentTotal *= numbers[index]
		case "||":
			currentTotal = concat(currentTotal, numbers[index])
		}

		// recursive call
		if check(numbers, index+1, result, currentTotal, partTwo) {
			return true
		}

		// undo operation.
		currentTotal = holder

	}

	// did not work with any operators
	return false

}

func concat(lhs, rhs int) int {
	lhsString := strconv.Itoa(lhs)
	rhsString := strconv.Itoa(rhs)
	concatString := lhsString + rhsString
	concat, _ := strconv.Atoi(concatString)
	return concat
}
