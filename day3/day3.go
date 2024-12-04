package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func calculate_mul(lines []string, use_do_instructions bool) int {
	mul_sequence_index_start := -1
	mul_sequence_index_end := -1
	mul_numbers := [][2]int{}
	do_enabled := true
	for _, line := range lines {
		for ichar, char := range line {
			if use_do_instructions {
				if ichar > 6 && line[ichar-7:ichar] == "don't()" {
					do_enabled = false
				}
				if ichar > 3 && line[ichar-4:ichar] == "do()" {
					do_enabled = true
				}
			}
			if do_enabled {
				if ichar > 3 {
					if line[ichar-4:ichar] == "mul(" {
						mul_sequence_index_start = ichar
					}
				}
				if mul_sequence_index_start != -1 {
					if string(char) == ")" {
						mul_sequence_index_end = ichar
						mul_content := line[mul_sequence_index_start:mul_sequence_index_end]
						numbers := strings.Split(mul_content, ",")
						if len(numbers) == 2 {
							n1, _ := strconv.Atoi(numbers[0])
							n2, _ := strconv.Atoi(numbers[1])
							mul_numbers = append(mul_numbers, [2]int{n1, n2})
						}
						mul_sequence_index_start = -1
						mul_sequence_index_end = -1
					}
				}
			}
		}
	}

	result := 0
	for _, num := range mul_numbers {
		result = result + num[0]*num[1]
	}
	return result
}

func main() {
	//input_file := "inputtest2"
	input_file := "input"
	lines := read_lines(input_file)

	result1 := calculate_mul(lines, false)
	fmt.Println("Result of mul operations without do instructions:")
	fmt.Println(result1)

	result2 := calculate_mul(lines, true)
	fmt.Println("Result of mul operations with do instructions:")
	fmt.Println(result2)
}
