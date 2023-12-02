package main

import (
	"fmt"
	"strings"
	"strconv"
)

func solve2(lines []string) int{
	total := 0
	for _, line := range lines {
		max_red := 0
		max_green := 0
		max_blue := 0

		split := strings.Split(line,":")
		sets := strings.Split(split[1],";")

		for _, subset := range sets {
			set:= strings.Split(subset, ",")
			for _, vals := range set {

				vals := strings.Split(strings.TrimSpace(vals), " ")
				num, err := strconv.Atoi(vals[0])
				if err != nil {
					fmt.Println("Error converting to integer")
				}
				color := vals[1]

				if color == "red" && max_red < num {
					max_red = num
				}
				if color == "green" && max_green < num {
					max_green = num
				}
				if color == "blue" && max_blue < num {
					max_blue = num
				}
			}

		}
		total += (max_red*max_green*max_blue)
	}
	return total
}
