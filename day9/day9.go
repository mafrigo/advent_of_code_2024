package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func extendRepresentation(line string) []string {
	extendedRepresentation := []string{}
	for ichar, char := range line {
		size, _ := strconv.Atoi(string(char))
		if ichar%2 == 0 {
			id := strconv.Itoa(ichar / 2)
			for i := 0; i < size; i++ {
				extendedRepresentation = append(extendedRepresentation, id)
			}
		} else {
			for i := 0; i < size; i++ {
				extendedRepresentation = append(extendedRepresentation, ".")
			}
		}
	}
	return extendedRepresentation
}

func getChecksum(representation []string) int {
	checksum := 0
	for ichar, char := range representation {
		if string(char) == "." {
			continue
		}
		number, _ := strconv.Atoi(string(char))
		checksum += number * ichar
	}
	return checksum
}

func reorder1(extendedRepresentation []string) []string {
	var finalRepresentation []string
	copy(finalRepresentation, extendedRepresentation)
	for ichar, char := range finalRepresentation {
		if string(char) == "." {
			for j := len(finalRepresentation) - 1; j >= 0; j-- {
				//fmt.Println(j, ichar, string(finalRepresentation[j]))
				if j <= ichar {
					break
				}
				if string(finalRepresentation[j]) != "." {
					fileId := string(finalRepresentation[j])
					finalRepresentation[ichar] = fileId
					finalRepresentation[j] = "."
					break
				}
			}
		}
	}
	return finalRepresentation
}

func reorder2(extendedRepresentation []string, line string) []string {
	finalRepresentation := extendedRepresentation
	lastFileId := "."
	for j := len(finalRepresentation) - 1; j >= 0; j-- {
		fileId := string(finalRepresentation[j])
		if fileId == "." || fileId == lastFileId {
			continue
		}

		blockLength := 0
		for i := j; i >= 0; i-- {
			if string(finalRepresentation[i]) != fileId {
				break
			} else {
				blockLength++
			}
		}

		for ichar, char := range finalRepresentation {
			if j <= ichar {
				break
			}
			if string(char) == "." {
				voidLength := 0
				for i := 0; i < len(finalRepresentation)-ichar; i++ {
					if string(finalRepresentation[ichar+i]) != "." {
						break
					} else {
						voidLength++
					}
				}
				if string(finalRepresentation[j]) != "." {
					//fmt.Println(blockLength, voidLength, finalRepresentation)
					if blockLength <= voidLength {
						for i := 0; i < blockLength; i++ {
							finalRepresentation[ichar+i] = fileId
							finalRepresentation[j-i] = "."
						}
						break
					}
				}
			}
		}
		lastFileId = fileId
	}
	return finalRepresentation
}

func main() {
	//inputFile := "inputtest2"
	inputFile := "input"
	line := readLines(inputFile)[0]
	fmt.Println(line)
	extendedRepresentation := extendRepresentation(line)
	fmt.Println(extendedRepresentation)

	//Problem 1
	finalRepresentation := reorder1(extendedRepresentation)
	fmt.Println(finalRepresentation)
	fmt.Println("checksum:")
	fmt.Println(getChecksum(finalRepresentation))

	//Problem 2
	finalRepresentation2 := reorder2(extendedRepresentation, line)
	fmt.Println(finalRepresentation2)
	fmt.Println("checksum 2:")
	fmt.Println(getChecksum(finalRepresentation2))
}
