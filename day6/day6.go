package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_lines(filename string) []string {
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
	return lines
}

type guard struct {
	position  [2]int
	direction string
}

func findGuard(lines []string) guard {
	myGuard := guard{position: [2]int{0, 0}, direction: "^"}
	for iline, line := range lines {
		for ichar, char := range line {
			if string(char) == "^" {
				myGuard.position = [2]int{iline, ichar}
			}
		}
	}
	return myGuard
}

func findObstacles(lines []string) [][2]int {
	obstacles := [][2]int{}
	for iline, line := range lines {
		for ichar, char := range line {
			if string(char) == "#" {
				obstacles = append(obstacles, [2]int{iline, ichar})
			}
		}
	}
	return obstacles
}

func turn90deg(myGuard guard) guard {
	turnMap := map[string]string{
		"^": ">",
		">": "v",
		"v": "<",
		"<": "^",
	}
	myGuard.direction = turnMap[myGuard.direction]
	return myGuard
}

func moveGuard(previousGuard guard, obstacles [][2]int, positionMap map[[2]int]string, guardMap map[guard]string, mapLimit [2]int) map[[2]int]string {
	nextPosition := previousGuard.position
	nextGuard := previousGuard
	if previousGuard.direction == "^" {
		nextPosition[0] = previousGuard.position[0] - 1
	}
	if previousGuard.direction == "v" {
		nextPosition[0] = previousGuard.position[0] + 1
	}
	if previousGuard.direction == ">" {
		nextPosition[1] = previousGuard.position[1] + 1
	}
	if previousGuard.direction == "<" {
		nextPosition[1] = previousGuard.position[1] - 1
	}
	foundObstacle := false
	for _, obstacle := range obstacles {
		if nextPosition == obstacle {
			foundObstacle = true
		}
	}
	if foundObstacle {
		nextGuard = turn90deg(nextGuard)
	} else {
		nextGuard.position = nextPosition
	}
	//fmt.Println(nextGuard.position, nextGuard.direction)

	_, ok := guardMap[nextGuard]
	if !ok {
		guardMap[nextGuard] = "#"
	} else {
		return map[[2]int]string{} //return empty map in case of closed loop
	}

	if nextGuard.position[0] >= 0 && nextGuard.position[0] < mapLimit[0] && nextGuard.position[1] >= 0 && nextGuard.position[1] < mapLimit[1] {
		positionMap[nextGuard.position] = "#"
		return moveGuard(nextGuard, obstacles, positionMap, guardMap, mapLimit)
	} else {
		return positionMap
	}
}

func main() {
	//input_file := "inputtest"
	input_file := "input"
	lines := read_lines(input_file)
	fmt.Println(lines)
	myGuard := findGuard(lines)
	fmt.Println(myGuard)
	obstacles := findObstacles(lines)
	fmt.Println(obstacles)
	positionMap := map[[2]int]string{myGuard.position: "#"}
	guardMap := map[guard]string{}
	positionMap = moveGuard(myGuard, obstacles, positionMap, guardMap, [2]int{len(lines), len(lines[0])})
	fmt.Println(positionMap)
	fmt.Println("Number of positions covered by the guard")
	fmt.Println(len(positionMap))

	nClosedLoops := 0
	for iline, line := range lines {
		fmt.Println(iline)
		for ichar, _ := range line {
			newObstacle := [2]int{iline, ichar}
			positionMap = map[[2]int]string{myGuard.position: "#"}
			guardMap = map[guard]string{}
			positionMap = moveGuard(myGuard, append(obstacles, newObstacle), positionMap, guardMap, [2]int{len(lines), len(lines[0])})
			if len(positionMap) == 0 {
				nClosedLoops = nClosedLoops + 1
			}
		}
	}
	fmt.Println("Number of closed loops after adding an obstacle")
	fmt.Println(nClosedLoops)
}
