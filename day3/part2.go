package main

import (
	// "fmt"
	"unicode"
	"strconv"
)

func isDigit(char string) bool {
	if (unicode.IsDigit(rune(char[0]))) {
		return true
	}
	return false
}

func isStar(char string) bool {
	return char == "*"
}

func findNumber(idx int, row int, lines []string) (int, int) {
	i := idx - 1
	endIdx := i
	num := string(lines[row][idx])
	
	for i >= 0 {
		if isDigit(string(lines[row][i])) {
			num = string(lines[row][i]) + num
			i--
		} else {
			break
		}
	}
	i = idx + 1
	endIdx = i
	for i < len(lines[row]) {
		if isDigit(string(lines[row][i])) {
			num = num + string(lines[row][i])
			i++
			endIdx++
		} else {
			break
		}
	}
	res, err := strconv.Atoi(num)
	if err != nil {
		return 0, 0
	}

	return res, endIdx
}

func findNumbers(idx int, row int, lines []string) []int {
	
	nums := []int{}
	startIdxSearch := idx
	endIdxSearch := idx

	if idx > 0 {
		startIdxSearch -= 1
	}

	if idx < len(lines[row]) - 1 {
		endIdxSearch += 1
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
		// search row up for a number including the diagonals if possible
		for i := startIdxSearch; i <= endIdxSearch; i++ {
			if isDigit(string(lines[row-1][i])) {
				num, endIdx := findNumber(i,row - 1, lines)
				if (num != 0) {
					nums = append(nums, num)
					if endIdx > i {
						i = endIdx
					}
				}
				
			}
		}
	}
	
	if searchDown {
		// search row down for a number including the diagonals if possible
		for i := startIdxSearch; i <= endIdxSearch; i++ {
			if isDigit(string(lines[row+1][i])) {
				num, endIdx := findNumber(i,row + 1, lines)
				if (num != 0) {
					nums = append(nums, num)
					if endIdx > i {
						i = endIdx
					}
				}
				
			}
		}
	}

	num, _ := findNumber(idx-1,row, lines)
	if (num != 0) {
		nums = append(nums, num)
	}
	num, _ = findNumber(idx+1,row, lines)
	if (num != 0) {
		nums = append(nums, num)
	}
	
	return nums
}

func solve2(lines []string) int{
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			char := string(lines[i][j])
			if isStar(char) {
				l := findNumbers(j,i,lines)
				if len(l) == 2 {
					multiplication := 1
					for _, v := range l {
						multiplication*=v 
					}
					total += multiplication
				}				
			}
		}
	}
	return total
}