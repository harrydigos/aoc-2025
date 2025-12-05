package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day03/input.txt")
	lines := strings.Split(string(input), "\n")

	fmt.Println("Part 1:", solvePart1(lines)) // 17343
}

func solvePart1(lines []string) int {
	total := 0

	for _, line := range lines {
		digits := []int{}

		for _, char := range line {
			digits = append(digits, int(char-'0'))
		}

		if len(digits) < 2 {
			continue
		}

		max1 := digits[0]
		maxIdx := 0
		for i := 1; i < len(digits)-1; i++ {
			if digits[i] > max1 {
				max1 = digits[i]
				maxIdx = i
			}
		}

		maxIdx++
		max2 := digits[maxIdx]
		for i := maxIdx + 1; i < len(digits); i++ {
			if digits[i] > max2 {
				max2 = digits[i]
				maxIdx = i
			}
		}

		total += max1*10 + max2
	}

	return total
}
