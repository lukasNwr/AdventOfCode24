package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertElem(arr []int, val, idx int) []int {
	if len(arr) < idx {
		return append(arr, val)
	}
	return append(arr[:idx], append([]int{val}, arr[idx:]...)...)
}

func removeElem(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func moveElem(arr []int, fromIdx, toIdx int) []int {
	val := arr[fromIdx]
	return insertElem(removeElem(arr, fromIdx), val, toIdx)
}

func evalRule(prec, succ int, arr []int) bool {
	for i := range arr {
		if succ == arr[i] {
			for j := i; j < len(arr); j++ {
				if arr[j] == prec {
					// Part 2
					moveElem(arr, i, i+1)
					// Part 1
					return false
				} else {
					continue
				}
			}
		}
	}

	return true
}

// Function to check if an array contains a number
func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
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

	correctOrders := make([]int, 0)
	for update_i := range updates {
		for rule_i := range rules {

			// Part 1
			if !evalRule(rules[rule_i][0], rules[rule_i][1], updates[update_i]) {
				wrongOrder = true
				break
			} else {
				continue
			}

		}

		if !wrongOrder {
			// totalSum += updates[update_i][len(updates[update_i])/2]
			if !contains(correctOrders, update_i) {
				correctOrders = append(correctOrders, update_i)
			}
			wrongOrder = false
			continue
		} else {
			wrongOrder = false
			// Part 2
			// totalSum += updates[update_i][len(updates[update_i])/2]
			continue
		}

	}

	// thisRule := false
	for updateIdx := range updates {
		for i := range correctOrders {
			if updateIdx == correctOrders[i] {
				continue
			}
		}

		for {
			allRulesSatisfied := true // Flag to track if all rules are satisfied

			for ruleIdx := range rules {
				// Check if the rule is satisfied on the first try
				if evalRule(rules[ruleIdx][0], rules[ruleIdx][1], updates[updateIdx]) {
					continue // If true, proceed to the next rule
				}

				// If not, keep looping until the rule is satisfied
				for !evalRule(rules[ruleIdx][0], rules[ruleIdx][1], updates[updateIdx]) {
					allRulesSatisfied = false // Mark that not all rules are satisfied
				}
			}

			// Break the outer loop if all rules are satisfied
			if allRulesSatisfied {
				break
			}
		}
	}

	for i := range updates {
		if !contains(correctOrders, i) {
			totalSum += updates[i][len(updates[i])/2]
		}
	}

	fmt.Printf("Total sum: %d\n", totalSum)
}
