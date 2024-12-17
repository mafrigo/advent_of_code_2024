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

func parseProgram(lines []string) ([3]int, []int) {
	var registers [3]int
	var program []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[0:11] == "Register A:" {
			registers[0], _ = strconv.Atoi(line[12:])
		}
		if line[0:11] == "Register B:" {
			registers[1], _ = strconv.Atoi(line[12:])
		}
		if line[0:11] == "Register C:" {
			registers[2], _ = strconv.Atoi(line[12:])
		}
		if line[0:8] == "Program:" {
			programOneString := line[9:]
			programManyStrings := strings.Split(programOneString, ",")
			for _, n := range programManyStrings {
				nInt, _ := strconv.Atoi(n)
				program = append(program, nInt)
			}
		}
	}
	return registers, program
}

func getComboOperand(registers [3]int, operand int) int {
	var comboOperand int
	if operand < 4 {
		comboOperand = operand
	} else if operand < 7 {
		comboOperand = registers[operand-4]
	} else {
		fmt.Println("combo operand 7 should not be reached")
		comboOperand = 8
	}
	return comboOperand
}

func adv(registers [3]int, operand int) [3]int {
	comboOperand := getComboOperand(registers, operand)
	registers[0] = registers[0] >> comboOperand // bit shift: equivalent to dividing by 2^n
	return registers
}

func bst(registers [3]int, operand int) [3]int {
	comboOperand := getComboOperand(registers, operand)
	registers[1] = comboOperand % 8
	return registers
}

func bxl(registers [3]int, operand int) [3]int {
	registers[1] = registers[1] ^ operand
	return registers
}

func jnz(registers [3]int, operand int) ([3]int, int) {
	if registers[0] == 0 {
		return registers, -1
	} else {
		return registers, operand
	}
}

func bxc(registers [3]int, operand int) [3]int {
	registers[1] = registers[1] ^ registers[2]
	return registers
}

func bdv(registers [3]int, operand int) [3]int {
	comboOperand := getComboOperand(registers, operand)
	registers[1] = registers[0] >> comboOperand // bit shift: equivalent to dividing by 2^n
	return registers
}

func cdv(registers [3]int, operand int) [3]int {
	comboOperand := getComboOperand(registers, operand)
	registers[2] = registers[0] >> comboOperand // bit shift: equivalent to dividing by 2^n
	return registers
}

func out(registers [3]int, operand int) int {
	comboOperand := getComboOperand(registers, operand)
	output := comboOperand % 8
	return output
}

func runInstructions(registers [3]int, program []int) []int {
	output := []int{}
	for instructionPointer := 0; instructionPointer < len(program); {
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]
		nextInstruction := instructionPointer + 2
		if instructionPointer >= len(program) {
			break
		}

		if opcode == 0 {
			registers = adv(registers, operand)
		}
		if opcode == 1 {
			registers = bxl(registers, operand)
		}
		if opcode == 2 {
			registers = bst(registers, operand)
		}
		if opcode == 3 {
			jump := 0
			registers, jump = jnz(registers, operand)
			if jump != -1 {
				nextInstruction = jump
			}
		}
		if opcode == 4 {
			registers = bxc(registers, operand)
		}
		if opcode == 5 {
			output = append(output, out(registers, operand))
		}
		if opcode == 6 {
			registers = bdv(registers, operand)
		}
		if opcode == 7 {
			registers = cdv(registers, operand)
		}

		instructionPointer = nextInstruction
	}
	return output
}

func joinOutput(output []int) string {
	outputString := ""
	for _, out := range output {
		if outputString == "" {
			outputString += strconv.Itoa(out)
		} else {
			outputString += "," + strconv.Itoa(out)
		}
	}
	return outputString
}

func isSameProgram(program1 []int, program2 []int) bool {
	if len(program1) != len(program2) {
		return false
	}
	isSame := true
	for i := 0; i < len(program1); i++ {
		if program1[i] != program2[i] {
			isSame = false
		}
	}
	return isSame
}

func main() {
	//inputFile := "inputtest4"
	inputFile := "input"
	lines := readLines(inputFile)
	registers, program := parseProgram(lines)
	//fmt.Println(registers, program)

	//Part 1
	output := runInstructions(registers, program)
	outputString := joinOutput(output)
	fmt.Println("Output (part1): ", outputString)

	//Part 2 - only works for real input (not for the tests)
	coeffMap := map[int][]int{}
	for i := len(program) - 1; i >= 0; i-- {
		coeffFound := false
		for coeff := 0; coeff < 8; coeff++ {
			if i == len(program)-1 && coeff == 0 {
				continue
			}
			aValue := 0
			//fmt.Println(coeffMap)
			for j := len(program) - 1; j > i; j-- {
				aValue += coeffMap[j][0] * int(math.Pow(8, float64(j)))
			}
			correctionToA := int(math.Pow(8, float64(i)))
			output2 := runInstructions([3]int{aValue + coeff*correctionToA, registers[1], registers[2]}, program)
			//fmt.Println(output2, aValue+coeff*correctionToA, i, coeff)
			if isSameProgram(output2[i:], program[i:]) {
				_, ok := coeffMap[i]
				if !ok {
					coeffMap[i] = []int{coeff}
				} else {
					coeffMap[i] = append(coeffMap[i], coeff)
				}
				coeffFound = true
			}
		}

		if !coeffFound {
			var newi int
			for newi = i; newi < len(program); newi++ {
				if len(coeffMap[newi]) > 1 {
					coeffMap[newi] = coeffMap[newi][1:]
					break
				} else {
					delete(coeffMap, newi)
				}
			}
			i = newi //rewind to latest iteration with a choice
		}
	}
	//fmt.Println(coeffMap)
	aValue := 0
	for j := len(program) - 1; j >= 0; j-- {
		aValue += coeffMap[j][0] * int(math.Pow(8, float64(j)))
	}
	//output2a := runInstructions([3]int{aValue, registers[1], registers[2]}, program)
	//fmt.Println(output2a)
	//fmt.Println(program)
	fmt.Println("Value of register A that causes the program to recreate itself (part 2): ", aValue)
}
