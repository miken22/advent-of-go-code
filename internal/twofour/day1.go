package twofour

import (
	"fmt"
	"sort"
	"strings"

	"github.com/miken22/advent-of-go-code.git/pkg/utils"
)

func Solve_day_one(lines []string) {

	// Part One

	list_one := []int{}
	list_two := []int{}

	for _, line := range lines {
		list_vals := strings.Split(line, "   ")

		x := utils.ToInteger(list_vals[0])
		list_one = append(list_one, x)

		y := utils.ToInteger(list_vals[1])
		list_two = append(list_two, y)
	}

	sort.Ints(list_one)
	sort.Ints(list_two)

	diff := 0
	for x := range list_one {
		tmp := list_one[x] - list_two[x]

		if tmp < 0 {
			diff += list_two[x] - list_one[x]
		} else {
			diff += tmp
		}
	}

	fmt.Printf("difference: %d\n", diff)

	// Part Two
	occurances := map[int]int{}

	for x := range list_two {
		if _, ok := occurances[list_two[x]]; !ok {
			occurances[list_two[x]] = 1
		} else {
			occurances[list_two[x]] += 1
		}
	}

	score := 0
	for x := range list_one {
		if count, ok := occurances[list_one[x]]; ok {
			score += count * list_one[x]
		}
	}

	fmt.Printf("score: %d\n", score)
}
