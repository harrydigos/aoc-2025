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

func isValidOperation(operation rune) bool {
	return operation == opAdd || operation == opMul
}

func parseInput(filename string) ([][]int, []rune) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var problems [][]int
	var operations []rune

	for scanner.Scan() {
		line := scanner.Text()
		columnsStr := strings.Fields(line)

		if isValidOperation(rune(columnsStr[0][0])) {
			for _, op := range columnsStr {
				operations = append(operations, rune(op[0]))
			}
			continue
		}

		for i, column := range columnsStr {
			num, _ := strconv.Atoi(column)
			if i >= len(problems) {
				problems = append(problems, []int{})
			}
			problems[i] = append(problems[i], num)
		}
	}

	return problems, operations
}

func main() {
	problems, operations := parseInput("day06/input.txt")
	fmt.Println("Part 1:", solvePart1(problems, operations)) // 3968933219902
}

func solvePart1(problems [][]int, operations []rune) int {
	total := 0
	for i, problem := range problems {
		op := operations[i]
		result := problem[0]
		for j := 1; j < len(problem); j++ {
			switch op {
			case opAdd:
				result += problem[j]
			case opMul:
				result *= problem[j]
			}
		}
		total += result
	}
	return total
}
