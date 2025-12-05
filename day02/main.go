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

	fmt.Println("Part 1:", solvePart1(ranges))
}

func solvePart1(items []string) int {
	total := 0

	for _, item := range items {
		slice := strings.Split(item, "-")
		start, _ := strconv.Atoi(slice[0])
		end, _ := strconv.Atoi(slice[1])

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
