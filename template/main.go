package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Specify the file path
	filePath := "input.txt"

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Convert the byte slice to string and split it into lines
	lines := strings.Split(string(content), "\n")
	solve1(lines);
	// solve2(lines);
	
}
