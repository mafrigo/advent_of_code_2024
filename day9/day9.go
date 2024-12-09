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

func reorder1(representation []string) []string {
	for ichar, char := range representation {
		if string(char) == "." {
			for j := len(representation) - 1; j >= 0; j-- {
				//fmt.Println(j, ichar, string(representation[j]))
				if j <= ichar {
					break
				}
				if string(representation[j]) != "." {
					fileId := string(representation[j])
					representation[ichar] = fileId
					representation[j] = "."
					break
				}
			}
		}
	}
	return representation
}

func reorder2(representation []string) []string {
	lastFileId := "."
	for j := len(representation) - 1; j >= 0; j-- {
		fileId := string(representation[j])
		if fileId == "." || fileId == lastFileId {
			continue
		}

		blockLength := 0
		for i := j; i >= 0; i-- {
			if string(representation[i]) != fileId {
				break
			} else {
				blockLength++
			}
		}

		for ichar, char := range representation {
			if j <= ichar {
				break
			}
			if string(char) == "." {
				voidLength := 0
				for i := 0; i < len(representation)-ichar; i++ {
					if string(representation[ichar+i]) != "." {
						break
					} else {
						voidLength++
					}
				}
				if string(representation[j]) != "." {
					//fmt.Println(blockLength, voidLength, representation)
					if blockLength <= voidLength {
						for i := 0; i < blockLength; i++ {
							representation[ichar+i] = fileId
							representation[j-i] = "."
						}
						break
					}
				}
			}
		}
		lastFileId = fileId
	}
	return representation
}

func main() {
	inputFile := "inputtest"
	//inputFile := "input"
	line := readLines(inputFile)[0]

	//Problem 1
	extendedRepresentation := extendRepresentation(line)
	//fmt.Println(extendedRepresentation)
	finalRepresentation := reorder1(extendedRepresentation)
	//fmt.Println(finalRepresentation)
	fmt.Println("checksum:")
	fmt.Println(getChecksum(finalRepresentation))

	//Problem 2
	extendedRepresentation2 := extendRepresentation(line)
	finalRepresentation2 := reorder2(extendedRepresentation2)
	//fmt.Println(finalRepresentation2)
	fmt.Println("checksum 2:")
	fmt.Println(getChecksum(finalRepresentation2))
}
