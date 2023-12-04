package main

import (
	"strings"
	"math"
)

func countOccurs(l []string, el string) int{
	total := 0
	for _, e := range l {
		if e == el {
			total += 1
		}
	}
	return total
}

func solve1(lines []string) int{
	// Loop through each line
	total := 0
	for _, line := range lines {
		row := strings.Split(line, "|")

		winners := strings.Split(strings.TrimSpace(strings.ReplaceAll(strings.Split(row[0], ":")[1],"  ", " ")), " ")
		nums := strings.Split(strings.TrimSpace(strings.ReplaceAll(row[1],"  "," ")), " ")
		m := 0
		for _ , w := range winners {
			m += countOccurs(nums, w)
		}
		if m != 0 {
			total += int(math.Pow(2,float64(m-1)))
		}
	}
	return total
}