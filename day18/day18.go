package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readBytes(filename string) [][2]int {
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
	bytes := [][2]int{}
	for _, line := range lines {
		bytesString := strings.Split(line, ",")
		byte1, _ := strconv.Atoi(bytesString[0])
		byte2, _ := strconv.Atoi(bytesString[1])
		bytes = append(bytes, [2]int{byte1, byte2})
	}
	return bytes
}

func printGrid(bytes [][2]int) {
	_, max := getMinMax(bytes)
	for i := 0; i <= max; i++ {
		for j := 0; j <= max; j++ {
			if isPosInList([2]int{i, j}, bytes) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getMinMax(bytes [][2]int) (int, int) {
	max := 0
	min := 100 * len(bytes)
	for _, byte := range bytes {
		if byte[0] < min {
			min = byte[0]
		}
		if byte[1] < min {
			min = byte[1]
		}
		if byte[0] > max {
			max = byte[0]
		}
		if byte[1] > max {
			max = byte[1]
		}
	}
	return min, max
}

func isPosInList(pos [2]int, posList [][2]int) bool {
	found := false
	for _, pos2 := range posList {
		if pos == pos2 {
			found = true
			break
		}
	}
	return found
}

func getLocWithSmallestDistance(mapDistance map[[2]int]int, unvisitedLocs map[[2]int]bool) [2]int {
	minimumDist := 100000000000
	minimumDistLoc := [2]int{-1, -1}
	for loc, _ := range unvisitedLocs {
		dist, _ := mapDistance[loc]
		if dist < minimumDist {
			minimumDist = dist
			minimumDistLoc = loc
		}
	}
	return minimumDistLoc
}

func shortestDistance(bytes [][2]int, start [2]int) int {
	unvisitedLocs := map[[2]int]bool{}
	mapDistance := map[[2]int]int{}
	bigNumber := 100000000000
	_, max := getMinMax(bytes)
	for i := 0; i <= max; i++ {
		for j := 0; j <= max; j++ {
			if !isPosInList([2]int{i, j}, bytes) {
				unvisitedLocs[[2]int{i, j}] = true
			}
		}
	}
	for loc, _ := range unvisitedLocs {
		if loc == start {
			mapDistance[loc] = 0
		} else {
			mapDistance[loc] = bigNumber
		}
	}
	for k := 0; len(unvisitedLocs) > 0; k++ {
		if len(unvisitedLocs) == 0 {
			break
		}
		currentLoc := getLocWithSmallestDistance(mapDistance, unvisitedLocs)
		if currentLoc == [2]int{-1, -1} {
			break
		}
		if currentLoc[0]-1 >= 0 {
			neighbour := [2]int{currentLoc[0] - 1, currentLoc[1]}
			if mapDistance[neighbour] > mapDistance[currentLoc] {
				mapDistance[neighbour] = mapDistance[currentLoc] + 1
			}
		}
		if currentLoc[0]+1 <= max {
			neighbour := [2]int{currentLoc[0] + 1, currentLoc[1]}
			if mapDistance[neighbour] > mapDistance[currentLoc] {
				mapDistance[neighbour] = mapDistance[currentLoc] + 1
			}
		}
		if currentLoc[1]-1 >= 0 {
			neighbour := [2]int{currentLoc[0], currentLoc[1] - 1}
			if mapDistance[neighbour] > mapDistance[currentLoc] {
				mapDistance[neighbour] = mapDistance[currentLoc] + 1
			}
		}
		if currentLoc[1]+1 <= max {
			neighbour := [2]int{currentLoc[0], currentLoc[1] + 1}
			if mapDistance[neighbour] > mapDistance[currentLoc] {
				mapDistance[neighbour] = mapDistance[currentLoc] + 1
			}
		}
		delete(unvisitedLocs, currentLoc)
	}
	return mapDistance[[2]int{max, max}]
}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	var part1Limit int
	if inputFile == "input" {
		part1Limit = 1024
	} else {
		part1Limit = 12
	}
	bytes := readBytes(inputFile)
	pathLength := shortestDistance(bytes[:part1Limit], [2]int{0, 0})
	//printGrid(bytes[:part1Limit])
	fmt.Println("Shortest distance after ", part1Limit, " bytes (part1): ", pathLength)

	//part2 - done manually by testing values; faster than testing every value in a loop
	testI := part1Limit + 1917
	fmt.Println(shortestDistance(bytes[:testI], [2]int{0, 0}))
	fmt.Println("Byte that breaks the path (part2): ", bytes[:testI][len(bytes[:testI])-1])
}
