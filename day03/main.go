package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day03/input.txt")
	lines := strings.Split(string(input), "\n")

	fmt.Println("Part 1:", solve(lines, 2))  // 17343
	fmt.Println("Part 2:", solve(lines, 12)) // 172664333119298
}

func solve(lines []string, keep int) int {
	total := 0

	for _, line := range lines {
		digits := []int{}

		for _, char := range line {
			digits = append(digits, int(char-'0'))
		}

		if len(digits) < keep {
			continue
		}

		total += countJoltage(selectMaxDigits(digits, keep))
	}

	return total
}

func selectMaxDigits(digits []int, keep int) []int {
	result := []int{}
	position := 0

	for len(result) < keep {
		remaining := keep - len(result)
		windowEnd := len(digits) - remaining + 1

		maxDigit := digits[position]
		maxIdx := position
		for i := position; i < windowEnd; i++ {
			if digits[i] > maxDigit {
				maxDigit = digits[i]
				maxIdx = i
			}
		}

		result = append(result, maxDigit)
		position = maxIdx + 1
	}

	return result
}

func countJoltage(digits []int) int {
	joltage := 0
	for _, digit := range digits {
		joltage = joltage*10 + digit
	}

	return joltage
}
