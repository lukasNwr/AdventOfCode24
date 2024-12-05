package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evalRule(prec, succ int, arr []int) bool {
	for i := range arr {
		if succ == arr[i] {
			for j := i; j < len(arr); j++ {
				if arr[j] == prec {
					// Found the wrong order
					return false
				} else {
					continue
				}
			}
		}
	}

	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var rules [][]int
	var updates [][]int
	newLineFound := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			newLineFound = true
			continue
		}

		if newLineFound {
			a := strings.Split(line, ",")
			nums := make([]int, len(a))
			for i := range a {
				nums[i], _ = strconv.Atoi(a[i])
			}
			updates = append(updates, nums)

		} else {
			a := strings.Split(line, "|")
			nums := make([]int, len(a))
			for i := range a {
				nums[i], _ = strconv.Atoi(a[i])
			}
			rules = append(rules, nums)
		}
	}

	totalSum := 0
	wrongOrder := false

	for update_i := range updates {
		for rule_i := range rules {
			if !evalRule(rules[rule_i][0], rules[rule_i][1], updates[update_i]) {
				wrongOrder = true
				break
			} else {
				continue
			}
		}

		fmt.Println("Rule: ", wrongOrder)

		if !wrongOrder {
			totalSum += updates[update_i][len(updates[update_i])/2]
			wrongOrder = false
		} else {
			wrongOrder = false
			continue
		}

	}

	fmt.Printf("Total sum: %d\n", totalSum)
}
