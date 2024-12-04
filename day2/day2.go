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

func get_reports(filename string) [][]int {
	reports := [][]int{}
	reports_str := read_lines(filename)
	for _, report := range reports_str {
		levels := []int{}
		levels_str := strings.Split(report, " ")
		for _, level_str := range levels_str {
			level, _ := strconv.Atoi(level_str)
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}
	return reports
}

func check_descending(levels []int, can_dampener_be_used bool) bool {
	if len(levels) == 1 {
		return true
	}
	initial_value := levels[0]
	next_value := levels[1]
	safe_with_dampener := false
	safe_without_dampener := false
	if initial_value-next_value >= 1 && initial_value-next_value <= 3 {
		safe_without_dampener = check_descending(levels[1:], can_dampener_be_used)
	}
	if can_dampener_be_used {
		if len(levels) <= 2 {
			safe_with_dampener = true
		} else {
			next_next_value := levels[2]
			if initial_value-next_next_value >= 1 && initial_value-next_next_value <= 3 {
				safe_with_dampener = check_descending(levels[2:], false)
			}
		}
	}
	if safe_without_dampener || safe_with_dampener {
		return true
	} else {
		return false
	}
}

func check_ascending(levels []int, can_dampener_be_used bool) bool {
	if len(levels) == 1 {
		return true
	}
	initial_value := levels[0]
	next_value := levels[1]
	safe_with_dampener := false
	safe_without_dampener := false
	if next_value-initial_value >= 1 && next_value-initial_value <= 3 {
		safe_without_dampener = check_ascending(levels[1:], can_dampener_be_used)
	}
	if can_dampener_be_used {
		if len(levels) <= 2 {
			safe_with_dampener = true
		} else {
			next_next_value := levels[2]
			if next_next_value-initial_value >= 1 && next_next_value-initial_value <= 3 {
				safe_with_dampener = check_ascending(levels[2:], false)
			}
		}
	}
	if safe_without_dampener || safe_with_dampener {
		return true
	} else {
		return false
	}
}

func main() {
	//input_file := "inputtest"
	input_file := "input"
	debug := false
	reports := get_reports(input_file)

	if debug {
		fmt.Println("Input:")
		fmt.Println(reports)
	}

	safe_counter := 0
	for _, report := range reports {
		safe_descending := check_descending(report, false)
		safe_ascending := check_ascending(report, false)
		safe := safe_descending || safe_ascending
		if debug {
			fmt.Println("Report:")
			fmt.Println(report)
			fmt.Println("Safety without dampener:")
			fmt.Println(safe_descending, safe_ascending, safe)
		}
		if safe {
			safe_counter = safe_counter + 1
		}
	}

	fmt.Println("Number of safe records without dampener:")
	fmt.Println(safe_counter)

	safe_counter_damp := 0
	for _, report := range reports {
		safe_descending := check_descending(report, true)
		safe_ascending := check_ascending(report, true)
		safe_descending_without_first := check_descending(report[1:], false)
		safe_ascending_without_first := check_ascending(report[1:], false)
		safe := safe_descending || safe_ascending || safe_descending_without_first || safe_ascending_without_first
		if debug {
			fmt.Println("Report:")
			fmt.Println(report)
			fmt.Println("Safety with_dampener:")
			fmt.Println(safe_descending, safe_ascending, safe)
		}
		if safe {
			safe_counter_damp = safe_counter_damp + 1
		}
	}

	fmt.Println("Number of safe records with dampener:")
	fmt.Println(safe_counter_damp)
}
