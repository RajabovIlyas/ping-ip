package utils

import (
	"strconv"
)

func ConvertArrayStrToInt(arrayStr []string) ([]int, error) {
	intArray := make([]int, len(arrayStr))

	for i, str := range arrayStr {
		intVal, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intArray[i] = intVal
	}

	return intArray, nil
}
