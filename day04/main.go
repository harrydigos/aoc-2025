package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("day04/input.txt")
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	fmt.Println("Part 1:", solvePart1(grid))
}

func solvePart1(grid [][]rune) int {
	total := 0

	directions := [][2]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '@' {
				continue
			}

			foundPapers := 0
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) {
					if grid[ni][nj] == '@' {
						foundPapers++
					}
				}
			}

			if foundPapers < 4 {
				total++
			}
		}
	}
	return total
}
