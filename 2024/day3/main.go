package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/Aki0x137/aoc/utils"
)

func main() {
	filename := "input.txt"
	absFilePath, err := filepath.Abs(filename)

	if err != nil {
		log.Fatalln("Error reading file path:\n", err)
		return
	}

	tcReader, err := utils.NewTCReader(absFilePath, "")
	defer tcReader.Close()

	if err != nil {
		log.Fatalln("Error reading file path:\n", err)
		return
	}

	memory := string(tcReader.ReadEntireFile())

	mulPattern := `mul\(\d{1,3}\,\d{1,3}\)`
	enablePattern := `do\(\)`
	disablePattern := `don\'t\(\)`

	patterns := []string{mulPattern, enablePattern, disablePattern}

	fmt.Printf("Solution of Day 3 Part 1: %d\n", getAllMultiplicationSum(mulPattern, memory))
	fmt.Printf("Solution of Day 3 Part 2: %d\n", getFilteredMultiplicationSum(patterns, memory))
}

func getAllMultiplicationSum(pattern string, memory string) int {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(memory, -1)

	result := 0
	for _, match := range matches {
		x, y, err := extractNumbersFromMul(match)
		if err != nil {
			log.Fatalln("error while extracting nums\n", err)
			continue
		}
		result += x * y
	}

	return result
}

func getFilteredMultiplicationSum(patterns []string, memory string) int {
	result := 0

	reMul := regexp.MustCompile(patterns[0])
	reEnable := regexp.MustCompile(patterns[1])
	reDisable := regexp.MustCompile(patterns[2])

	mulIndices := reMul.FindAllStringIndex(memory, -1)
	enableIndices := reEnable.FindAllStringIndex(memory, -1)
	disableIndices := reDisable.FindAllStringIndex(memory, -1)

	for _, mulIndex := range mulIndices {
		start := mulIndex[0]
		enableIndex := binarySearch(enableIndices, start)
		disableIndex := binarySearch(disableIndices, start)

		if disableIndex == -1 || (enableIndex != -1 && enableIndex > disableIndex) {
			x, y, err := extractNumbersFromMul(memory[start:mulIndex[1]])
			if err != nil {
				log.Fatalln("error while extracting nums\n", err)
				continue
			}

			result += x * y
		}
	}

	return result
}

func extractNumbersFromMul(str string) (int, int, error) {
	str = strings.TrimPrefix(str, "mul(")
	str = strings.TrimSuffix(str, ")")

	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input string: %s", str)
	}

	x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func binarySearch(space [][]int, key int) int {
	start := 0
	end := len(space) - 1
	indx := -1

	for start <= end {
		mid := start + ((end - start) / 2)
		if space[mid][0] < key {
			indx = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if indx == -1 {
		return -1
	}

	return space[indx][0]
}
