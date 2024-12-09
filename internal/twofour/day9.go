package twofour

import (
	"log"

	"github.com/miken22/advent-of-go-code.git/pkg/utils"
)

var emptyInt int = -1

func Solve_day_nine(input []string) {

	diskMap := input[0]

	diskMapArr := makeStringMap(diskMap)

	// part 1

	rearrange(diskMapArr)

	log.Print(checksum(diskMapArr))

	// part 2

	diskMapArr = makeStringMap(diskMap)

	defrag(diskMapArr)

	log.Printf("defragged: %v", diskMapArr)

	log.Print(checksum(diskMapArr))

}

var movedFiles = map[int]bool{}

func defrag(diskMapArr []int) {

	rightPosition := len(diskMapArr) - 1

	for {
		if rightPosition < 0 {
			break
		}

		fileId := diskMapArr[rightPosition]

		if fileId == emptyInt {
			rightPosition--
			continue
		}

		if _, ok := movedFiles[fileId]; ok {
			rightPosition--
			continue
		}

		movedFiles[fileId] = true

		fileSize := 0
		for {
			if rightPosition < 0 {
				fileSize = -1
				break
			}
			if diskMapArr[rightPosition] != fileId {
				break
			}
			fileSize++
			rightPosition--
		}

		candidates := findFreeBlocks(fileSize, rightPosition)
		if len(candidates) == 0 {
			continue
		}

		leftPosition := len(diskMapArr)
		blockSize := 0
		for _, candidate := range candidates {
			if candidate.location < leftPosition {
				leftPosition = candidate.location
				blockSize = candidate.size
			}
		}
		for j := leftPosition; j < leftPosition+fileSize; j++ {
			diskMapArr[j] = fileId
		}
		for j := rightPosition + 1; j < rightPosition+fileSize+1; j++ {
			diskMapArr[j] = -1
		}

		delete(freeSpace, leftPosition)

		if fileSize < blockSize {
			freeSpace[leftPosition+fileSize] = blockSize - fileSize
		}

	}

}

type pair struct {
	location int
	size     int
}

func findFreeBlocks(fileSize, rightPosition int) []pair {
	if fileSize == -1 {
		return []pair{}
	}
	candidates := []pair{}
	for key, value := range freeSpace {
		if key <= rightPosition && value >= fileSize {
			candidates = append(candidates, pair{location: key, size: value})
		}
	}
	return candidates
}

func checksum(diskMapArr []int) int {
	sum := 0
	for index, char := range diskMapArr {
		if char == emptyInt {
			continue
		}
		sum += index * diskMapArr[index]
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

var freeSpace = map[int]int{}

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
			freeSpace[len(diskMapArr)] = val
			for range val {
				diskMapArr = append(diskMapArr, -1)
			}
		}
	}
	return diskMapArr
}
