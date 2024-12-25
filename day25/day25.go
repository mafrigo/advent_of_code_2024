package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readGrids(filename string) [][][]string {
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
	grids := [][][]string{}
	currentGrid := [][]string{}
	for _, line := range lines {
		if len(line) == 0 {
			grids = append(grids, currentGrid)
			currentGrid = [][]string{}
			continue
		}
		gridLine := []string{}
		for _, char := range line {
			gridLine = append(gridLine, string(char))
		}
		currentGrid = append(currentGrid, gridLine)
	}
	if len(currentGrid) > 0 {
		grids = append(grids, currentGrid)
	}
	return grids
}

func getKeysLocks(grids [][][]string) ([][5]int, [][5]int) {
	keys := [][5]int{}
	locks := [][5]int{}
	for _, grid := range grids {
		if grid[0][0] == "#" { //lock
			currentLock := [5]int{-1, -1, -1, -1, -1}
			for iline, line := range grid {
				for ichar, char := range line {
					if currentLock[ichar] == -1 && string(char) == "." {
						currentLock[ichar] = iline - 1
					}
				}
			}
			locks = append(locks, currentLock)
		} else if grid[0][0] == "." { // key
			currentKey := [5]int{-1, -1, -1, -1, -1}
			for iline, line := range grid {
				for ichar, char := range line {
					if currentKey[ichar] == -1 && string(char) == "#" {
						currentKey[ichar] = len(line) + 1 - iline
					}
				}
			}
			keys = append(keys, currentKey)
		}
	}
	return keys, locks
}

func testKeyLock(key [5]int, lock [5]int) bool {
	fits := true
	for i := 0; i < 5; i++ {
		if key[i] > 5-lock[i] {
			fits = false
		}
	}
	return fits
}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	grids := readGrids(inputFile)
	keys, locks := getKeysLocks(grids)
	numKeysThatFit := 0
	for _, lock := range locks {
		for _, key := range keys {
			if testKeyLock(key, lock) {
				numKeysThatFit++
			}
		}
	}
	fmt.Println("Number of key-lock combinations that fit (part 1): ", numKeysThatFit)
	fmt.Println("Deliver the chronicle (part 2).")
}
