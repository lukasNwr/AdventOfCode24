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

// Part 1
// type MoveResult struct {
// 	newX, newY int
// 	explored   bool
// 	leaving    bool
// }

// Part 2
type MoveResult struct {
	newX, newY int
	explored   bool
	leaving    bool
	looping    bool
	oldDir     rune
}

var directions = map[rune]Direction{
	'^': {'^', '>', 0, -1},
	'v': {'v', '<', 0, 1},
	'>': {'>', 'v', 1, 0},
	'<': {'<', '^', -1, 0},
}

func deepCopy(guardMap [][]rune) [][]rune {
	// Create a new slice with the same length as the original
	mapCopy := make([][]rune, len(guardMap))

	// Copy each slice from guardMap into mapCopy
	for i, guard := range guardMap {
		mapCopy[i] = make([]rune, len(guard)) // Allocate space for each slice
		copy(mapCopy[i], guard)               // Copy data into mapCopy[i]
	}

	return mapCopy
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

func moveGuard(mapArray [][]rune, posX, posY int, dir rune, visited []MoveResult) MoveResult {
	direction := directions[dir]
	newX, newY := posX+direction.deltaX, posY+direction.deltaY
	looping := false

	for _, visited := range visited {
		if visited.newX == newX && visited.newY == newY {
			looping = true
		}
	}

	// Check if guard is at map boundary
	if isAtBoundary(mapArray, newX, newY) {
		mapArray[posY][posX] = dir
		return MoveResult{posX, posY, false, true, looping, dir}
	}

	// Check if guard hit a wall
	if isWall(mapArray, newX, newY) {
		rotateGuard(mapArray, posX, posY, direction.nextDir)
		return MoveResult{posX, posY, false, false, looping, dir}
	}

	return handleMovement(mapArray, posX, posY, newX, newY, dir, looping)
}

func isAtBoundary(mapArray [][]rune, x, y int) bool {
	return y < 0 || y >= len(mapArray) ||
		x < 0 || x >= len(mapArray[0])
}
func isWall(mapArray [][]rune, x, y int) bool {
	return mapArray[y][x] == '#'
}
func handleMovement(mapArray [][]rune, oldX, oldY, newX, newY int, dir rune, looping bool) MoveResult {
	switch mapArray[newY][newX] {

	case '.':
		// Moving to unexplored territory
		mapArray[oldY][oldX] = 'x'
		mapArray[newY][newX] = dir
		return MoveResult{newX, newY, true, false, looping, dir}
	case 'x':
		// Moving to already explored territory
		mapArray[oldY][oldX] = 'x'
		mapArray[newY][newX] = dir
		return MoveResult{newX, newY, false, false, looping, dir}
	default:
		return MoveResult{oldX, oldY, false, false, looping, dir}
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

	// Create a clean copy of the original map before running Part 1
	originalMap := deepCopy(guardMap)

	// Part 1
	startingPosX, startingPosY := findInitialPos(guardMap)
	if startingPosX == -1 || startingPosY == -1 {
		panic("Could not find initial position")
	}

	// 1 because we include starting position
	totalVisited := 1
	looping := false

	moveResult := MoveResult{startingPosX, startingPosY, false, false, looping, guardMap[startingPosY][startingPosX]}
	var visited []MoveResult

	for {
		posX, posY := moveResult.newX, moveResult.newY
		moveResult = moveGuard(guardMap, posX, posY, guardMap[posY][posX], visited)

		// printMapReplace(guardMap, totalVisited)
		// time.Sleep(200 * time.Millisecond)

		if moveResult.leaving {
			break
		}

		if moveResult.explored {
			totalVisited++
			visited = append(visited, moveResult)
		}
	}


	// Part 2
	bads := 0
	initialDir := originalMap[startingPosY][startingPosX]
	
	// Test each position the guard visited from Part 1
	for i := 0; i < len(visited); i++ {
		// Create fresh map from the original
		testMap := deepCopy(originalMap)
		
		// Block the test position with a wall
		testPos := visited[i]
		testMap[testPos.newY][testPos.newX] = '#'
		
		// Start fresh path finding from the beginning
		currX, currY := startingPosX, startingPosY
		testMap[startingPosY][startingPosX] = initialDir
		
		moveResult := MoveResult{currX, currY, false, false, false, initialDir}
		testVisited := 1
		
		// Track position+direction combinations to detect loops
		type posDir struct {
			x, y int
			dir  rune
		}
		visited := make(map[posDir]bool)
		
		for {
			posX, posY := moveResult.newX, moveResult.newY
			currentDir := testMap[posY][posX]
			
			// Check for loop
			pd := posDir{posX, posY, currentDir}
			if visited[pd] {
				bads++
				break
			}
			visited[pd] = true
			
			moveResult = moveGuard(testMap, posX, posY, testMap[posY][posX], nil)
			
			if moveResult.leaving {
				break  // Guard found the exit
			}
			
			if moveResult.explored {
				testVisited++
			}

			// printMapReplace(testMap, testVisited)
			// time.Sleep(50 * time.Millisecond)
		}
		
		// time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Total visited in Part 1: ", totalVisited)
	fmt.Println("Bad positions in Part 2: ", bads)
}

