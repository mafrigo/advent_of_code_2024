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

var mapNumericKeypad map[string][2]int
var mapDirectionKeypad map[string][2]int

func initKeypads() {
	mapNumericKeypad = map[string][2]int{
		"0": {3, 1},
		"A": {3, 2},
		"1": {2, 0},
		"2": {2, 1},
		"3": {2, 2},
		"4": {1, 0},
		"5": {1, 1},
		"6": {1, 2},
		"7": {0, 0},
		"8": {0, 1},
		"9": {0, 2},
	}
	mapDirectionKeypad = map[string][2]int{
		"<": {1, 0},
		"v": {1, 1},
		">": {1, 2},
		"^": {0, 1},
		"A": {0, 2},
	}
}

func orderDirections(vertical string, horizontal string) string {
	if len(vertical) == 0 {
		return horizontal
	}
	if len(horizontal) == 0 {
		return vertical
	}
	if string(vertical[0]) == "v" && string(horizontal[0]) == ">" {
		return vertical + horizontal
	} else if string(vertical[0]) == "v" && string(horizontal[0]) == "<" {
		return horizontal + vertical
	} else if string(vertical[0]) == "^" && string(horizontal[0]) == ">" {
		return vertical + horizontal
	} else if string(vertical[0]) == "^" && string(horizontal[0]) == "<" {
		return horizontal + vertical
	} else {
		return ""
	}
}

func getPathBetweenKeys(key1 string, key2 string, numericKeypad bool) string {
	keyList := ""
	verticalKeyList := ""
	horizontalKeyList := ""
	var distX int
	var distY int
	if numericKeypad {
		distY = mapNumericKeypad[key2][0] - mapNumericKeypad[key1][0]
		distX = mapNumericKeypad[key2][1] - mapNumericKeypad[key1][1]
	} else {
		distY = mapDirectionKeypad[key2][0] - mapDirectionKeypad[key1][0]
		distX = mapDirectionKeypad[key2][1] - mapDirectionKeypad[key1][1]
	}
	if distY > 0 {
		for i := 0; i < distY; i++ {
			verticalKeyList += "v"
		}
	}
	if distY < 0 {
		for i := 0; i > distY; i-- {
			verticalKeyList += "^"
		}
	}
	if distX > 0 {
		for i := 0; i < distX; i++ {
			horizontalKeyList += ">"
		}
	}
	if distX < 0 {
		for i := 0; i > distX; i-- {
			horizontalKeyList += "<"
		}
	}
	if numericKeypad {
		if mapNumericKeypad[key1][1] == 0 && mapNumericKeypad[key2][0] == 3 {
			keyList = horizontalKeyList + verticalKeyList
		} else if mapNumericKeypad[key2][1] == 0 && mapNumericKeypad[key1][0] == 3 {
			keyList = verticalKeyList + horizontalKeyList
		} else {
			keyList = orderDirections(verticalKeyList, horizontalKeyList)
		}
	} else {
		if mapDirectionKeypad[key1][1] == 0 && mapDirectionKeypad[key2][0] == 0 {
			keyList = horizontalKeyList + verticalKeyList
		} else if mapDirectionKeypad[key2][1] == 0 && mapDirectionKeypad[key1][0] == 0 {
			keyList = verticalKeyList + horizontalKeyList
		} else {
			keyList = orderDirections(verticalKeyList, horizontalKeyList)
		}
	}
	return keyList
}

func getKeypadLength(code string, numericKeypad bool, nSteps int) int {
	instructionsLength, found := mapKeypadSteps[codeSteps{code, nSteps}]
	if found {
		return instructionsLength
	}
	nInstructions := 0
	start := "A"
	for _, key := range code {
		instructions := getPathBetweenKeys(start, string(key), numericKeypad) + "A"
		start = string(key)
		if nSteps > 1 {
			nInstructions += getKeypadLength(string(instructions), false, nSteps-1)
		} else {
			nInstructions += len(instructions)
		}
	}
	mapKeypadSteps[codeSteps{code, nSteps}] = nInstructions
	return nInstructions
}

type codeSteps struct {
	code   string
	nSteps int
}

var mapKeypadSteps map[codeSteps]int

func getComplexity(codes []string, nIntermediateRobots int) int {
	sumComplexity := 0
	for _, code := range codes {
		lengthAfterRobots := getKeypadLength(code, true, nIntermediateRobots+1)
		number, _ := strconv.Atoi(code[:len(code)-1])
		sumComplexity += lengthAfterRobots * number
	}
	return sumComplexity
}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	codes := readLines(inputFile)
	fmt.Println(codes)
	initKeypads()
	mapKeypadSteps = make(map[codeSteps]int)
	fmt.Println("Sum of the code complexities with 2 robots (part 1) : ", getComplexity(codes, 2))
	fmt.Println("Sum of the code complexities with 25 robots (part 2) : ", getComplexity(codes, 25))
}
