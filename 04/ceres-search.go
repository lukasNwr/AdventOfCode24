package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkHorizontal(arr [][]rune, x, y int) bool {
	if x <= len(arr[y])-4 {
		if string(arr[y][x:x+4]) == "XMAS" || string(arr[y][x:x+4]) == "SAMX" {
			return true
		}
	}

	return false
}

func checkVertical(arr [][]rune, x, y int) bool {
	if y <= len(arr)-4 {
		vertical := string([]rune{
			arr[y][x],
			arr[y+1][x],
			arr[y+2][x],
			arr[y+3][x],
		})

		if vertical == "XMAS" || vertical == "SAMX" {
			return true

		}

	}

	return false
}

func checkDiagonal(arr [][]rune, x, y int) (bool, int) {
	nXmasFound := 0
	xmasFound := false
	if x <= len(arr[y])-4 && y <= len(arr)-4 {
		if string(arr[y][x]) == "X" && string(arr[y+1][x+1]) == "M" && string(arr[y+2][x+2]) == "A" && string(arr[y+3][x+3]) == "S" {
			nXmasFound++
			xmasFound = true
		} else if string(arr[y][x]) == "S" && string(arr[y+1][x+1]) == "A" && string(arr[y+2][x+2]) == "M" && string(arr[y+3][x+3]) == "X" {
			nXmasFound++
			xmasFound = true
		}
	}

	if x >= 3 && y <= len(arr)-4 {
		if string(arr[y][x]) == "S" && string(arr[y+1][x-1]) == "A" && string(arr[y+2][x-2]) == "M" && string(arr[y+3][x-3]) == "X" {
			nXmasFound++
			xmasFound = true
		} else if string(arr[y][x]) == "X" && string(arr[y+1][x-1]) == "M" && string(arr[y+2][x-2]) == "A" && string(arr[y+3][x-3]) == "S" {
			nXmasFound++
			xmasFound = true
		}
	}

	return xmasFound, nXmasFound
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	data := make([][]rune, 0)

	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		// Converting the string into a slice of runes so that I can access the individual characters
		runes := []rune(line)
		line_rune := make([]rune, 0)
		line_rune = append(line_rune, runes...)

		data = append(data, line_rune)
		i++
	}

	totalXMAS := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if checkHorizontal(data, x, y) {
				totalXMAS++
				fmt.Printf("Found XMAS at %d %d, totalCount: %d, horizontal\n", y, x, totalXMAS)
			}
			if checkVertical(data, x, y) {
				totalXMAS++
				fmt.Printf("Found XMAS at %d %d, totalCount: %d, vertical\n", y, x, totalXMAS)
			}

			b, v := checkDiagonal(data, x, y)
			if b {
				totalXMAS += v
				fmt.Printf("Found XMAS at %d %d, totalCount: %d, diagonal\n", y, x, totalXMAS)
			}

		}
	}

	fmt.Println("Total XMAS found:", totalXMAS)
}
