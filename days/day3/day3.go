package days

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func Part1() {
	fileContent, err := os.ReadFile("days/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid := strings.Split(string(fileContent), "\n")

	indexes := make(map[int]int)

	for i, row := range grid {
		for j, char := range row {
			if unicode.IsDigit(char) || char == '.' {
				continue
			}
		}
	}

}
