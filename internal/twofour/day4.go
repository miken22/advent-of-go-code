package twofour

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/miken22/advent-of-go-code.git/pkg/utils"
)

func Solve_day_four(input []string) {

	// Part 1

	rules := make(map[int][]int)

	breakLine := 0
	for index, line := range input {
		if line == "" {
			breakLine = index
			break
		}
		ordering := strings.Split(line, "|")
		pageBefore, _ := strconv.Atoi(ordering[0])
		pageAfter, _ := strconv.Atoi(ordering[1])

		if _, ok := rules[pageBefore]; !ok {
			rules[pageBefore] = make([]int, 0)
		}

		rules[pageBefore] = append(rules[pageBefore], pageAfter)

	}

	validOrderings := make([][]int, 0)
	for x := breakLine + 1; x < len(input); x++ {
		pages := utils.StringToIntSlice(input[x], ",")

		if isValidOrdering(rules, pages) {
			validOrderings = append(validOrderings, pages)
		}

	}

	sum := 0
	for _, ordering := range validOrderings {
		midpoint := len(ordering) / 2
		sum += ordering[midpoint]
	}

	fmt.Printf("sum - %d\n", sum)

	// Part 2

	validOrderings = make([][]int, 0)
	for x := breakLine + 1; x < len(input); x++ {
		pages := utils.StringToIntSlice(input[x], ",")

		if !isValidOrdering(rules, pages) {
			correctedOrdering := correctOrdering(rules, pages)
			validOrderings = append(validOrderings, correctedOrdering)
		}

	}

	sum = 0
	for _, ordering := range validOrderings {
		midpoint := len(ordering) / 2
		sum += ordering[midpoint]
	}

	fmt.Printf("sum - %d\n", sum)

}

func correctOrdering(rules map[int][]int, pages []int) []int {
	correctSlice := pages
	for index := 0; index < len(pages); index++ {
		for nestedIndex := index + 1; nestedIndex < len(pages); nestedIndex++ {
			if ordering, ok := rules[pages[nestedIndex]]; ok {
				page := correctSlice[index]
				if slices.Contains(ordering, page) {
					correctSlice = utils.Swap(correctSlice, index, nestedIndex)
					index = -1
					break
				}
			}
		}
	}
	return correctSlice
}

func isValidOrdering(rules map[int][]int, pages []int) bool {

	for index, page := range pages {

		for nestedIndex := index + 1; nestedIndex < len(pages); nestedIndex++ {
			if ordering, ok := rules[pages[nestedIndex]]; ok {
				if slices.Contains(ordering, page) {
					return false
				}
			}
		}

	}

	return true
}
