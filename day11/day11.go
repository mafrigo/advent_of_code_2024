package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func blink(stones []string) []string {
	newStones := []string{}
	for _, stone := range stones {
		if stone == "0" {
			newStones = append(newStones, "1")
		} else if len(stone)%2 == 0 {
			newStone1, _ := strconv.Atoi(stone[0 : len(stone)/2])
			newStone2, _ := strconv.Atoi(stone[len(stone)/2:])
			newStones = append(newStones, strconv.Itoa(newStone1))
			newStones = append(newStones, strconv.Itoa(newStone2))
		} else {
			stoneInt, _ := strconv.Atoi(stone)
			newStones = append(newStones, strconv.Itoa(stoneInt*2024))
		}
	}
	return newStones
}

func efficientBlink(stones []string, nIterations int) int {
	nStones := 0
	var newStones int
	for _, stone := range stones {
		value, stoneAlreadyEncountered := blinkMap[key{stone, nIterations}]
		if stoneAlreadyEncountered {
			nStones += value
		} else if stone == "0" {
			if nIterations >= 2 {
				newStones = efficientBlink([]string{"1"}, nIterations-1)
			} else {
				newStones = 1
			}
			blinkMap[key{stone, nIterations}] = newStones
			nStones += newStones
		} else if len(stone)%2 == 0 {
			if nIterations >= 2 {
				newStone1, _ := strconv.Atoi(stone[0 : len(stone)/2])
				newStone2, _ := strconv.Atoi(stone[len(stone)/2:])
				newStones = efficientBlink([]string{strconv.Itoa(newStone1), strconv.Itoa(newStone2)}, nIterations-1)
			} else {
				newStones = 2
			}
			blinkMap[key{stone, nIterations}] = newStones
			nStones += newStones
		} else {
			if nIterations >= 2 {
				stoneInt, _ := strconv.Atoi(stone)
				newStones = efficientBlink([]string{strconv.Itoa(stoneInt * 2024)}, nIterations-1)
			} else {
				newStones = 1
			}
			blinkMap[key{stone, nIterations}] = newStones
			nStones += newStones
		}
	}
	return nStones
}

type key struct {
	stone     string
	iteration int
}

var blinkMap = map[key]int{}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	line := readLines(inputFile)[0]
	stones := strings.Split(line, " ")
	fmt.Println(stones)

	nIterations := 75 //Part1: 25, Part2: 75
	fmt.Println("Number of stones after " + strconv.Itoa(nIterations) + " iterations:")
	fmt.Println(efficientBlink(stones, nIterations))
}
