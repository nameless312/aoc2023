package main

import (
	"fmt"
	"unicode"
	"strconv"
)

func solve1(lines []string) int {
	sum := 0
	for _, line := range lines {
		val := ""
		// find 1st digit
		for i := 0; i < len(line); i++ {
			c := rune(line[i])
			if unicode.IsDigit(c) {
				val += string(c)
				break
			}
		}
		// find last digit
		for i := len(line)-1; i >= 0; i-- {
			c := rune(line[i])
			if unicode.IsDigit(c) {
				val += string(c)
				break
			}
		}
		ival, err := strconv.Atoi(val);
		if err != nil {
			fmt.Println("Cant convert to an integer")
		}
		sum += ival
		fmt.Println(val)
	}
	return sum
}
