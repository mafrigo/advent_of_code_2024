package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) (map[string]int, [][]string) {
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
	secondPart := false
	state := map[string]int{}
	operations := [][]string{}
	for _, line := range lines {
		if len(line) == 0 {
			secondPart = true
		} else if !secondPart {
			splitResult := strings.Split(line, ": ")
			state[splitResult[0]], _ = strconv.Atoi(splitResult[1])
		} else {
			operations = append(operations, strings.Split(line, " "))
		}
	}
	return state, operations
}

func findClosestUnvisitedOperation(operations [][]string, unvisitedOperations map[[3]string]bool, operationDistance map[[3]string]int, unvisitedOperationsByOperand map[[3]string][2]bool) []string {
	closest := []string{"notfound", "", "", "", ""}
	operationDistance[[3]string{closest[0], closest[2], closest[4]}] = 10000000000000000
	for _, operation := range operations {
		currentOp := [3]string{operation[0], operation[2], operation[4]}
		if unvisitedOperations[currentOp] {
			if !unvisitedOperationsByOperand[currentOp][0] && !unvisitedOperationsByOperand[currentOp][1] && operationDistance[currentOp] < operationDistance[[3]string{closest[0], closest[2], closest[4]}] {
				closest = operation
			}
		}
	}
	return closest
}

func orderOperations(operations [][]string, state map[string]int) [][]string {
	infinity := 100000000000
	unvisitedOperations := map[[3]string]bool{}
	unvisitedOperationsByOperand := map[[3]string][2]bool{}
	operationDistance := map[[3]string]int{}
	for _, operation := range operations {
		unvisitedOperations[[3]string{operation[0], operation[2], operation[4]}] = true
		_, foundInState1 := state[operation[0]]
		_, foundInState2 := state[operation[2]]
		if foundInState1 && !foundInState2 {
			unvisitedOperationsByOperand[[3]string{operation[0], operation[2], operation[4]}] = [2]bool{false, true}
		}
		if !foundInState1 && foundInState2 {
			unvisitedOperationsByOperand[[3]string{operation[0], operation[2], operation[4]}] = [2]bool{true, false}
		}
		if foundInState1 && foundInState2 {
			operationDistance[[3]string{operation[0], operation[2], operation[4]}] = 0
			unvisitedOperationsByOperand[[3]string{operation[0], operation[2], operation[4]}] = [2]bool{false, false}
		} else {
			operationDistance[[3]string{operation[0], operation[2], operation[4]}] = infinity
			unvisitedOperationsByOperand[[3]string{operation[0], operation[2], operation[4]}] = [2]bool{true, true}
		}
	}
	for i := 0; i < len(operations); i++ {
		operation := findClosestUnvisitedOperation(operations, unvisitedOperations, operationDistance, unvisitedOperationsByOperand)
		if operation[0] == "notfound" {
			break
		}
		operand1 := operation[0]
		operand2 := operation[2]
		resultLocation := operation[4]
		currentOp := [3]string{operand1, operand2, resultLocation}
		delete(unvisitedOperations, currentOp)

		for operation2, _ := range operationDistance {
			if operation2[0] == resultLocation {
				unvisitedOperationsByOperand[operation2] = [2]bool{false, unvisitedOperationsByOperand[operation2][1]}
			}
			if operation2[1] == resultLocation {
				unvisitedOperationsByOperand[operation2] = [2]bool{unvisitedOperationsByOperand[operation2][0], false}
			}
			if operation2[0] == resultLocation || operation2[1] == resultLocation {
				if operationDistance[currentOp]+1 > operationDistance[operation2] || operationDistance[operation2] == infinity {
					operationDistance[operation2] = operationDistance[currentOp] + 1
				}
			}
		}
	}
	sort.Slice(operations, func(i, j int) bool {
		opI := [3]string{operations[i][0], operations[i][2], operations[i][4]}
		opJ := [3]string{operations[j][0], operations[j][2], operations[j][4]}
		return operationDistance[opI] < operationDistance[opJ]
	})
	return operations
}

func calculateOperation(state map[string]int, operation []string) map[string]int {
	operand1, found1 := state[operation[0]]
	operand2, found2 := state[operation[2]]
	if !found1 || !found2 {
		fmt.Println("Operand not found, something is wrong...", operation[0], found1, operation[2], found2)
	}
	action := operation[1]
	resultLocation := operation[4]
	var result int
	if action == "AND" {
		result = operand1 * operand2
	}
	if action == "OR" {
		if operand1 == 1 || operand2 == 1 {
			result = 1
		} else {
			result = 0
		}
	}
	if action == "XOR" {
		result = operand1 ^ operand2
	}
	state[resultLocation] = result
	return state
}

func getResult(state map[string]int, keyId string) int {
	zList := []string{}
	indexList := []int{}
	mapZ := map[int]string{}
	for key, value := range state {
		if string(key[0]) == keyId {
			zList = append(zList, strconv.Itoa(value))
			index, _ := strconv.Atoi(key[1:])
			indexList = append(indexList, index)
			mapZ[index] = zList[len(zList)-1]
		}
	}
	binaryString := ""
	for i := 0; i < len(mapZ); i++ {
		binaryString += mapZ[len(mapZ)-i-1]
	}
	i, _ := strconv.ParseInt(binaryString, 2, 64)
	return int(i)
}

