package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func RunDayOne() {
	readFile()
}
func readFile() {
	file, err := os.Open("input.txt")
	if err != nil {
		exit("Failed to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elvesCaloryCount := make([]int, 1)
	currentElf := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elvesCaloryCount = append(elvesCaloryCount, 0)
			currentElf++
		} else {
			caloryCount, err := strconv.Atoi(line)
			if err != nil {
				exit(err.Error())
			}
			elvesCaloryCount[currentElf] += caloryCount
		}
	}
	sort.Ints(elvesCaloryCount)
	length := len(elvesCaloryCount)
	max := elvesCaloryCount[len(elvesCaloryCount)-1]
	topThree := elvesCaloryCount[length-3 : length]
	fmt.Println(topThree)
	fmt.Println(max)
	fmt.Println(topThree[0] + topThree[1] + topThree[2])

	if err := scanner.Err(); err != nil {
		exit(err.Error())
	}

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
