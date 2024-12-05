package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines(filename string) ([]string, []string) {
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
	dividing_line := 0
	for iline, line := range lines {
		if line == "" {
			dividing_line = iline
		}
	}
	return lines[0:dividing_line], lines[dividing_line:len(lines)]
}

func getUpdateInt(updateString string) []int {
	updateSplit := strings.Split(updateString, ",")
	updateInt := []int{}
	for _, pageString := range updateSplit {
		page, _ := strconv.Atoi(pageString)
		updateInt = append(updateInt, page)
	}
	return updateInt
}

func checkUpdateOrder(update string, orderRules []string) bool {
	updateInt := getUpdateInt(update)
	updateInOrder := true
	for ipage1, page1 := range updateInt {
		if updateInOrder == false {
			break
		}
		for _, page2 := range updateInt[ipage1+1:] {
			if updateInOrder == false {
				break
			}
			for _, rule := range orderRules {
				pageNumbers := strings.Split(rule, "|")
				nBefore, _ := strconv.Atoi(pageNumbers[0])
				nAfter, _ := strconv.Atoi(pageNumbers[1])
				if nBefore == page1 && nAfter == page2 {
					break
				}
				if nBefore == page2 && nAfter == page1 {
					updateInOrder = false
					break
				}
			}
		}
	}
	return updateInOrder
}

func orderUpdate(update []int, orderRules []string) []int {
	newUpdate := []int{}
	page1 := 0
	for ipage, page2 := range update {
		if ipage == 0 {
			page1 = page2
			newUpdate = []int{page1}
			continue
		} else {
			page1 = newUpdate[len(newUpdate)-1]
		}
		for _, rule := range orderRules {
			pageNumbers := strings.Split(rule, "|")
			nBefore, _ := strconv.Atoi(pageNumbers[0])
			nAfter, _ := strconv.Atoi(pageNumbers[1])
			if nBefore == page1 && nAfter == page2 {
				if ipage == 0 {
					newUpdate = append(newUpdate, page1)
				}
				newUpdate = append(newUpdate, page2)
				break
			}
			if nBefore == page2 && nAfter == page1 {
				if len(newUpdate) > 0 {
					newUpdate = newUpdate[:len(newUpdate)-1]
				}
				newUpdate = append(newUpdate, page2)
				newUpdate = append(newUpdate, page1)
				break
			}
		}
		fmt.Println(newUpdate)
	}
	return newUpdate
}

func getMiddle(update string) int {
	updateInt := getUpdateInt(update)
	middleIndex := len(updateInt) / 2
	return updateInt[middleIndex]
}

func main() {
	//input_file := "inputtest"
	input_file := "input"
	orderRules, updates := readLines(input_file)

	fmt.Println("Order rules:")
	fmt.Println(orderRules)
	fmt.Println("List of updates:")
	fmt.Println(updates)

	fmt.Println("Checking update order:")
	sumOrdered := 0
	sumUnordered := 0
	for _, update := range updates {
		updateValid := checkUpdateOrder(update, orderRules)
		if updateValid {
			sumOrdered = sumOrdered + getMiddle(update)
		} else {
			orderedUpdate := orderUpdate(getUpdateInt(update), orderRules)
			for _, _ = range getUpdateInt(update) {
				orderedUpdate = orderUpdate(orderedUpdate, orderRules)
			} // very inefficient trick :(
			fmt.Println(orderedUpdate)
			sumUnordered = sumUnordered + orderedUpdate[len(orderedUpdate)/2]
		}
	}
	fmt.Println("Sum of middle values of ordered updates (problem 1):")
	fmt.Println(sumOrdered)
	fmt.Println("Sum of middle values of unordered updates (problem 2):")
	fmt.Println(sumUnordered)
}
