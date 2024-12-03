package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func checkMul(batch []byte) bool {
	// This will be done only if in the main loop we find "m" char
	// Check if the string is proper "mul" string
	// return the boolead based the check

	if string(batch[0]) == "m" && string(batch[1]) == "u" && string(batch[2]) == "l" {
		// fmt.Println("Found mul")
		return true
	} else {
		// fmt.Println("Not found mul", string(batch[0]), string(batch[1]), string(batch[2]))
		return false
	}

}

func checkPar(batch []byte) (int, int) {
	// This will be done only if the checkMul returns true
	// Check the paranthesis content/multilply
	// return the num of idx (length) and multiply value

	re := regexp.MustCompile(`^[0-9]+$`)
	idxCounter := 0

	if string(batch[0]) != "(" {
		return 0, 0
	} else {
		idxCounter++
	}

	digitsCounter := 0
	firstNumberS := ""
	secondNumberS := ""
	secondNumber := false

	result := 0

	for i := 1; i < len(batch); i++ {
		if !re.MatchString(string(batch[i])) {
			if len(firstNumberS) > 0 && string(batch[i]) == "," {
				idxCounter++
				digitsCounter = 0
				secondNumber = true
				continue
			}
			if len(firstNumberS) > 0 && len(secondNumberS) > 0 && string(batch[i]) == ")" {
				idxCounter++
				// Convert string to int
				firstNumber, _ := strconv.Atoi(firstNumberS)
				secondNumber, _ := strconv.Atoi(secondNumberS)
				return idxCounter, firstNumber * secondNumber
			}
			return 0, 0
		} else {
			if digitsCounter > 3 {
				return 0, 0
			}

			if !secondNumber {
				firstNumberS += string(batch[i])
			} else {
				secondNumberS += string(batch[i])
			}

			digitsCounter++
			idxCounter++
		}
	}

	return idxCounter, result
}

func main() {
	// Try to open the file
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Read the file into memory (this can be problematic with big files)
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	totalSum := 0

	for i, v := range data {
		if v == 'm' {
			endIdx := i + 3
			if !checkMul(data[i:endIdx]) {
				continue
			} else {
				i += 3
				endIdx = i + 9
				// fmt.Println("par: ", string(data[i:endIdx]))
				parLenght, result := checkPar(data[i:endIdx])
				i += parLenght
				totalSum += result
			}

		}
	}

	fmt.Printf("Total sum: %d\n", totalSum)
}
