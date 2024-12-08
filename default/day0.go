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

func main() {
	//inputFile := "inputtest2"
	inputFile := "input"
	lines := readLines(inputFile)
	fmt.Println(lines)
}
