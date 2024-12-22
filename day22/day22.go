package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readPrices(filename string) []int {
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
	prices := []int{}
	for _, line := range lines {
		price, _ := strconv.Atoi(line)
		prices = append(prices, price)
	}
	return prices
}

func getNextSecret(price int) int {
	price = price ^ (price*64)%16777216
	price = price ^ (price/32)%16777216
	price = price ^ (price*2048)%16777216
	return price
}

var priceMap map[[4]int]int

func main() {
	//inputFile := "inputtest2"
	inputFile := "input"
	prices := readPrices(inputFile)
	nIterations := 2000
	priceMap = make(map[[4]int]int)
	sumOfEndPrices := 0
	for _, value := range prices {
		price0 := 0
		price1 := 0
		price2 := 0
		price3 := 0
		price4 := 0
		priceMapBySeller := map[[4]int]int{}
		for i := 0; i < nIterations; i++ {
			value = getNextSecret(value)
			valueString := strconv.Itoa(value)
			price4 = price3
			price3 = price2
			price2 = price1
			price1 = price0
			price0, _ = strconv.Atoi(string(valueString[len(valueString)-1]))
			if i >= 4 {
				sequence := [4]int{price3 - price4, price2 - price3, price1 - price2, price0 - price1}
				_, ok := priceMapBySeller[sequence]
				if !ok {
					priceMapBySeller[sequence] += price0
				}
			}
		}
		for sequence, price := range priceMapBySeller {
			_, alreadyInGeneralMap := priceMap[sequence]
			if !alreadyInGeneralMap {
				priceMap[sequence] = price
			} else {
				priceMap[sequence] += price
			}
		}
		sumOfEndPrices += value
	}
	fmt.Println("Sum of the end prices (part 1): ", sumOfEndPrices)

	bestSequence := [4]int{0, 0, 0, 0}
	for sequence, price := range priceMap {
		if price > priceMap[bestSequence] {
			bestSequence = sequence
		}
	}
	fmt.Println("Best sequence is ", bestSequence, ", and it gives this many bananas (part 2): ", priceMap[bestSequence])
}
