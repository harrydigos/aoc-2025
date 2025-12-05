package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day02/input.txt")
	ranges := strings.Split(string(input), ",")

	fmt.Println("Part 1:", solvePart1(ranges)) // 55916882972
	fmt.Println("Part 2:", solvePart2(ranges)) // 76169125915
}

func solvePart1(items []string) int {
	total := 0

	for _, item := range items {
		slice := strings.Split(item, "-")
		start, err1 := strconv.Atoi(strings.TrimSpace(slice[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(slice[1]))

		if err1 != nil || err2 != nil {
			continue
		}

		for i := start; i <= end; i++ {
			numStr := strconv.Itoa(i)
			if len(numStr)%2 != 0 {
				continue
			}
			mid := len(numStr) / 2
			firstHalf := numStr[:mid]
			secondHalf := numStr[mid:]

			if firstHalf == secondHalf {
				total += i
			}
		}
	}

	return total
}

func isRepeatingSequence(n int) bool {
	s := strconv.Itoa(n)

	for length := 1; length <= len(s)/2; length++ {
		if len(s)%length != 0 {
			continue
		}

		if strings.Repeat(s[:length], len(s)/length) == s {
			return true
		}
	}
	return false
}

func solvePart2(items []string) int {
	total := 0

	for _, item := range items {
		slice := strings.Split(item, "-")
		start, err1 := strconv.Atoi(strings.TrimSpace(slice[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(slice[1]))

		if err1 != nil || err2 != nil {
			continue
		}

		for i := start; i <= end; i++ {
			if isRepeatingSequence(i) {
				total += i
			}
		}

	}

	return total
}
