package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func partition(low, high int, arr []int) int {
	p := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < p {
			i++
			swap(i, j, arr)
		}
	}

	swap(i+1, high, arr)

	// Return the index of the pivot
	return i + 1
}

func swap(idx1, idx2 int, arr []int) {
	arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
}

func quicksort(lowIdx, highIdx int, arr []int) {
	if lowIdx < highIdx {
		pivotIdx := partition(lowIdx, highIdx, arr)
		quicksort(lowIdx, pivotIdx-1, arr)
		quicksort(pivotIdx+1, highIdx, arr)
	}
}

func countOccurences(arr []int, target int) int {
	result := 0

	for _, v := range arr {
		if v == target {
			result++
		}
	}

	return result
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	// Read the file line by line with the use of scanner
	scanner := bufio.NewScanner(f)

	// Split the input into words
	scanner.Split(bufio.ScanWords)

	// Close the file
	defer f.Close()

	var i int = 0
	// Create two slices to store the left and right side of the input with the initial size of 0
	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		i++
		// Convert the string to an integer
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		// Append the number to the left or right slice depending on the index
		if i%2 == 0 {
			right = append(right, num)
		} else {
			left = append(left, num)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort the left slice
	low := 0
	high := len(left) - 1
	quicksort(low, high, left)

	// Sort the right slice
	low = 0
	high = len(right) - 1
	quicksort(low, high, right)

	totalDiff := 0
	for i, v := range left {
		totalDiff += int(math.Abs(float64(v - right[i])))
	}

	fmt.Printf("The total difference is: %d\n", totalDiff)

	// Part 2
	iterVals := make(map[int]int)
	similarityScore := 0

	for _, v := range left {
		// Check if the value is already in the map
		val, ok := iterVals[v]
		if ok {
			similarityScore += v * val
		} else {
			iterVals[v] = countOccurences(right, v)
			similarityScore += v * iterVals[v]
		}
	}

	fmt.Printf("The similarity score is: %d\n", similarityScore)

	// Print the left slice
	// fmt.Printf("%v", left)

}