func formatLevel(level int) string {

	var levelString string
	if level >= 10 {
		levelString = strconv.Itoa(level)
	} else {
		levelString = "0" + strconv.Itoa(level)
	}
	return levelString
}

func checkIfAddition(zValue string, operations [][]string) []string {
	currentLevel, _ := strconv.Atoi(zValue[1:])
	xCurrent := "x" + formatLevel(currentLevel)
	yCurrent := "y" + formatLevel(currentLevel)
	xPrevious := "x" + formatLevel(currentLevel-1)
	yPrevious := "y" + formatLevel(currentLevel-1)
	wrongOutputs := []string{}
	currentXor := []string{}
	previousXor := []string{}
	previousAnd := []string{}
	orOperation := []string{}
	xorOperation := []string{}
	andOperation := []string{}
	zOperation := []string{}
	debug := false
	for _, operation := range operations {
		if ((operation[0] == xCurrent && operation[2] == yCurrent) || (operation[0] == yCurrent && operation[2] == xCurrent)) && (operation[1] == "XOR") {
			currentXor = operation
			if debug {
				fmt.Println(currentXor, "currentXor")
			}
			for _, operation2 := range operations {
				if (operation2[0] == currentXor[4] || operation2[2] == currentXor[4]) && operation2[1] == "XOR" {
					xorOperation = operation2
					if debug {
						fmt.Println(xorOperation, "xorOperation")
					}
				}
			}
		}
		if ((operation[0] == xPrevious && operation[2] == yPrevious) || (operation[0] == yPrevious && operation[2] == xPrevious)) && (operation[1] == "XOR") {
			previousXor = operation
			if debug {
				fmt.Println(previousXor, "previousXor")
			}
			for _, operation2 := range operations {
				if (operation2[0] == previousXor[4] || operation2[2] == previousXor[4]) && operation2[1] == "AND" {
					andOperation = operation2
					if debug {
						fmt.Println(andOperation, "andOperation")
					}
				}
			}
		}
		if ((operation[0] == xPrevious && operation[2] == yPrevious) || (operation[0] == yPrevious && operation[2] == xPrevious)) && (operation[1] == "AND") {
			previousAnd = operation
			if debug {
				fmt.Println(previousAnd, "previousAnd")
			}

			for _, operation2 := range operations {
				if (operation2[0] == previousAnd[4] || operation2[2] == previousAnd[4]) && operation2[1] == "OR" {
					orOperation = operation2
					if debug {
						fmt.Println(orOperation, "orOperation")
					}
				}
			}
		}
		if operation[4] == zValue {
			zOperation = operation
			if debug {
				fmt.Println(zOperation, "zOperation")
			}
		}
	}
	if len(xorOperation) > 0 {
		if xorOperation[4] != zValue {
			wrongOutputs = append(wrongOutputs, []string{xorOperation[4], zValue}...)
		}
	} else {
		if currentXor[4] != zOperation[0] && orOperation[4] == zOperation[2] {
			wrongOutputs = append(wrongOutputs, currentXor[4], zOperation[0])
		} else if currentXor[4] == zOperation[0] && orOperation[4] != zOperation[2] {
			wrongOutputs = append(wrongOutputs, currentXor[4], zOperation[2])
		}
	}
	return wrongOutputs
}

func toOrderedString(myList []string) string {
	sort.Slice(myList, func(i, j int) bool {
		return int(myList[j][0])*10000+int(myList[j][1])*100+int(myList[j][2]) > int(myList[i][0])*10000+int(myList[i][1])*100+int(myList[i][2])
	})
	return strings.Join(myList, ",")
}

func main() {
	//inputFile := "inputtest4"
	inputFile := "input"
	state, operations := readFile(inputFile)
	operations = orderOperations(operations, state)
	for _, operation := range operations {
		state = calculateOperation(state, operation)
	}
	//fmt.Println(state)
	fmt.Println("Result of operations (part 1): ", getResult(state, "z"))
	x := getResult(state, "x")
	y := getResult(state, "y")
	expectedZ := x + y
	zBinary := strconv.FormatInt(int64(expectedZ), 2)
	wrongZValues := []string{}
	for i, digit := range zBinary {
		var zString string
		if len(zBinary)-i-1 >= 10 {
			zString = "z" + strconv.Itoa(len(zBinary)-i-1)
		} else {
			zString = "z0" + strconv.Itoa(len(zBinary)-i-1)
		}
		if string(digit) != strconv.Itoa(state[zString]) {
			wrongZValues = append(wrongZValues, zString)
		}
	}
	fmt.Println("Wrong z values: ", wrongZValues)
	wrongOutputs := []string{}
	for _, zValue := range wrongZValues {
		wrongOutputs = append(wrongOutputs, checkIfAddition(zValue, operations)...)

	}
	fmt.Println("Values that need to be ordered (part 2): ", toOrderedString(wrongOutputs))
	/*
		7: gmt z07
		11: cbj qjj
		18: dmn z18
		35: z35 cfk
	*/
}
