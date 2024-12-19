package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getTowels(lines []string) ([]string, []string) {
	availableTowelsUntrimmed := strings.Split(lines[0], ",")
	availableTowels := []string{}
	for _, towel := range availableTowelsUntrimmed {
		availableTowels = append(availableTowels, strings.TrimSpace(towel))
	}
	targetTowels := lines[2:]
	return availableTowels, targetTowels
}

var possibleDesigns = make(map[string]int)

func numberOfPossibleDesigns(target string, availableTowels []string) int {
	number, designAlreadyAnalyzed := possibleDesigns[target]
	if designAlreadyAnalyzed {
		//fmt.Println("Using map")
		return number
	} else {
		possibleDesigns[target] = 0
		for _, towel := range availableTowels {
			//fmt.Println(towel, target)
			if towel == target {
				possibleDesigns[target] += 1
			} else if len(towel) <= len(target) && towel == target[:len(towel)] {
				possibleDesigns[target] += numberOfPossibleDesigns(target[len(towel):], availableTowels)
			} else {
				continue
			}
		}
		return possibleDesigns[target]
	}
}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	lines := readLines(inputFile)
	availableTowels, targetTowels := getTowels(lines)
	//fmt.Println(availableTowels, targetTowels)
	nPossibleDesigns := 0
	nPossibleCombos := 0
	for _, target := range targetTowels {
		if numberOfPossibleDesigns(target, availableTowels) > 0 {
			//fmt.Println(target, " is possible")
			nPossibleDesigns++
			nPossibleCombos += possibleDesigns[target]
		}
	}
	fmt.Println("Number of possible designs (part 1): ", nPossibleDesigns)
	fmt.Println("Number of ways to make up the patterns (part 2): ", nPossibleCombos)
}
