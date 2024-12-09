package utils

func Swap[T any](input []T, originalIndex, newIndex int) []T {
	tmp := input[originalIndex]
	input[originalIndex] = input[newIndex]
	input[newIndex] = tmp
	return input
}

func Remove(slice []string, index int) []string {

	outSlice := make([]string, 0)

	for i := range slice {
		if i != index {
			outSlice = append(outSlice, slice[i])
		}
	}

	return outSlice
}
