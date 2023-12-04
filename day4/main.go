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

	/// Convert the byte slice to string and split it into lines
	stripped_content := strings.TrimSpace(string (content))
	lines := strings.Split(stripped_content, "\n")
	fmt.Println("Part 1 - ", solve1(lines));
	fmt.Println("Part 2 - ", solve2(lines));
	
}
