package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkSafety(report []int) bool {
	temp := report[0]
	dec := false
	dampened := false

	if report[0] < report[1] {
		dec = false
	} else {
		dec = true
	}

	for _, v := range report[1:] {
		if dec && temp < v {
			if !dampened {
				dampened = true
				temp = v
				continue
			}
			return false
		}
		if !dec && temp > v {
			if !dampened {
				dampened = true
				temp = v
				continue
			}
			return false
		}

		if math.Abs(float64(v-temp)) > 3 || v == temp {
			if !dampened {
				dampened = true
				temp = v
				continue
			}
			return false
		}

		temp = v
	}

	return true
}

func main() {
	f, err := os.Open("report_input.txt")
	if err != nil {
		panic(err)
	}

	data := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		report := make([]int, 0)

		// Split the string on line into substrings
		for _, v := range strings.Split(line, " ") {
			// _ is error, but I don't need it so it's ignored, i = int
			i, _ := strconv.Atoi(v)
			report = append(report, i)
		}

		data = append(data, report)
	}

	totalSafe := 0
	for _, v := range data {
		if checkSafety(v) {
			totalSafe++
		}
	}

	fmt.Println("Total safe reports: ", totalSafe)

}
