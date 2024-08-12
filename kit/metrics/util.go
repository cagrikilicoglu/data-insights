package metrics

import "strconv"

func GetTopElements[T any](slice []T, numberOfElements int) []T {
	return slice[:numberOfElements]
}

func GetBottomElements[T any](slice []T, numberOfElements int) []T {
	if numberOfElements > len(slice) {
		numberOfElements = len(slice)
	}
	bottomElements := slice[len(slice)-numberOfElements:]

	// Reverse the order of bottomElements
	for i, j := 0, len(bottomElements)-1; i < j; i, j = i+1, j-1 {
		bottomElements[i], bottomElements[j] = bottomElements[j], bottomElements[i]
	}

	return bottomElements
}

// Parse string to float64 with error handling
func parseStringToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return value
}
