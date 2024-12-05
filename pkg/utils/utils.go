package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// LoadTestFileLines opens the specified file path and parses each
// line into a slice of strings.
func LoadTestFileLines(filePath string) []string {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// LLoadTestFile loads the entire contents of the given file into
// a byte slice.
func LoadTestFile(filePath string) []byte {
	file, _ := os.ReadFile(filePath)
	return file
}

// ToInteger is a helper function to convert strings to integers.
func ToInteger(s string) int {
	integer, _ := strconv.Atoi(s)
	return integer
}

func StringToIntSlice(input, delim string) []int {
	slice := strings.Split(input, delim)
	outSlice := make([]int, 0)
	for _, element := range slice {
		val, err := strconv.Atoi(element)
		if err != nil {
			panic(err)
		}
		outSlice = append(outSlice, val)
	}
	return outSlice
}
