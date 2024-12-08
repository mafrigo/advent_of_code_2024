package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func findAntennas(lines []string) map[string][][2]int {
	antennaMap := make(map[string][][2]int)

	for iline, line := range lines {
		for ichar, char := range line {
			if string(char) != "." {
				_, ok := antennaMap[string(char)]
				if ok {
					antennaMap[string(char)] = append(antennaMap[string(char)], [2]int{iline, ichar})
				} else {
					antennaMap[string(char)] = [][2]int{[2]int{iline, ichar}}
				}
			}
		}
	}
	return antennaMap
}

func addAntinodeIfUniqueAndValid(antinode [2]int, antinodeList [][2]int, mapLimits [2]int) [][2]int {
	antinodeFound := false
	for _, existingAntinode := range antinodeList {
		if existingAntinode == antinode {
			antinodeFound = true
		}
	}
	if !antinodeFound && antinode[0] >= 0 && antinode[0] < mapLimits[0] && antinode[1] >= 0 && antinode[1] < mapLimits[1] {
		antinodeList = append(antinodeList, antinode)
	}
	return antinodeList
}

func findAntinodes1(antennaMap map[string][][2]int, mapLimits [2]int) [][2]int {
	antinodePositions := [][2]int{}
	for frequency, antennaPositions := range antennaMap {
		fmt.Println("Analyzing frequency ", frequency)
		for _, pos1 := range antennaPositions {
			for _, pos2 := range antennaPositions {
				if pos1 == pos2 {
					continue
				}
				antinodeSeparation := [2]int{pos1[0] - pos2[0], pos1[1] - pos2[1]}
				newAntinode1 := [2]int{pos1[0] + antinodeSeparation[0], pos1[1] + antinodeSeparation[1]}
				antinodePositions = addAntinodeIfUniqueAndValid(newAntinode1, antinodePositions, mapLimits)
				newAntinode2 := [2]int{pos2[0] - antinodeSeparation[0], pos2[1] - antinodeSeparation[1]}
				antinodePositions = addAntinodeIfUniqueAndValid(newAntinode2, antinodePositions, mapLimits)
			}
		}
	}
	return antinodePositions
}

func findAntinodes2(antennaMap map[string][][2]int, mapLimits [2]int) [][2]int {
	antinodePositions := [][2]int{}
	for frequency, antennaPositions := range antennaMap {
		fmt.Println("Analyzing frequency ", frequency)
		for _, pos1 := range antennaPositions {
			for _, pos2 := range antennaPositions {
				if pos1 == pos2 {
					continue
				}
				antinodeSeparation := [2]int{pos1[0] - pos2[0], pos1[1] - pos2[1]}
				for i := 0; i < mapLimits[0]; i++ {
					if antinodeSeparation[1]*(pos1[0]-i)%antinodeSeparation[0] == 0 {
						newAntinode := [2]int{i, pos1[1] - antinodeSeparation[1]*(pos1[0]-i)/antinodeSeparation[0]}
						antinodePositions = addAntinodeIfUniqueAndValid(newAntinode, antinodePositions, mapLimits)
					}
				}
			}
		}
	}
	return antinodePositions
}

func printMap(lines []string, antinodePositions [][2]int) {
	for iline, line := range lines {
		for ichar, _ := range line {
			antinodeFound := false
			for _, antinode := range antinodePositions {
				if [2]int{iline, ichar} == antinode {
					antinodeFound = true
				}
			}
			if antinodeFound {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	lines := readLines(inputFile)
	fmt.Println(lines)
	antennaMap := findAntennas(lines)
	fmt.Println(antennaMap)

	// Part 1
	antinodePositions := findAntinodes1(antennaMap, [2]int{len(lines), len(lines[0])})
	printMap(lines, antinodePositions)
	fmt.Println("Number of antinode positions (part 1):")
	fmt.Println(len(antinodePositions))

	// Part 2
	antinodePositions = findAntinodes2(antennaMap, [2]int{len(lines), len(lines[0])})
	printMap(lines, antinodePositions)
	fmt.Println("Number of antinode positions (part 2):")
	fmt.Println(len(antinodePositions))
}
