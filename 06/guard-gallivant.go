package main

import (
	"bufio"
	"fmt"
	"os"
)

// Direction represents the guard's movement direction
type Direction struct {
	symbol  rune
	nextDir rune
	deltaX  int
	deltaY  int
}

type MoveResult struct {
	newX, newY int
	explored   bool
	leaving    bool
}

var directions = map[rune]Direction{
	'^': {'^', '>', 0, -1},
	'v': {'v', '<', 0, 1},
	'>': {'>', 'v', 1, 0},
	'<': {'<', '^', -1, 0},
}

func colorize(s rune) string {
	switch s {
	case 'x':
		// blue
		return "\033[38;5;32m"

	case '.':
		// white
		return "\033[38;5;255m"
	case '#':
		// orange
		return "\033[38;5;214m"
	}
	// green
	return "\033[38;5;46m"
}

func printMap(mapArray [][]rune, nMoves int) {
	for _, row := range mapArray {
		fmt.Println(string(row))
	}

	fmt.Println("Moves: ", nMoves)
	fmt.Println("----------------------------")
}

// Function to print the map and replace the old one
func printMapReplace(mapArray [][]rune, nMoves int) {
	// Clear the screen by printing enough newlines (you can adjust this as needed)
	fmt.Print("\033[H\033[2J") // ANSI escape code to clear the terminal screen

	// Print the updated map

	for _, row := range mapArray {
		for _, val := range row {
			// Print the rune with its color
			fmt.Print(colorize(val), string(val), "\033[0m") // Color the rune and reset color after printing
		}
		fmt.Println()
	}

	// Print the number of moves
	fmt.Println("Moves: ", nMoves)
	fmt.Println("----------------------------")
}

func findInitialPos(mapArray [][]rune) (int, int) {
	for posY, row := range mapArray {
		for posX, val := range row {
			if val == '^' || val == 'v' || val == '<' || val == '>' {
				return posX, posY
			}
		}
	}
	return -1, -1
}

func rotateGuard(mapArray [][]rune, currPosX, currPosY int, newDir rune) {
	mapArray[currPosY][currPosX] = newDir
}

func moveGuard(mapArray [][]rune, posX, posY int, dir rune) MoveResult {
	direction := directions[dir]
	newX, newY := posX+direction.deltaX, posY+direction.deltaY

	// Check if guard is at map boundary
	if isAtBoundary(mapArray, newX, newY) {
		mapArray[posY][posX] = dir
		return MoveResult{posX, posY, false, true}
	}

	// Check if guard hit a wall
	if isWall(mapArray, newX, newY) {
		rotateGuard(mapArray, posX, posY, direction.nextDir)
		return MoveResult{posX, posY, false, false}
	}

	return handleMovement(mapArray, posX, posY, newX, newY, dir)
}

func isAtBoundary(mapArray [][]rune, x, y int) bool {
	return y < 0 || y >= len(mapArray) ||
		x < 0 || x >= len(mapArray[0])
}

func isWall(mapArray [][]rune, x, y int) bool {
	return mapArray[y][x] == '#'
}

func handleMovement(mapArray [][]rune, oldX, oldY, newX, newY int, dir rune) MoveResult {
	switch mapArray[newY][newX] {
	case '.':
		// Moving to unexplored territory
		mapArray[oldY][oldX] = 'x'
		mapArray[newY][newX] = dir
		return MoveResult{newX, newY, true, false}
	case 'x':
		// Moving to already explored territory
		mapArray[oldY][oldX] = 'x'
		mapArray[newY][newX] = dir
		return MoveResult{newX, newY, false, false}
	default:
		return MoveResult{oldX, oldY, false, false}
	}
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Create a 2D array to store the map
	var guardMap [][]rune

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// Convert the string to a rune slice
		line := []rune(scanner.Text())
		guardMap = append(guardMap, line)
	}

	currentPosX, currentPosY := findInitialPos(guardMap)
	if currentPosX == -1 || currentPosY == -1 {
		panic("Could not find initial position")
	}

	// 1 because we include starting position
	totalVisited := 1
	i := 0

	moveResult := MoveResult{currentPosX, currentPosY, false, false}

	for {

		posX, posY := moveResult.newX, moveResult.newY

		moveResult = moveGuard(guardMap, posX, posY, guardMap[posY][posX])

		// fmt.Println("Move: ", totalVisited)
		// printMap(guardMap)
		printMapReplace(guardMap, totalVisited)
		// time.Sleep(200 * time.Millisecond)

		if moveResult.leaving {
			break
		}

		if moveResult.explored {
			totalVisited++
		}
		i++
	}

	fmt.Println("Total visited: ", totalVisited)

}
