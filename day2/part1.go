package main

import (
	"fmt"
	"strings"
	"strconv"
)

func solve1(lines []string) int{
	sum := 0
	for _, line := range lines {
		max_red := 0
		max_green := 0
		max_blue := 0

		split := strings.Split(line,":")
		id, err := strconv.Atoi(strings.Split(split[0]," ")[1])
		if err != nil {
			fmt.Println("Error converting id to integer")
		}
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
		if max_red <= 12 && max_green <= 13 && max_blue <= 14 {
			sum += id
		}
	}
	return sum
}
