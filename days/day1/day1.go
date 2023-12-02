package days

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var validStringNums = [9]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// DO_AOC_PART_TWO is a constant that determines whether to do part two of the
// advent of code challenge. https://adventofcode.com/2023/day/1
const DO_AOC_PART_TWO = true

func Day1() {
	fileContent, err := os.ReadFile("days/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fileContent), "\n")
	sum := 0
	for _, line := range lines {
		digits := []string{}
		for i, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
			}
			if DO_AOC_PART_TWO {
				for j, num := range validStringNums {
					if strings.HasPrefix(line[i:], num) {
						digits = append(digits, strconv.Itoa(j+1))
					}
				}
			}
		}
		concatenated := digits[0] + digits[len(digits)-1]
		added, _ := strconv.Atoi(concatenated)
		sum += added
	}
	fmt.Println(sum)
}
