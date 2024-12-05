package main

import (
	"os"

	"github.com/miken22/advent-of-go-code.git/internal/twofour"
	"github.com/miken22/advent-of-go-code.git/pkg/utils"
)

func main() {

	input := utils.LoadTestFileLines(os.Args[1])

	twofour.Solve_day_four(input)

}
