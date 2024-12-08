package utils

import "slices"

func RemoveAtIndex[T any](arr []T, index int) []T {
	if index < 0 || index >= len(arr) {
		return slices.Clone(arr)
	}

	result := make([]T, 0, len(arr)-1)
	result = append(result, arr[:index]...)
	result = append(result, arr[index+1:]...)

	return result
}

func DeleteAtIndex[T any](arr []T, index int) []T {
	if index < 0 || index >= len(arr) {
		return arr
	}

	return slices.Delete(arr, index, index+1)
}
