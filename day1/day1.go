package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
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

func get_locations_from_file(filename string) ([]int, []int) {
	lines := read_lines(filename)
	var locationId1 []int
	var locationId2 []int
	for _, line := range lines {
		splitted := strings.Split(line, "   ")
		n1, _ := strconv.Atoi(splitted[0])
		n2, _ := strconv.Atoi(splitted[1])
		locationId1 = append(locationId1, n1)
		locationId2 = append(locationId2, n2)
	}
	return locationId1, locationId2
}

func abs_value(value int) int {
	return max(value, -value)
}

func get_distance(loc1 []int, loc2 []int) int {
	slices.Sort(loc1)
	slices.Sort(loc2)
	distance := 0
	for index, _ := range loc1 {
		distance = distance + abs_value(loc2[index]-loc1[index])
	}
	return distance
}

func get_similarity(loc1 []int, loc2 []int) int {
	similarity := 0
	occurrence_map := make(map[int]int)
	for _, n1 := range loc1 {
		counter := 0
		occurrence, ok := occurrence_map[n1]
		if !ok {
			for _, n2 := range loc2 {
				if n2 == n1 {
					counter = counter + 1
				}
			}
			occurrence_map[n1] = counter
			occurrence = counter
		}
		similarity = similarity + n1*occurrence
	}
	return similarity
}

func main() {
	input_file := "inputtest"
	//input_file := "input"
	debug := false
	loc1, loc2 := get_locations_from_file(input_file)

	if debug {
		fmt.Println("Loc1:")
		fmt.Println(loc1)
		fmt.Println("Loc2:")
		fmt.Println(loc2)
	}

	//problem 1
	distance := get_distance(loc1, loc2)
	fmt.Println("Distance:")
	fmt.Println(distance)

	//problem 2
	similarity := get_similarity(loc1, loc2)
	fmt.Println("Similarity:")
	fmt.Println(similarity)
}
