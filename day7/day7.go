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

func parseNumbers(line string) (int, []int) {
	splitResult := strings.Split(line, ":")
	testValue, _ := strconv.Atoi(splitResult[0])
	numbersStr := strings.Split(strings.TrimSpace(splitResult[1]), " ")
	numbers := []int{}
	for _, nstr := range numbersStr {
		nint, _ := strconv.Atoi(nstr)
		numbers = append(numbers, nint)
	}
	return testValue, numbers
}

func evaluateNumbers(numbers []int, testValue int) bool {
	if len(numbers) == 2 {
		return numbers[0]+numbers[1] == testValue || numbers[0]*numbers[1] == testValue
	} else {
		return evaluateNumbers(append([]int{numbers[0] + numbers[1]}, numbers[2:]...), testValue) || evaluateNumbers(append([]int{numbers[0] * numbers[1]}, numbers[2:]...), testValue)
	}
}

func evaluateNumbersWithConcat(numbers []int, testValue int) bool {
	if len(numbers) == 2 {
		return numbers[0]+numbers[1] == testValue || numbers[0]*numbers[1] == testValue || strconv.Itoa(numbers[0])+strconv.Itoa(numbers[1]) == strconv.Itoa(testValue)
	} else {
		concatNumbers, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1]))
		return evaluateNumbersWithConcat(append([]int{numbers[0] + numbers[1]}, numbers[2:]...), testValue) || evaluateNumbersWithConcat(append([]int{numbers[0] * numbers[1]}, numbers[2:]...), testValue) || evaluateNumbersWithConcat(append([]int{concatNumbers}, numbers[2:]...), testValue)
	}
}

func main() {
	input_file := "inputtest"
	//input_file := "input"
	lines := read_lines(input_file)
	sumSuccessfulEquations := 0
	for _, line := range lines {
		testValue, numbers := parseNumbers(line)
		testValueReached := evaluateNumbers(numbers, testValue)
		if testValueReached {
			sumSuccessfulEquations = sumSuccessfulEquations + testValue
		}
	}
	fmt.Println("Sum of test values of successful equations:")
	fmt.Println(sumSuccessfulEquations)

	sumSuccessfulEquations = 0
	for _, line := range lines {
		testValue, numbers := parseNumbers(line)
		testValueReached := evaluateNumbersWithConcat(numbers, testValue)
		if testValueReached {
			sumSuccessfulEquations = sumSuccessfulEquations + testValue
		}
	}
	fmt.Println("Sum of test values of successful equations with concatenation:")
	fmt.Println(sumSuccessfulEquations)
}
