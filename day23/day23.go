package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func readConnections(filename string) [][]string {
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
	connections := [][]string{}
	for _, line := range lines {
		splitResult := strings.Split(line, "-")
		connections = append(connections, splitResult)
	}
	return connections
}

var mapConnections map[string][]string

func addConnectionToMap(source string, target string) {
	existingConnections, existingEntry := mapConnections[source]
	if !existingEntry {
		mapConnections[source] = []string{target}
	} else {
		mapConnections[source] = append(existingConnections, target)
	}
}

func writeMapConnections(connections [][]string) {
	mapConnections = make(map[string][]string)
	for _, connection := range connections {
		addConnectionToMap(connection[0], connection[1])
		addConnectionToMap(connection[1], connection[0])
	}
}

func isConnectedTo(source string, target string) bool {
	isConnected := false
	connectedComputers, foundInMap := mapConnections[source]
	if !foundInMap {
		return false
	}
	for _, computer := range connectedComputers {
		if computer == target {
			isConnected = true
			break
		}
	}
	return isConnected
}

func appendIfNew(listTriplets [][3]string, triplet [3]string) [][3]string {
	tripletIsInList := false
	for _, triplet2 := range listTriplets {
		if (triplet[0] == triplet2[0] && triplet[1] == triplet2[1] && triplet[2] == triplet2[2]) ||
			(triplet[0] == triplet2[1] && triplet[1] == triplet2[0] && triplet[2] == triplet2[2]) ||
			(triplet[0] == triplet2[2] && triplet[1] == triplet2[1] && triplet[2] == triplet2[0]) ||
			(triplet[0] == triplet2[0] && triplet[1] == triplet2[2] && triplet[2] == triplet2[1]) ||
			(triplet[0] == triplet2[1] && triplet[1] == triplet2[2] && triplet[2] == triplet2[0]) ||
			(triplet[0] == triplet2[2] && triplet[1] == triplet2[0] && triplet[2] == triplet2[1]) {
			tripletIsInList = true
		}
	}
	//fmt.Println(triplet, listTriplets, tripletIsInList)
	if !tripletIsInList {
		listTriplets = append(listTriplets, triplet)
	}
	return listTriplets
}

func findTriplets(startLetter string) int {
	listTriplets := [][3]string{}
	for computer, connectedComputers := range mapConnections {
		if string(computer[0]) == startLetter {
			for _, computer2 := range connectedComputers {
				connectedComputers2 := mapConnections[computer2]
				for _, computer3 := range connectedComputers2 {
					if isConnectedTo(computer, computer3) {
						listTriplets = appendIfNew(listTriplets, [3]string{computer, computer2, computer3})
					}
				}
			}
		}
	}
	//fmt.Println(listTriplets)
	return len(listTriplets)
}

func findConnectedClusters(connections [][]string) []map[string]bool {
	connectedClusters := []map[string]bool{}
	for _, connection := range connections {
		connectedClusters = append(connectedClusters, map[string]bool{connection[0]: true, connection[1]: true})
	}

	for computer, _ := range mapConnections {
		for _, cluster := range connectedClusters {
			connectedToEveryComputer := true
			alreadyInCluster := false
			for clusterComputer, _ := range cluster {
				if !isConnectedTo(computer, clusterComputer) {
					connectedToEveryComputer = false
				}
				if computer == clusterComputer {
					alreadyInCluster = true
				}
			}
			if connectedToEveryComputer && !alreadyInCluster {
				cluster[computer] = true
			}
		}
	}
	return connectedClusters
}

func getLargestCluster(connectedClusters []map[string]bool) []string {
	largest := connectedClusters[0]
	for _, cluster := range connectedClusters {
		if len(cluster) > len(largest) {
			largest = cluster
		}
	}
	largestList := []string{}
	for computer, _ := range largest {
		largestList = append(largestList, computer)
	}
	return largestList
}

func toOrderedString(computerList []string) string {
	sort.Slice(computerList, func(i, j int) bool {
		return int(computerList[j][0])*100+int(computerList[j][1]) > int(computerList[i][0])*100+int(computerList[i][1])
	})
	return strings.Join(computerList, ",")
}

func main() {
	//inputFile := "inputtest"
	inputFile := "input"
	connections := readConnections(inputFile)
	writeMapConnections(connections) //saved into mapConnections
	nTriplets := findTriplets("t")
	fmt.Println("Number of triple connections with one computer starting with t (part 1) : ", nTriplets)
	connectedClusters := findConnectedClusters(connections)
	largestCluster := getLargestCluster(connectedClusters)
	fmt.Println("Largest interconnected group (part 2): ", toOrderedString(largestCluster))
}
