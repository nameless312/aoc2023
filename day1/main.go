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

	stripped_content := strings.TrimSpace(string (content))
	lines := strings.Split(stripped_content, "\n")
	// fmt.Println(solve1(lines));
	fmt.Println(solve2(lines));
	
}
