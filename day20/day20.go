package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readGrid(filename string) ([2]int, [2]int, map[[2]int]bool) {
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
	var start [2]int
	var end [2]int
	mapObstacles := map[[2]int]bool{}
	for iline, line := range lines {
		for ichar, char := range line {
			if string(char) == "#" {
				mapObstacles[[2]int{iline, ichar}] = true
			} else {
				mapObstacles[[2]int{iline, ichar}] = false
			}
			if string(char) == "S" {
				start = [2]int{iline, ichar}
			}
			if string(char) == "E" {
				end = [2]int{iline, ichar}
			}
		}
	}
	return start, end, mapObstacles
}

func getMinMax(mapObstacles map[[2]int]bool) (int, int) {
	max := 0
	min := 100000000000
	for pos, _ := range mapObstacles {
		if pos[0] > max {
			max = pos[0]
		}
		if pos[0] < min {
			min = pos[0]
		}
	}
	return min, max
}

func printGrid(start [2]int, end [2]int, mapObstacles map[[2]int]bool) {
	min, max := getMinMax(mapObstacles)
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			pos := [2]int{i, j}
			if mapObstacles[pos] {
				fmt.Print("#")
			} else if pos == start {
				fmt.Print("S")
			} else if pos == end {
				fmt.Print("E")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getShortestDistance(mapUnvisited map[[2]int]bool, mapDistance map[[2]int]int, infinity int) [2]int {
	shortest := [2]int{0, 0}
	distance := infinity
	for loc, _ := range mapUnvisited {
		if mapDistance[loc] < distance {
			shortest = loc
			distance = mapDistance[loc]
		}
	}
	return shortest
}

func getDistanceFromStart(start [2]int, mapObstacles map[[2]int]bool) map[[2]int]int {
	unvisitedLocations := map[[2]int]bool{}
	mapDistance := map[[2]int]int{}
	infinity := 1000000000000
	min, max := getMinMax(mapObstacles)
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			if !mapObstacles[[2]int{i, j}] {
				unvisitedLocations[[2]int{i, j}] = true
				if [2]int{i, j} == start {
					mapDistance[[2]int{i, j}] = 0
				} else {
					mapDistance[[2]int{i, j}] = infinity
				}
			}
		}
	}

	for k := min; k <= len(mapDistance); k++ {
		currentPosition := getShortestDistance(unvisitedLocations, mapDistance, infinity)
		if len(unvisitedLocations) == 0 || mapDistance[currentPosition] == infinity {
			break
		}
		positionUp := [2]int{currentPosition[0] - 1, currentPosition[1]}
		if !mapObstacles[positionUp] && mapDistance[positionUp] > mapDistance[currentPosition]+1 {
			mapDistance[positionUp] = mapDistance[currentPosition] + 1
		}
		positionDown := [2]int{currentPosition[0] + 1, currentPosition[1]}
		if !mapObstacles[positionDown] && mapDistance[positionDown] > mapDistance[currentPosition]+1 {
			mapDistance[positionDown] = mapDistance[currentPosition] + 1
		}
		positionLeft := [2]int{currentPosition[0], currentPosition[1] - 1}
		if !mapObstacles[positionLeft] && mapDistance[positionLeft] > mapDistance[currentPosition]+1 {
			mapDistance[positionLeft] = mapDistance[currentPosition] + 1
		}
		positionRight := [2]int{currentPosition[0], currentPosition[1] + 1}
		if !mapObstacles[positionRight] && mapDistance[positionRight] > mapDistance[currentPosition]+1 {
			mapDistance[positionRight] = mapDistance[currentPosition] + 1
		}
		delete(unvisitedLocations, currentPosition)
	}
	return mapDistance
}

func distanceBetweenPos(pos1 [2]int, pos2 [2]int) int {
	distX := pos1[0] - pos2[0]
	if distX < 0 {
		distX = -1 * distX
	}
	distY := pos1[1] - pos2[1]
	if distY < 0 {
		distY = -1 * distY
	}
	return distX + distY
}

func findCheats(mapDistance map[[2]int]int, targetTimeSave int, cheatDistance int) int {
	nGoodCheats := 0
	for pos1, time1 := range mapDistance {
		for pos2, time2 := range mapDistance {
			if pos1 == pos2 {
				continue
			}
			if distanceBetweenPos(pos1, pos2) <= cheatDistance {
				savedTime := time1 - time2
				if savedTime < 0 {
					savedTime = savedTime * -1
				}
				savedTime = savedTime - distanceBetweenPos(pos1, pos2)
				if savedTime >= targetTimeSave {
					//fmt.Println(pos1, pos2)
					nGoodCheats++
				}
			}
		}
	}
	return nGoodCheats / 2
}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	start, end, mapObstacles := readGrid(inputFile)
	printGrid(start, end, mapObstacles)
	mapDistance := getDistanceFromStart(start, mapObstacles)
	distanceToEnd := mapDistance[end]
	fmt.Println(distanceToEnd)

	//nGoodCheats := findCheats(mapDistance, 12, 2)
	nGoodCheats := findCheats(mapDistance, 100, 2)
	fmt.Println("Number of 2-picosecond cheats saving 100 picoseconds (part 1): ", nGoodCheats)

	//nGoodCheats2 := findCheats(mapDistance, 72, 20)
	nGoodCheats2 := findCheats(mapDistance, 100, 20)
	fmt.Println("Number of 20-picosecond cheats saving 100 picoseconds (part 2): ", nGoodCheats2)
}
