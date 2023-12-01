package main

import (
	"fmt"
	"unicode"
	"strconv"
)

func solve2(lines []string) int {
	numberMap := map[string]string{
        "one":   "1",
        "two":   "2",
        "three": "3",
        "four":  "4",
        "five":  "5",
        "six":   "6",
        "seven": "7",
        "eight": "8",
        "nine":  "9",
    }

	sum := 0
	for _, line := range lines {
		val := ""
		// find 1st digit
		outer1:
			for i := 0; i < len(line); i++ {
				c := rune(line[i])
				if unicode.IsDigit(c) {
					val += string(c)
					break
				} else {
					// check if one of the keys of the map is the digit with a forward loop
					for key := range numberMap {
				        if i + len(key)-1 < len(line)-1 && line[i:i + len(key)] == key {
				        	val += numberMap[key]
				        	break outer1
				        }
				    }
				}
			}
		outer2:
			// find last digit
			for i := len(line)-1; i >= 0; i-- {
				c := rune(line[i])
				if unicode.IsDigit(c) {
					val += string(c)
					break
				} else {
					// check if one of the keys of the map is the digit with a reverse loop
					for key := range numberMap {
				        if i - len(key)+1 > 0 && line[i - len(key) + 1:i+1] == key {
				        	val += numberMap[key]
				        	break outer2
				        }
				    }
				}
			}
		ival, err := strconv.Atoi(val);
		if err != nil {
			fmt.Println("Cant convert to an integer")
		}
		sum += ival
	}
	return sum
}
