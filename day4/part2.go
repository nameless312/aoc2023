package main

import (
	"strings"
)

func makeArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
	      arr[i] = 1
	}
	return arr
}

func solve2(lines []string) int{
	// Loop through each line
	winners := [][]string{}
	nums := [][]string{}
	counts := makeArr(len(lines))
	for i, line := range lines {
		row := strings.Split(line, "|")
		winners = append(winners,strings.Split(strings.TrimSpace(strings.ReplaceAll(strings.Split(row[0], ":")[1],"  ", " ")), " "))
		nums = append(nums,strings.Split(strings.TrimSpace(strings.ReplaceAll(row[1],"  "," ")), " "))
		occ := 0

		for j := 0; j < len(winners[i]); j++ {
			occ += countOccurs(nums[i],winners[i][j])
		}
		for repeatTimes := 0; repeatTimes < counts[i]; repeatTimes++ {
			for z := 0; z < occ; z++ {
				counts[i + z + 1] += 1
			}
		}
		
	}
	total := 0
	for _,c := range counts {
		total += c
	}

	return total
}