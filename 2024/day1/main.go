package main

import (
	"fmt"
	"log"
	"math"
	"path/filepath"
	"sort"
	"strings"

	utils "github.com/Aki0x137/aoc/utils"
)

func main() {
	filename := "input.txt"
	absFilePath, err := filepath.Abs(filename)

	if err != nil {
		log.Fatalln("Error reading file path", err)
		return
	}

	delimiter := strings.Repeat(" ", 3)
	tcReader, err := utils.NewTCReader(absFilePath, delimiter)
	defer tcReader.Close()

	if err != nil {
		log.Fatalln("Error reading file path", err)
		return
	}

	dest1 := make([]int, 0)
	dest2 := make([]int, 0)
	freqMap := make(map[int]int)

	for tcReader.Scan() {
		fields := tcReader.Next()
		input, err := utils.ConvertSlice(fields, utils.StringToInt)
		if err != nil {
			log.Fatalln("Error parsing input", err)
		}

		if len(input) >= 2 {
			dest1 = append(dest1, input[0])
			dest2 = append(dest2, input[1])

			key := input[1]
			if val, ok := freqMap[key]; ok {
				freqMap[key] = val + 1
			} else {
				freqMap[key] = 1
			}
		}
	}

	n := 0
	if len(dest1) > len(dest2) {
		n = len(dest2)
	} else {
		n = len(dest1)
	}

	fmt.Printf("Solution of Part1: %d\n", solvePart1(dest1, dest2, n))
	fmt.Printf("Solution of Part2: %d\n", solvePart2(dest1, freqMap))
}

func solvePart1(dest1 []int, dest2 []int, n int) int {
	sort.Ints(dest1)
	sort.Ints(dest2)

	result := 0

	for i := range n {
		x := dest1[i]
		y := dest2[i]
		diff := x - y
		result += int(math.Abs(float64(diff)))
	}

	return result
}

func solvePart2(dest1 []int, freqMap map[int]int) int {
	result := 0
	for _, destId := range dest1 {
		if val, ok := freqMap[destId]; ok {
			result += destId * val
		}
	}

	return result
}
