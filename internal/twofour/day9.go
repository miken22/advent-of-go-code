package twofour

import (
	"encoding/json"
	"log"
	"os"
	"slices"

	"github.com/miken22/advent-of-go-code.git/pkg/utils"
)

var emptyByte byte = '.'
var emptyInt int = -1

func Solve_day_nine(input []string) {

	diskMap := input[0]

	diskMapArr := makeStringMap(diskMap)

	// part 1

	rearrange(diskMapArr)

	log.Print(checksum(diskMapArr))

	// part 2

	diskMapArr = makeStringMap(diskMap)

	rearrangeWholeFiles(diskMapArr)

	log.Printf("rearranged: %v", diskMapArr)

	file, _ := os.ReadFile("../../output")

	type comparison struct {
		theirs []string
	}
	comp := comparison{}
	json.Unmarshal(file, &comp)
	compMap := make([]int, len(diskMapArr))
	for index, val := range comp.theirs {
		if val == "." {
			compMap[index] = -1
		} else {
			compMap[index] = utils.ToInteger(val)
		}
	}

	if slices.Equal(diskMapArr, compMap) {
		print("ahhh")
	}

	log.Print(checksum(diskMapArr))

}

func rearrangeWholeFiles(diskMapArr []int) {

	leftPosition, rightPosition := findLeftRightPosition(diskMapArr)

	for {
		if rightPosition <= 0 {
			break
		}
		if diskMapArr[rightPosition] == emptyInt {
			rightPosition--
			continue
		}

		fileSize := 0
		leftPosition, fileSize = findSpaceAndFileSize(diskMapArr, rightPosition)

		if leftPosition != -1 && leftPosition < rightPosition {
			swap2(diskMapArr, leftPosition, rightPosition, fileSize)
			leftPosition += fileSize
			rightPosition -= fileSize
			// log.Printf("updated: %v", diskMapArr)
		} else {
			rightPosition -= fileSize
			leftPosition = 0
		}

	}

}

func swap2(diskMapArr []int, leftPosition, rightPosition, size int) {
	for {
		if size == 0 {
			break
		}
		diskMapArr[leftPosition] = diskMapArr[rightPosition]
		diskMapArr[rightPosition] = emptyInt
		leftPosition++
		rightPosition--
		size--
	}
}

func findSpaceAndFileSize(diskMapArr []int, rightPosition int) (int, int) {
	fileSize := 1
	fileId := diskMapArr[rightPosition]
	rightPosition -= 1

	for {
		if rightPosition < 0 {
			break
		}
		if diskMapArr[rightPosition] != fileId {
			break
		}
		rightPosition--
		fileSize++
	}

	emptyChunk := 0
	for leftPosition := 0; leftPosition <= len(diskMapArr)/2; leftPosition++ {
		if diskMapArr[leftPosition] != emptyInt {
			emptyChunk = 0
		}
		if emptyChunk >= fileSize {
			return leftPosition - fileSize + 1, fileSize
		}
		emptyChunk++
	}

	return -1, fileSize
}

func checksum(diskMapArr []int) int {
	sum := 0
	fileId := 0
	for index, char := range diskMapArr {
		print(index, " ", diskMapArr[index], "\n")
		if char == emptyInt {
			continue
		}
		sum += index * diskMapArr[index]
		fileId++
	}
	return sum
}

func rearrange(diskMapArr []int) []int {

	leftPosition, rightPosition := findLeftRightPosition(diskMapArr)

	for {
		if leftPosition >= rightPosition {
			break
		}
		if diskMapArr[leftPosition] != emptyInt {
			leftPosition++
			continue
		}
		if diskMapArr[rightPosition] == emptyInt {
			rightPosition--
			continue
		}
		diskMapArr = utils.Swap(diskMapArr, leftPosition, rightPosition)
		leftPosition++
		rightPosition--
	}

	return diskMapArr

}

func findLeftRightPosition(diskMapArr []int) (int, int) {

	leftPosition := 0
	rightPosition := len(diskMapArr) - 1

	for {
		if diskMapArr[rightPosition] == emptyInt {
			rightPosition--
		} else {
			break
		}
	}

	for {
		if diskMapArr[leftPosition] != emptyInt {
			leftPosition++
		} else {
			break
		}
	}
	return leftPosition, rightPosition
}

func makeStringMap(diskMap string) []int {
	diskMapArr := make([]int, 0)
	fileId := 0
	for index, char := range diskMap {
		val := utils.ToInteger(string(char))
		if index%2 == 0 {
			for range val {
				diskMapArr = append(diskMapArr, fileId)
			}
			fileId++
		} else {
			for range val {
				diskMapArr = append(diskMapArr, -1)
			}
		}
	}
	return diskMapArr
}
