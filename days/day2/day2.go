package days

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	RedMax   = 12
	GreenMax = 13
	BlueMax  = 14
)

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func Day2() {
	fileContent, err := os.ReadFile("days/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fileContent), "\n")
	sum := 0
	for i, line := range lines {
		temp := strings.Split(line, ": ")
		line = temp[1]
		set := SplitAny(line, ",;")
		valid := true
		for _, numColorPair := range set {
			trimmedPair := strings.Trim(numColorPair, " ")
			num, _ := strconv.Atoi(strings.Split(trimmedPair, " ")[0])
			if (strings.Contains(trimmedPair, "red") && num > RedMax) || (strings.Contains(trimmedPair, "green") && num > GreenMax) || (strings.Contains(trimmedPair, "blue") && num > BlueMax) {
				valid = false
				break
			}
		}
		if valid {
			sum += i + 1
		}
	}
	fmt.Println(sum)
}
