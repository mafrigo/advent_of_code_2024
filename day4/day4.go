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

func search_horizontal_xmas(lines []string) int {
	xmas_counter := 0
	for iline, line := range lines {
		for ichar, _ := range line {
			if ichar < 3 {
				continue
			}

			horizontal_word := string(lines[iline][ichar]) + string(lines[iline][ichar-1]) + string(lines[iline][ichar-2]) + string(lines[iline][ichar-3])
			if horizontal_word == "XMAS" || horizontal_word == "SAMX" {
				xmas_counter = xmas_counter + 1
			}
		}
	}
	return xmas_counter
}

func search_vertical_xmas(lines []string) int {
	xmas_counter := 0
	for iline, line := range lines {
		for ichar, _ := range line {
			if iline < 3 {
				continue
			}
			vertical_word := string(lines[iline][ichar]) + string(lines[iline-1][ichar]) + string(lines[iline-2][ichar]) + string(lines[iline-3][ichar])
			if vertical_word == "XMAS" || vertical_word == "SAMX" {
				xmas_counter = xmas_counter + 1
			}
		}
	}
	return xmas_counter
}

func search_diagonal_upright_xmas(lines []string) int {
	xmas_counter := 0
	for iline, line := range lines {
		for ichar, _ := range line {
			if ichar < 3 || iline > len(lines)-4 {
				continue
			}
			diagonal_word := string(lines[iline][ichar]) + string(lines[iline+1][ichar-1]) + string(lines[iline+2][ichar-2]) + string(lines[iline+3][ichar-3])
			if diagonal_word == "XMAS" || diagonal_word == "SAMX" {
				xmas_counter = xmas_counter + 1
			}
		}
	}
	return xmas_counter
}

func search_diagonal_downright_xmas(lines []string) int {
	xmas_counter := 0
	for iline, line := range lines {
		for ichar, _ := range line {
			if ichar > len(lines)-4 || iline > len(lines)-4 {
				continue
			}
			diagonal_word := string(lines[iline][ichar]) + string(lines[iline+1][ichar+1]) + string(lines[iline+2][ichar+2]) + string(lines[iline+3][ichar+3])
			if diagonal_word == "XMAS" || diagonal_word == "SAMX" {
				xmas_counter = xmas_counter + 1
			}
		}
	}
	return xmas_counter
}

func search_xmas_square(lines []string) int {
	xmas_counter := 0
	for iline, line := range lines {
		for ichar, _ := range line {
			if ichar > len(lines)-3 || iline > len(lines)-3 {
				continue
			}
			diagonal_word_upright := string(lines[iline][ichar+2]) + string(lines[iline+1][ichar+1]) + string(lines[iline+2][ichar])
			diagonal_word_downright := string(lines[iline][ichar]) + string(lines[iline+1][ichar+1]) + string(lines[iline+2][ichar+2])
			if (diagonal_word_downright == "MAS" || diagonal_word_downright == "SAM") && (diagonal_word_upright == "MAS" || diagonal_word_upright == "SAM") {
				xmas_counter = xmas_counter + 1
			}
		}
	}
	return xmas_counter
}

func main() {
	input_file := "input"
	//input_file := "inputtest"
	lines := read_lines(input_file)
	fmt.Println(lines)

	n_horizontal := search_horizontal_xmas(lines)
	fmt.Println(n_horizontal)
	n_vertical := search_vertical_xmas(lines)
	fmt.Println(n_vertical)
	n_diagonal_upright := search_diagonal_upright_xmas(lines)
	fmt.Println(n_diagonal_upright)
	n_diagonal_downright := search_diagonal_downright_xmas(lines)
	fmt.Println(n_diagonal_downright)
	fmt.Println("Number of XMAS")
	fmt.Println(n_horizontal + n_vertical + n_diagonal_upright + n_diagonal_downright)

	n_square := search_xmas_square(lines)
	fmt.Println("Number of XMAS squares")
	fmt.Println(n_square)
}
