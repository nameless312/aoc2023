package main

import (
	"fmt"
	"unicode"
	"strconv"
)

func isSymbol(char string) bool {
	return char != "." && !unicode.IsDigit(rune(char[0]))
}

func hasSymbolAdj(startIdx int, endIdx int, row int, lines []string) bool {
	startIdxSearch := startIdx
	endIdxSearch := endIdx
	if startIdx > 0 {
		// set the start index to include the left side and the left diagonal down
		startIdxSearch = startIdx - 1
		if isSymbol(string(lines[row][startIdxSearch])) {
			return true
		}
	}
	if endIdx < len(lines[row]) - 1 {
		// set the start index to include the right side and the right diagonal down
		endIdxSearch = endIdx + 1
		if isSymbol(string(lines[row][endIdxSearch])) {
			return true
		}
	}
	searchUp := true
	searchDown := true
	
	if row == 0 {
		searchUp = false
	}
	if row == len(lines) - 1{
		searchDown = false
	}

	if searchUp {
		// search row up for a symbol including the diagonals if possible
		for i := startIdxSearch; i <= endIdxSearch; i++ {
			if isSymbol(string(lines[row-1][i])) {
				return true
			}
		}
	}
	
	if searchDown {
		// search row down for a symbol including the diagonals if possible
		for i := startIdxSearch; i <= endIdxSearch; i++ {
			if isSymbol(string(lines[row+1][i])) {
				return true
			}
		}
	}
	return false
}

func solve1(lines []string) int{
	total := 0
	for i := 0; i < len(lines); i++ {
		num := 0
		for j := 0; j < len(lines[i]); j++ {
			char := string(lines[i][j])
			// println("char ",char)
			// if a digit was found add to num
			if (unicode.IsDigit(rune(char[0]))) {
				digit, err := strconv.Atoi(char)
				if err != nil {
					fmt.Println("Could not parse digit")
				}
				num = (num * 10) + digit
				// end of the line, check the value before its reset
				if j + 1 == len(lines[i]) {
					if hasSymbolAdj(j - len(string(num)),j - 1,i,lines) {
						total += num
					}
				}
			} else if num != 0 {
				if hasSymbolAdj(j - len(strconv.Itoa(num)),j - 1,i,lines) {
					total += num
				}
				num = 0
			}
		}
	}
	return total
}