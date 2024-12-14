package main

import (
	"fmt"
	"log"
	"path/filepath"
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

	tcReader, err := utils.NewTCReader(absFilePath, "")
	defer tcReader.Close()

	if err != nil {
		log.Fatalln("Error reading file path:\n", err)
		return
	}

	searchSpace := make([][]string, 0)
	XIndexList := make([][]int, 0)
	AIndexList := make([][]int, 0)

	rowIndex := 0
	n := 0
	for tcReader.Scan() {

		row := tcReader.Next()
		searchSpace = append(searchSpace, row)
		n = max(n, len(row))
		for index, ch := range row {
			if ch == "X" {
				position := []int{rowIndex, index}
				XIndexList = append(XIndexList, position)
			}
			if ch == "A" {
				position := []int{rowIndex, index}
				AIndexList = append(AIndexList, position)
			}
		}

		rowIndex++
	}

	m := rowIndex

	// fmt.Printf("Solution of Day 4 Part 1: %d\n", findXMAS(searchSpace, XIndexList, m, n))
	fmt.Printf("Solution of Day 4 Part 2: %d\n", findXShapedMAS(searchSpace, AIndexList, m, n))
}

func findXShapedMAS(searchSpace [][]string, AIndexList [][]int, m int, n int) int {
	answer := 0
	match1 := "MAS"
	match2 := "SAM"

	for _, position := range AIndexList {
		row, col := position[0], position[1]

		if row > 0 && row < m-1 && col > 0 && col < n-1 {
			str1 := strings.Join([]string{searchSpace[row-1][col-1], searchSpace[row][col], searchSpace[row+1][col+1]}, "")
			str2 := strings.Join([]string{searchSpace[row-1][col+1], searchSpace[row][col], searchSpace[row+1][col-1]}, "")

			if (str1 == match1 || str1 == match2) && (str2 == match1 || str2 == match2) {
				answer++
			}
		}
	}

	return answer
}

func findXMAS(searchSpace [][]string, XIndexList [][]int, m int, n int) int {
	match1 := "XMAS"
	match2 := "SAMX"
	answer := 0
	for _, position := range XIndexList {
		row, col := position[0], position[1]

		// right
		str1 := ""
		for j := col; j < col+4 && j < n; j++ {
			str1 = strings.Join([]string{str1, searchSpace[row][j]}, "")
		}

		if str1 == match1 || str1 == match2 {
			answer++
		}

		// left
		str2 := ""
		for j := col; j >= 0 && j > col-4; j-- {
			str2 = strings.Join([]string{str2, searchSpace[row][j]}, "")
		}

		if str2 == match1 || str2 == match2 {
			answer++
		}

		// top
		str3 := ""
		for i := row; i >= 0 && i > row-4; i-- {
			str3 = strings.Join([]string{str3, searchSpace[i][col]}, "")
		}

		if str3 == match1 || str3 == match2 {
			answer++
		}

		// bottom
		str4 := ""
		for i := row; i < m && i < row+4; i++ {
			str4 = strings.Join([]string{str4, searchSpace[i][col]}, "")
		}

		if str4 == match1 || str4 == match2 {
			answer++
		}

		// diag 1
		str5 := ""
		for i, j := row, col; i >= 0 && i > row-4 && j < n && j < col+4; {
			str5 = strings.Join([]string{str5, searchSpace[i][j]}, "")
			i--
			j++
		}

		if str5 == match1 || str5 == match2 {
			answer++
		}

		// diag 2
		str6 := ""
		for i, j := row, col; i >= 0 && i > row-4 && j >= 0 && j > col-4; {
			str6 = strings.Join([]string{str6, searchSpace[i][j]}, "")
			i--
			j--
		}

		if str6 == match1 || str6 == match2 {
			answer++
		}

		// diag 3
		str7 := ""
		for i, j := row, col; i < m && i < row+4 && j >= 0 && j > col-4; {
			str7 = strings.Join([]string{str7, searchSpace[i][j]}, "")
			i++
			j--
		}

		if str7 == match1 || str7 == match2 {
			answer++
		}

		// diag 4
		str8 := ""
		for i, j := row, col; i < m && i < row+4 && j < n && j < col+4; {
			str8 = strings.Join([]string{str8, searchSpace[i][j]}, "")
			i++
			j++
		}

		if str8 == match1 || str8 == match2 {
			answer++
		}

	}

	return answer
}
