package metrics

import "strconv"

func GetTopElements[T any](slice []T, numberOfElements int) []T {
	return slice[:numberOfElements]
}

// todo terse çevir
func GetBottomElements[T any](slice []T, numberOfElements int) []T {
	return slice[len(slice)-numberOfElements:]
}

// Parse string to float64 with error handling
func parseStringToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return value
}
