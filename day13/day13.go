package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

type machine struct {
	buttonA [2]int
	buttonB [2]int
	prize   [2]int
}

func readMachines(lines []string, part2 bool) []machine {
	machines := []machine{}
	var currentMachine machine
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line[0:8] == "Button A" {
			splitString := strings.Split(line, "+")
			yNum, _ := strconv.Atoi(splitString[2])
			splitString2 := strings.Split(splitString[1], ",")
			xNum, _ := strconv.Atoi(splitString2[0])
			currentMachine.buttonA = [2]int{xNum, yNum}
		} else if line[0:8] == "Button B" {
			splitString := strings.Split(line, "+")
			yNum, _ := strconv.Atoi(splitString[2])
			splitString2 := strings.Split(splitString[1], ",")
			xNum, _ := strconv.Atoi(splitString2[0])
			currentMachine.buttonB = [2]int{xNum, yNum}
		} else if line[0:8] == "Prize: X" {
			splitString := strings.Split(line, "=")
			yNum, _ := strconv.Atoi(splitString[2])
			splitString2 := strings.Split(splitString[1], ",")
			xNum, _ := strconv.Atoi(splitString2[0])
			if !part2 {
				currentMachine.prize = [2]int{xNum, yNum}
			} else {
				currentMachine.prize = [2]int{xNum + 10000000000000, yNum + 10000000000000}
			}
			machines = append(machines, currentMachine)
		}
	}
	return machines
}

func getMinTokens(current machine, maxPresses int) int {
	bPresses := int(math.Round((float64(current.prize[1]) - float64(current.prize[0]*current.buttonA[1])/float64(current.buttonA[0])) / (float64(current.buttonB[1]) - float64(current.buttonB[0]*current.buttonA[1])/float64(current.buttonA[0]))))
	aPresses := int(math.Round(float64(current.prize[0]-bPresses*current.buttonB[0]) / float64(current.buttonA[0])))
	if maxPresses != 0 && (maxPresses < aPresses || maxPresses < bPresses) {
		return 0
	} else if aPresses < 0 || bPresses < 0 {
		return 0
	} else if current.buttonA[0]*aPresses+current.buttonB[0]*bPresses != current.prize[0] || current.buttonA[1]*aPresses+current.buttonB[1]*bPresses != current.prize[1] {
		return 0
	} else {
		return 3*aPresses + bPresses
	}
}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	lines := readLines(inputFile)
	machines := readMachines(lines, false)
	nTokens := 0
	for _, currentMachine := range machines {
		nTokens += getMinTokens(currentMachine, 100)
	}
	fmt.Println("Minimum number of tokens ", nTokens)

	//part2
	machines = readMachines(lines, true)
	nTokens = 0
	for _, currentMachine := range machines {
		nTokens += getMinTokens(currentMachine, 0)
	}
	fmt.Println("Minimum number of tokens (part2) ", nTokens)
}
