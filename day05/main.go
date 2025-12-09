package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day05/input.txt")
	defer file.Close()

	var ranges [][2]int
	var items []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}

		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, [2]int{start, end})
	}

	for scanner.Scan() {
		line := scanner.Text()
		item, _ := strconv.Atoi(line)
		items = append(items, item)
	}

	fmt.Println("Part 1:", solvePart1(ranges, items)) // 735
	fmt.Println("Part 2:", solvePart2(ranges))        // 344306344403172
}

func solvePart1(ranges [][2]int, items []int) int {
	total := 0

	for _, item := range items {
		foundAlready := false
		for _, r := range ranges {
			if item >= r[0] && item <= r[1] && !foundAlready {
				total++
				foundAlready = true
			}
		}
	}

	return total
}

func solvePart2(ranges [][2]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][2]int{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		prev := &merged[len(merged)-1]
		current := ranges[i]

		if current[0] <= prev[1]+1 {
			if current[1] > prev[1] {
				prev[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	total := 0
	for _, r := range merged {
		total += r[1] - r[0] + 1
	}

	return total
}
