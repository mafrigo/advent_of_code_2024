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

func main() {
	inputFile := "inputtest"
	//inputFile := "input"
	lines := readLines(inputFile)
	fmt.Println(lines)
	grid := readGrid(inputFile)
	printGrid(grid)
}
