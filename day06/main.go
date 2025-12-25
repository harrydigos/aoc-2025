package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	opAdd = '+'
	opMul = '*'
)

func main() {
	lines := readLines("day06/input.txt")
	fmt.Println("Part 1:", solve(parseHorizontal(lines))) // 3968933219902
	fmt.Println("Part 2:", solve(parseVertical(lines)))   // 6019576291014
}

func solve(problems [][]int, operations []rune) int {
	total := 0
	for i, problem := range problems {
		result := problem[0]
		for _, n := range problem[1:] {
			if operations[i] == opAdd {
				result += n
			} else {
				result *= n
			}
		}
		total += result
	}
	return total
}

func readLines(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseHorizontal(lines []string) ([][]int, []rune) {
	var problems [][]int
	var operations []rune

	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0][0] == opAdd || fields[0][0] == opMul {
			for _, op := range fields {
				operations = append(operations, rune(op[0]))
			}
			continue
		}

		for i, field := range fields {
			num, _ := strconv.Atoi(field)
			if i >= len(problems) {
				problems = append(problems, []int{})
			}
			problems[i] = append(problems[i], num)
		}
	}

	return problems, operations
}

func parseVertical(lines []string) ([][]int, []rune) {
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	for i := range lines {
		for len(lines[i]) < maxLen {
			lines[i] += " "
		}
	}

	dataRows := lines[:len(lines)-1]
	opRow := lines[len(lines)-1]

	var problems [][]int
	var operations []rune

	for col := 0; col < maxLen; {
		for col < maxLen && isSpaceColumn(dataRows, col) {
			col++
		}
		if col >= maxLen {
			break
		}

		var numbers []int
		var op rune
		for col < maxLen && !isSpaceColumn(dataRows, col) {
			numbers = append(numbers, readColumnNumber(dataRows, col))
			if opRow[col] == opAdd || opRow[col] == opMul {
				op = rune(opRow[col])
			}
			col++
		}

		problems = append(problems, numbers)
		operations = append(operations, op)
	}

	return problems, operations
}

func isSpaceColumn(rows []string, col int) bool {
	for _, row := range rows {
		if row[col] != ' ' {
			return false
		}
	}
	return true
}

func readColumnNumber(rows []string, col int) int {
	var s strings.Builder
	for _, row := range rows {
		if row[col] >= '0' && row[col] <= '9' {
			s.WriteByte(row[col])
		}
	}
	n, _ := strconv.Atoi(s.String())
	return n
}
