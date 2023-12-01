package main

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

func main() {
	fileContent, err := os.ReadFile("1/input.txt")
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
			for j, num := range validStringNums {
				if strings.HasPrefix(line[i:], num) {
					digits = append(digits, strconv.Itoa(j+1))
				}
			}
		}
		concatenated := digits[0] + digits[len(digits)-1]
		added, _ := strconv.Atoi(concatenated)
		sum += added
	}
	fmt.Println(sum)
}
