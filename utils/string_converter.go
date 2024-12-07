package utils

import (
	"fmt"
	"strconv"
)

// ConvertSlice converts a slice of strings to a slice of a specified type.
func ConvertSlice[T any](input []string, convertFunc func(string) (T, error)) ([]T, error) {
	var result []T
	for _, s := range input {
		converted, err := convertFunc(s)
		if err != nil {
			return nil, fmt.Errorf("error converting value %q: %w", s, err)
		}
		result = append(result, converted)
	}
	return result, nil
}

// StringToInt is helper functions for string to integer conversions
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StringToFloat64 is helper functions for string to integer conversions
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// StringToBool is helper functions for string to integer conversions
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}
