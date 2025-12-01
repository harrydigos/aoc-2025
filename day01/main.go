package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("day01/input.txt")
	lines := strings.Split(string(input), "\n")

	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func rotate(value int, direction rune, amount int) int {
	switch direction {
	case 'L':
		return (value - amount%100 + 100) % 100
	case 'R':
		return (value + amount%100) % 100
	default:
		return value
	}
}

func solvePart1(lines []string) int {
	count, dial := 0, 50

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		direction := rune(line[0])
		amount, _ := strconv.Atoi(line[1:])
		dial = rotate(dial, direction, amount)

		if dial == 0 {
			count++
		}
	}

	return count
}

func solvePart2(lines []string) int {
	count, dial := 0, 50

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		direction := rune(line[0])
		amount, _ := strconv.Atoi(line[1:])

		if direction == 'R' {
			count += (dial + amount) / 100
		} else {
			if dial == 0 {
				count += amount / 100
			} else if amount >= dial {
				count += 1 + (amount-dial)/100
			}
		}

		dial = rotate(dial, direction, amount)

	}

	return count
}
