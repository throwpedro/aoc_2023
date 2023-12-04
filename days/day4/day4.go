package days

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	fileContent, err := os.ReadFile("days/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fileContent), "\n")

	tally := 0
	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		tempWinNums := strings.Trim(strings.Split(line, "|")[0], " ")
		tempMyNums := strings.Trim(strings.Split(line, "|")[1], " ")
		winNumsStrArr := strings.Split(tempWinNums, " ")
		myNumsStrArr := strings.Split(tempMyNums, " ")
		winNums := []int{}
		for _, num := range winNumsStrArr {
			n, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			winNums = append(winNums, n)
		}
		myNums := []int{}
		for _, num := range myNumsStrArr {
			n, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			myNums = append(myNums, n)
		}
		hits := 0
		for _, num := range myNums {
			if slices.Contains(winNums, num) {
				if hits == 0 {
					hits = 1
				} else {
					hits = hits * 2
				}
			}
		}
		tally += hits
	}
	fmt.Println(tally)
}
