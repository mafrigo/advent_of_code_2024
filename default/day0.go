package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_lines(filename string) []string {
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
	//input_file := "inputtest2"
	input_file := "input"
	lines := read_lines(input_file)
	fmt.Println(lines)
}
