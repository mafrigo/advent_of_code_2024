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

func getAllPlants(grid [][]string) map[string][][2]int {
	plantMap := map[string][][2]int{}
	for iline, line := range grid {
		for ichar, char := range line {
			_, plantAlreadyInMap := plantMap[string(char)]
			if plantAlreadyInMap {
				plantMap[string(char)] = append(plantMap[string(char)], [2]int{iline, ichar})
			} else {
				plantMap[string(char)] = [][2]int{{iline, ichar}}
			}
		}
	}
	return plantMap
}

func getNumNeighbours(currentPlant [2]int, positionList [][2]int) int {
	nNeighbours := 0
	for _, otherPlant := range positionList {
		if otherPlant[0] == currentPlant[0]-1 && otherPlant[1] == currentPlant[1] || otherPlant[0] == currentPlant[0]+1 && otherPlant[1] == currentPlant[1] || otherPlant[0] == currentPlant[0] && otherPlant[1] == currentPlant[1]-1 || otherPlant[0] == currentPlant[0] && otherPlant[1] == currentPlant[1]+1 {
			nNeighbours++
		}
	}
	return nNeighbours
}

func isPosInList(pos [2]int, posList [][2]int) bool {
	posFound := false
	for _, pos2 := range posList {
		if pos == pos2 {
			posFound = true
		}
	}
	return posFound
}

func getNumCorners(currentPlant [2]int, positionList [][2]int) int {
	nCorners := 0
	if !isPosInList([2]int{currentPlant[0] - 1, currentPlant[1]}, positionList) && !isPosInList([2]int{currentPlant[0], currentPlant[1] - 1}, positionList) {
		nCorners++
	}
	if !isPosInList([2]int{currentPlant[0] + 1, currentPlant[1]}, positionList) && !isPosInList([2]int{currentPlant[0], currentPlant[1] - 1}, positionList) {
		nCorners++
	}
	if !isPosInList([2]int{currentPlant[0] - 1, currentPlant[1]}, positionList) && !isPosInList([2]int{currentPlant[0], currentPlant[1] + 1}, positionList) {
		nCorners++
	}
	if !isPosInList([2]int{currentPlant[0] + 1, currentPlant[1]}, positionList) && !isPosInList([2]int{currentPlant[0], currentPlant[1] + 1}, positionList) {
		nCorners++
	}
	if !isPosInList([2]int{currentPlant[0] - 1, currentPlant[1] - 1}, positionList) && isPosInList([2]int{currentPlant[0] - 1, currentPlant[1]}, positionList) && isPosInList([2]int{currentPlant[0], currentPlant[1] - 1}, positionList) {
		nCorners++
	}
	if !isPosInList([2]int{currentPlant[0] + 1, currentPlant[1] - 1}, positionList) && isPosInList([2]int{currentPlant[0] + 1, currentPlant[1]}, positionList) && isPosInList([2]int{currentPlant[0], currentPlant[1] - 1}, positionList) {
		nCorners++
	}
	if !isPosInList([2]int{currentPlant[0] - 1, currentPlant[1] + 1}, positionList) && isPosInList([2]int{currentPlant[0] - 1, currentPlant[1]}, positionList) && isPosInList([2]int{currentPlant[0], currentPlant[1] + 1}, positionList) {
		nCorners++
	}
	if !isPosInList([2]int{currentPlant[0] + 1, currentPlant[1] + 1}, positionList) && isPosInList([2]int{currentPlant[0] + 1, currentPlant[1]}, positionList) && isPosInList([2]int{currentPlant[0], currentPlant[1] + 1}, positionList) {
		nCorners++
	}
	return nCorners
}

func getIndexByPosition(position [2]int, regionMap map[int][][2]int) int {
	for regionId, positions := range regionMap {
		for _, regPos := range positions {
			if position == regPos {
				return regionId
			}
		}
	}
	return -1
}

func getUniqueNeighbouringIndices(currentPlant [2]int, positionList [][2]int, regionMap map[int][][2]int) []int {
	neighbouringIndices := []int{}
	for _, otherPlant := range positionList {
		if otherPlant[0] == currentPlant[0]-1 && otherPlant[1] == currentPlant[1] || otherPlant[0] == currentPlant[0]+1 && otherPlant[1] == currentPlant[1] || otherPlant[0] == currentPlant[0] && otherPlant[1] == currentPlant[1]-1 || otherPlant[0] == currentPlant[0] && otherPlant[1] == currentPlant[1]+1 {
			index := getIndexByPosition(otherPlant, regionMap)
			if index != -1 {
				indexFound := false
				for _, uniqueIndex := range neighbouringIndices {
					if index == uniqueIndex {
						indexFound = true
					}
				}
				if !indexFound {
					neighbouringIndices = append(neighbouringIndices, index)
				}
			}
		}
	}
	return neighbouringIndices
}

func joinRegions(regionMap map[int][][2]int, regionIds []int, targetRegion int) map[int][][2]int {
	for regionId, positions := range regionMap {
		indexFound := false
		for _, index := range regionIds {
			if regionId == index {
				indexFound = true
			}
		}
		if indexFound {
			if regionId != targetRegion {
				regionMap[targetRegion] = append(regionMap[targetRegion], positions...)
				delete(regionMap, regionId)
			}
		}
	}
	return regionMap
}

func getRegions(positionList [][2]int) map[int][][2]int {
	regionMap := map[int][][2]int{}
	for _, currentPlant := range positionList {
		neighbouringRegions := getUniqueNeighbouringIndices(currentPlant, positionList, regionMap)
		if len(neighbouringRegions) == 0 { //no known neighbouring regions -> create a new one with current plant
			maxRegionIndex := 0
			for regionId, _ := range regionMap {
				if regionId > maxRegionIndex {
					maxRegionIndex = regionId
				}
			}
			regionMap[maxRegionIndex+1] = [][2]int{currentPlant}
		} else if len(neighbouringRegions) == 1 { //only one known neighbouring region -> add current plant to it
			regionMap[neighbouringRegions[0]] = append(regionMap[neighbouringRegions[0]], currentPlant)
		} else if len(neighbouringRegions) > 1 { //multiple known neighbouring regions -> join them and add new plant
			minRegionIndex := len(positionList)
			for _, index := range neighbouringRegions {
				if index < minRegionIndex {
					minRegionIndex = index
				}
			}
			regionMap = joinRegions(regionMap, neighbouringRegions, minRegionIndex)
			regionMap[minRegionIndex] = append(regionMap[minRegionIndex], currentPlant)
		}
	}
	return regionMap
}

func main() {
	//inputFile := "inputtest5"
	inputFile := "input"
	grid := readGrid(inputFile)
	plantMap := getAllPlants(grid)

	price1 := 0
	price2 := 0
	for _, positions := range plantMap {
		regionMap := getRegions(positions)
		for _, regPositions := range regionMap {
			perimeter := 0
			area := 0
			nCorners := 0 // the number of corners equals the number of sides and is easier to calculate
			for _, position := range regPositions {
				perimeter += (4 - getNumNeighbours(position, regPositions))
				area++
				nCorners += getNumCorners(position, regPositions)
			}
			price1 += perimeter * area
			price2 += nCorners * area
		}
	}
	fmt.Println("Total Price (part 1): ", price1)
	fmt.Println("Total Price (part 2): ", price2)
}
