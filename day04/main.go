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

	fmt.Println("Part 1:", solvePart1(grid)) // 1395
	fmt.Println("Part 2:", solvePart2(grid)) // 8451
}

var directions = [][2]int{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

func solvePart1(grid [][]rune) int {
	total := 0

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

type Position struct {
	row int
	col int
}

func solvePart2(grid [][]rune) int {
	total := 0
	removed := make(map[Position]bool)
	lastRoundFound := 1

	for lastRoundFound > 0 {
		lastRoundFound = 0

		for i := range grid {
			for j := range grid[i] {
				if (grid[i][j] != '@' || removed[Position{i, j}]) {
					continue
				}

				foundPapers := 0
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[ni]) {
						if grid[ni][nj] == '@' && !removed[Position{ni, nj}] {
							foundPapers++
						}
					}
				}

				if foundPapers < 4 {
					total++
					removed[Position{i, j}] = true
					lastRoundFound++
				}
			}
		}
	}

	return total
}
