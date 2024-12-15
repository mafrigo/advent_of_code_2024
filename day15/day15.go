package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readGrid(filename string, part2 bool) (map[[2]int]string, string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	grid := [][]string{}
	instructions := ""
	instructionsBegun := false
	gridMap := map[[2]int]string{}
	for iline, line := range lines {
		if line == "" {
			instructionsBegun = true
			continue
		} else if !instructionsBegun {
			gridLine := []string{}
			for ichar, char := range line {
				gridLine = append(gridLine, string(char))
				if string(char) != "." {
					if !part2 {
						gridMap[[2]int{iline, ichar}] = string(char)
					} else {
						if string(char) == "O" {
							gridMap[[2]int{iline, 2 * ichar}] = "["
							gridMap[[2]int{iline, 2*ichar + 1}] = "]"
						} else if string(char) == "#" {
							gridMap[[2]int{iline, 2 * ichar}] = "#"
							gridMap[[2]int{iline, 2*ichar + 1}] = "#"
						} else {
							gridMap[[2]int{iline, 2 * ichar}] = "@"
						}
					}
				}
			}
			grid = append(grid, gridLine)
		} else {
			instructions += line
		}
	}
	return gridMap, instructions
}

func printGrid(grid map[[2]int]string) {
	limitY := 0
	limitX := 0
	for loc, _ := range grid {
		if loc[0] > limitY {
			limitY = loc[0]
		}
		if loc[1] > limitX {
			limitX = loc[1]
		}
	}
	for y := 0; y <= limitY; y++ {
		for x := 0; x <= limitX; x++ {
			item, ok := grid[[2]int{y, x}]
			if ok {
				fmt.Print(string(item))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func canItemMove(grid map[[2]int]string, itemLoc [2]int, move [2]int) bool {
	itemCanMove := false
	nextLoc := [2]int{itemLoc[0] + move[0], itemLoc[1] + move[1]}
	itemType, spaceOccupied := grid[nextLoc]
	if !spaceOccupied {
		itemCanMove = true
	} else {
		if itemType == "#" {
			itemCanMove = false
		}
		if itemType == "[" || itemType == "]" {
			itemCanMove = canBigBoxMove(grid, nextLoc, move)
		}
		if itemType == "O" {
			itemCanMove = canItemMove(grid, nextLoc, move)
		}
	}
	return itemCanMove
}

func canBigBoxMove(grid map[[2]int]string, boxLoc [2]int, move [2]int) bool {
	boxType, _ := grid[boxLoc]
	var boxLeftLoc [2]int
	var boxRightLoc [2]int
	if boxType == "[" {
		boxLeftLoc = boxLoc
		boxRightLoc = [2]int{boxLoc[0], boxLoc[1] + 1}
	} else if boxType == "]" {
		boxRightLoc = boxLoc
		boxLeftLoc = [2]int{boxLoc[0], boxLoc[1] - 1}
	}
	if move[1] == -1 {
		return canItemMove(grid, boxLeftLoc, move)
	} else if move[1] == 1 {
		return canItemMove(grid, boxRightLoc, move)
	} else {
		return canItemMove(grid, boxLeftLoc, move) && canItemMove(grid, boxRightLoc, move)
	}
}

func moveItems(grid map[[2]int]string, itemLoc [2]int, move [2]int) map[[2]int]string {
	currentItemType, _ := grid[itemLoc]
	if currentItemType == "[" || currentItemType == "]" {
		if canBigBoxMove(grid, itemLoc, move) {
			if move == [2]int{1, 0} || move == [2]int{-1, 0} {
				var itemLoc2 [2]int
				if currentItemType == "[" {
					itemLoc2 = [2]int{itemLoc[0], itemLoc[1] + 1}
				}
				if currentItemType == "]" {
					itemLoc2 = [2]int{itemLoc[0], itemLoc[1] - 1}
				}
				nextLocation := [2]int{itemLoc[0] + move[0], itemLoc[1] + move[1]}
				nextLocation2 := [2]int{itemLoc2[0] + move[0], itemLoc2[1] + move[1]}
				grid = moveItems(grid, nextLocation, move)
				grid = moveItems(grid, nextLocation2, move)
				grid[nextLocation] = grid[itemLoc]
				grid[nextLocation2] = grid[itemLoc2]
				delete(grid, itemLoc)
				delete(grid, itemLoc2)
			} else {
				nextLocation := [2]int{itemLoc[0] + move[0]*2, itemLoc[1] + move[1]*2}
				nextLocation2 := [2]int{itemLoc[0] + move[0], itemLoc[1] + move[1]}
				grid = moveItems(grid, nextLocation, move)
				grid[nextLocation] = grid[nextLocation2]
				grid[nextLocation2] = grid[itemLoc]
				delete(grid, itemLoc)
			}
		}
	} else if currentItemType == "@" || currentItemType == "O" {
		nextLocation := [2]int{itemLoc[0] + move[0], itemLoc[1] + move[1]}
		if canItemMove(grid, itemLoc, move) {
			grid = moveItems(grid, nextLocation, move)
			grid[nextLocation] = grid[itemLoc]
			delete(grid, itemLoc)
		}
	}
	//printGrid(grid)
	return grid
}

func findRobot(grid map[[2]int]string) [2]int {
	var robotPos [2]int
	for loc, item := range grid {
		if item == "@" {
			robotPos = loc
			break
		}
	}
	return robotPos
}

func moveRobot(grid map[[2]int]string, instructions string) map[[2]int]string {
	moveMap := map[string][2]int{"^": [2]int{-1, 0}, "v": [2]int{1, 0}, "<": [2]int{0, -1}, ">": [2]int{0, 1}}
	for i := 0; i < len(instructions); i++ {
		robotPos := findRobot(grid)
		nextMove := moveMap[string(instructions[i])]
		grid = moveItems(grid, robotPos, nextMove)
	}
	return grid
}

func calcGPSScore(grid map[[2]int]string) int {
	gps := 0
	for loc, item := range grid {
		if item == "O" || item == "[" {
			gps += 100*loc[0] + loc[1]
		}
	}
	return gps
}

func main() {
	//inputFile := "inputtest2"
	inputFile := "input"

	// Part 1
	grid, instructions := readGrid(inputFile, false)
	//printGrid(grid)
	//fmt.Println("Instructions: ", instructions)
	grid = moveRobot(grid, instructions)
	printGrid(grid)
	gps := calcGPSScore(grid)
	fmt.Println("Sum of GPS coordinates: ", gps)

	// Part 2
	grid2, _ := readGrid(inputFile, true)
	//printGrid(grid2)
	grid = moveRobot(grid2, instructions)
	printGrid(grid2)
	gps2 := calcGPSScore(grid2)
	fmt.Println("Sum of GPS coordinates (part2): ", gps2)
}
