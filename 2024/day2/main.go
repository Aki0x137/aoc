package main

import (
	"fmt"
	"log"
	"path/filepath"
	"slices"

	utils "github.com/Aki0x137/aoc/utils"
)

func main() {
	filename := "input.txt"
	filepath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalln("Error reading file path", err)
		return
	}

	delimiter := " "
	tcReader, err := utils.NewTCReader(filepath, delimiter)
	tcReader.Close()

	if err != nil {
		log.Fatalln("Error reading file path", err)
		return
	}

	fmt.Printf("Solution of Day 2 Part 1: %d\n", getSafeReportCount(tcReader, false))
	fmt.Printf("Solution of Day 2 Part 2: %d\n", getSafeReportCount(tcReader, true))
}

func getSafeReportCount(tcReader *utils.TCReader, isDampenerEnabled bool) int {
	result := 0

	for tcReader.Scan() {
		row := tcReader.Next()
		report, err := utils.ConvertSlice(row, utils.StringToInt)
		if err != nil {
			log.Fatalln("Error parsing input", err)
		}

		if len(report) < 2 || isSafeReport(report, isDampenerEnabled) {
			result++
		}
	}

	return result
}

func isSafeReport(report []int, isDampenerEnabled bool) bool {
	if report[0] > report[len(report)-1] {
		slices.Reverse(report)
	}

	if isDampenerEnabled {
		return isIncreasingSafelyWithDampener(report)
	}

	return isIncreasingSafely(report)
}

func isIncreasingSafely(report []int) bool {
	isSafe := true
	for indx, val := range report[1:] {
		prev := report[indx]
		if val <= prev || val-prev > 3 {
			isSafe = false
			break
		}
	}

	return isSafe
}

func isIncreasingSafelyWithDampener(report []int) bool {
	isSafe := true
	for indx, val := range report {
		if indx == 0 {
			continue
		}
		prev := report[indx-1]
		if val <= prev || val-prev > 3 {
			withoutPrev := utils.RemoveAtIndex(report, indx-1)
			withoutCurr := utils.RemoveAtIndex(report, indx)
			if isIncreasingSafely(withoutPrev) || isIncreasingSafely(withoutCurr) {
				isSafe = true
				break
			} else {
				isSafe = false
				break
			}
		}
	}

	return isSafe
}
