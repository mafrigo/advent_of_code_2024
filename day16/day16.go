package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readGrid(filename string) [][]string {
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
	for _, line := range lines {
		gridLine := []string{}
		for _, char := range line {
			gridLine = append(gridLine, string(char))
		}
		grid = append(grid, gridLine)
	}
	return grid
}

func printGrid(grid [][]string, path [][2]int) {
	for iline, line := range grid {
		for ichar, char := range line {
			inPath := false
			for _, cell := range path {
				if cell == [2]int{iline, ichar} {
					inPath = true
				}
			}
			if inPath && string(char) != "S" && string(char) != "E" {
				fmt.Print(string("O"))
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

type positionOrientation struct {
	position    [2]int
	orientation string
}

var scoreMap map[positionOrientation]int

func findEndScore(grid [][]string, scoreMap map[positionOrientation]int) (positionOrientation, int) {
	var endPos [2]int
	for iline, line := range grid {
		for ichar, char := range line {
			if string(char) == "E" {
				endPos = [2]int{iline, ichar}
			}
		}
	}
	minScore := 10000000000000
	var endPosOri positionOrientation
	orientations := []string{">", "<", "^", "v"}
	for _, orientation := range orientations {
		if scoreMap[positionOrientation{position: endPos, orientation: orientation}] < minScore {
			minScore = scoreMap[positionOrientation{position: endPos, orientation: orientation}]
			endPosOri = positionOrientation{position: endPos, orientation: orientation}
		}
	}
	return endPosOri, minScore
}

func initScoreMap(grid [][]string) map[positionOrientation]int {
	scoreMap = make(map[positionOrientation]int)
	orientations := []string{">", "<", "^", "v"}
	for iline, line := range grid {
		for ichar, char := range line {
			for _, orientation := range orientations {
				if string(char) == "S" && orientation == ">" {
					scoreMap[positionOrientation{position: [2]int{iline, ichar}, orientation: orientation}] = 0
				} else if string(char) != "#" {
					scoreMap[positionOrientation{position: [2]int{iline, ichar}, orientation: orientation}] = 1000000000000000000
				} else {
					continue
				}
			}
		}
	}
	return scoreMap
}

func populateScoreMap(grid [][]string) map[positionOrientation]int {
	scoreMap = initScoreMap(grid)
	unvisitedMap := make(map[positionOrientation]int)
	for posOri, _ := range scoreMap {
		unvisitedMap[posOri] = 1
	}

	moveStraightMap := map[string][2]int{"^": [2]int{-1, 0}, "v": [2]int{1, 0}, "<": [2]int{0, -1}, ">": [2]int{0, 1}}
	moveLeftMap := map[string]string{"^": "<", "v": ">", "<": "v", ">": "^"}
	moveRightMap := map[string]string{"^": ">", "v": "<", "<": "^", ">": "v"}
	for i := 0; i < len(scoreMap); i++ {
		if i%500 == 0 {
			fmt.Println(int(100*float64(i)/float64(len(scoreMap))), " %")
		}
		currentPositionOrientation := getSmallestScoreUnvisitedLocation(scoreMap, unvisitedMap)
		if currentPositionOrientation.position == [2]int{0, 0} {
			break
		}
		//Move straight
		straightPosition := [2]int{currentPositionOrientation.position[0] + moveStraightMap[currentPositionOrientation.orientation][0], currentPositionOrientation.position[1] + moveStraightMap[currentPositionOrientation.orientation][1]}
		if grid[straightPosition[0]][straightPosition[1]] != "#" {
			straightMove := positionOrientation{position: straightPosition, orientation: currentPositionOrientation.orientation}
			if scoreMap[currentPositionOrientation]+1 < scoreMap[straightMove] {
				scoreMap[straightMove] = scoreMap[currentPositionOrientation] + 1
			}
		}
		// Turn 90 degrees
		leftMove := positionOrientation{position: currentPositionOrientation.position, orientation: moveLeftMap[currentPositionOrientation.orientation]}
		if scoreMap[currentPositionOrientation]+1000 < scoreMap[leftMove] {
			scoreMap[leftMove] = scoreMap[currentPositionOrientation] + 1000
		}
		rightMove := positionOrientation{position: currentPositionOrientation.position, orientation: moveRightMap[currentPositionOrientation.orientation]}
		if scoreMap[currentPositionOrientation]+1000 < scoreMap[rightMove] {
			scoreMap[rightMove] = scoreMap[currentPositionOrientation] + 1000
		}
		delete(unvisitedMap, currentPositionOrientation)
	}

	return scoreMap
}

func getSmallestScoreUnvisitedLocation(scoreMap map[positionOrientation]int, unvisitedMap map[positionOrientation]int) positionOrientation {
	smallestPosOri := positionOrientation{position: [2]int{0, 0}, orientation: ">"}
	smallestScore := 1000000000000000000
	for posOri, _ := range unvisitedMap {
		score := scoreMap[posOri]
		if score < smallestScore {
			smallestPosOri = posOri
			smallestScore = score
		}
	}
	return smallestPosOri
}

func getBestPaths(scoreMap map[positionOrientation]int, end positionOrientation) [][2]int {
	moveStraightMap := map[string][2]int{"^": [2]int{-1, 0}, "v": [2]int{1, 0}, "<": [2]int{0, -1}, ">": [2]int{0, 1}}
	moveLeftMap := map[string]string{"^": "<", "v": ">", "<": "v", ">": "^"}
	moveRightMap := map[string]string{"^": ">", "v": "<", "<": "^", ">": "v"}
	posOrisBestPath := []positionOrientation{end}
	for i := 0; i < len(scoreMap); i++ {
		nNewPosOris := 0
		currentPositionOrientation := posOrisBestPath[i]
		straightPosition := [2]int{currentPositionOrientation.position[0] - moveStraightMap[currentPositionOrientation.orientation][0], currentPositionOrientation.position[1] - moveStraightMap[currentPositionOrientation.orientation][1]}
		straightMove := positionOrientation{position: straightPosition, orientation: currentPositionOrientation.orientation}
		if scoreMap[straightMove] == scoreMap[currentPositionOrientation]-1 {
			posOrisBestPath = append(posOrisBestPath, straightMove)
			nNewPosOris++
		}
		leftMove := positionOrientation{position: currentPositionOrientation.position, orientation: moveLeftMap[currentPositionOrientation.orientation]}
		if scoreMap[leftMove] == scoreMap[currentPositionOrientation]-1000 {
			posOrisBestPath = append(posOrisBestPath, leftMove)
			nNewPosOris++
		}
		rightMove := positionOrientation{position: currentPositionOrientation.position, orientation: moveRightMap[currentPositionOrientation.orientation]}
		if scoreMap[rightMove] == scoreMap[currentPositionOrientation]-1000 {
			posOrisBestPath = append(posOrisBestPath, rightMove)
			nNewPosOris++
		}
		if nNewPosOris == 0 {
			break
		}
	}
	posBestPath := [][2]int{}
	for _, posOri := range posOrisBestPath {
		alreadyInList := false
		for _, pos := range posBestPath {
			if posOri.position == pos {
				alreadyInList = true
				break
			}
		}
		if !alreadyInList {
			posBestPath = append(posBestPath, posOri.position)
		}
	}
	return posBestPath
}

func main() {
	inputFile := "inputtest2"
	//inputFile := "input"
	grid := readGrid(inputFile)
	//printGrid(grid, [][2]int{})
	scoreMap = populateScoreMap(grid)
	//fmt.Println(scoreMap)
	endPosOri, minScore := findEndScore(grid, scoreMap)
	fmt.Println("Minimum score (part1): ", minScore)

	pathCells := getBestPaths(scoreMap, endPosOri)
	printGrid(grid, pathCells)
	fmt.Println("Number of cells on one of the best paths (part2): ", len(pathCells))
}
