package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(filename string) []string {
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

func printGrid(grid [][]string) {
	for _, line := range grid {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func getTrailheads(grid [][]string) [][2]int {
	trailheads := [][2]int{}
	for iline, line := range grid {
		for ichar, char := range line {
			if string(char) == "0" {
				trailheads = append(trailheads, [2]int{iline, ichar})
			}
		}
	}
	return trailheads
}

func uniqueAppend(positions1 [][2]int, positions2 [][2]int) [][2]int {
	for _, pos2 := range positions2 {
		found := false
		for _, pos1 := range positions1 {
			if pos1 == pos2 {
				found = true
			}
		}
		if !found {
			positions1 = append(positions1, pos2)
		}
	}
	return positions1
}

func pathSearch(grid [][]string, start [2]int) ([][2]int, int) {
	currentHeight, _ := strconv.Atoi(grid[start[0]][start[1]])
	if currentHeight == 9 {
		return [][2]int{start}, 1
	}
	listOfPeaks := [][2]int{}
	totalRating := 0
	if start[0] > 0 && grid[start[0]-1][start[1]] == strconv.Itoa(currentHeight+1) {
		peaks, rating := pathSearch(grid, [2]int{start[0] - 1, start[1]})
		totalRating += rating
		listOfPeaks = uniqueAppend(listOfPeaks, peaks)
	}
	if start[0] < len(grid)-1 && grid[start[0]+1][start[1]] == strconv.Itoa(currentHeight+1) {
		peaks, rating := pathSearch(grid, [2]int{start[0] + 1, start[1]})
		totalRating += rating
		listOfPeaks = uniqueAppend(listOfPeaks, peaks)
	}
	if start[1] > 0 && grid[start[0]][start[1]-1] == strconv.Itoa(currentHeight+1) {
		peaks, rating := pathSearch(grid, [2]int{start[0], start[1] - 1})
		totalRating += rating
		listOfPeaks = uniqueAppend(listOfPeaks, peaks)
	}
	if start[1] < len(grid[0])-1 && grid[start[0]][start[1]+1] == strconv.Itoa(currentHeight+1) {
		peaks, rating := pathSearch(grid, [2]int{start[0], start[1] + 1})
		totalRating += rating
		listOfPeaks = uniqueAppend(listOfPeaks, peaks)
	}
	return listOfPeaks, totalRating
}

func main() {
	//inputFile := "inputtest2"
	inputFile := "input"
	grid := readGrid(inputFile)
	printGrid(grid)
	trailheads := getTrailheads(grid)
	sumTrailheadScores := 0
	sumRatings := 0
	for _, trailhead := range trailheads {
		destinations, rating := pathSearch(grid, trailhead)
		trailheadScore := len(destinations)
		sumTrailheadScores += trailheadScore
		sumRatings += rating
	}
	fmt.Println("Sum of trailhead scores (part1):")
	fmt.Println(sumTrailheadScores)
	fmt.Println("Sum of trailhead ratings (part2):")
	fmt.Println(sumRatings)
}
