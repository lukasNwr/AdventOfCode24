package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func orOp(a, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	combined := aStr + bStr

	num, _ := strconv.Atoi(combined)
	return num
}

func findCombination(numbers []int, testValue int) bool {
	var backtrack func(index, currentValue int, currentOp string) bool
	backtrack = func(index, currentValue int, currentOp string) bool {

		// Check if the value is what am I looking for when I reach the end of the array
		if index == len(numbers) {
			return currentValue == testValue
		}

		operators := []string{"+", "*", "||"}

		for _, op := range operators {
			// Initial number, just do the backtrack from next number
			if index == 0 {
				if backtrack(index+1, numbers[index], op) {
					return true
				}
			} else {
				newValue := 0
				switch currentOp {
				case "+":
					newValue = currentValue + numbers[index]
				case "*":
					newValue = currentValue * numbers[index]
				case "||":
					newValue = orOp(currentValue, numbers[index])
				}

				if backtrack(index+1, newValue, op) {
					return true
				}
			}
		}

		return false
	}

	return backtrack(0, 0, "+") || backtrack(0, 0, "*") || backtrack(0, 0, "||")
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	totalCalibration := 0

	for scanner.Scan() {
		var numbers []int

		line := scanner.Text()
		processed := strings.Split(line, ":")
		testValue, _ := strconv.Atoi(processed[0])

		nums := strings.Split(processed[1], " ")
		for _, v := range nums {
			num, _ := strconv.Atoi(v)
			numbers = append(numbers, num)
		}

		if findCombination(numbers, testValue) {
			totalCalibration += testValue
		}
	}
	fmt.Println("Total calibration: ", totalCalibration)
}
