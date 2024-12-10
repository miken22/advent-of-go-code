package main

import (
	"os"

	"github.com/miken22/advent-of-go-code.git/internal/twofour"
	"github.com/miken22/advent-of-go-code.git/pkg/utils"
)

func main() {

	practiceInput := utils.LoadTestFileLines(os.Args[1])

	// twofour.Solve_day_one(practiceInput)
	// twofour.Solve_day_two(practiceInput)
	// twofour.Solve_day_three(practiceInput[0])
	// twofour.Solve_day_four(practiceInput)
	// twofour.Solve_day_five(practiceInput)
	// twofour.Solve_day_six(practiceInput)
	// twofour.Solve_day_seven(practiceInput)
	// twofour.Solve_day_eight(practiceInput)
	// twofour.Solve_day_nine(practiceInput)
	twofour.Solve_day_ten(practiceInput)

	input := utils.LoadTestFileLines(os.Args[2])

	// twofour.Solve_day_one(input)
	// twofour.Solve_day_two(input)
	// twofour.Solve_day_three(input[0])
	// twofour.Solve_day_four(input)
	// twofour.Solve_day_five(input)
	// twofour.Solve_day_six(input)
	// twofour.Solve_day_seven(input)
	// twofour.Solve_day_eight(input)
	// twofour.Solve_day_nine(input)
	twofour.Solve_day_ten(input)

}
