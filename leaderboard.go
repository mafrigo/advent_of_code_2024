package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func getJsonData(filename string) map[string]any {
	var leaderboard map[string]any
	file, err := os.ReadFile("apidata.json")
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}
	errJson := json.Unmarshal(file, &leaderboard)
	if errJson != nil {
		log.Fatalf("Unable to marshal JSON due to %s", errJson)
	}
	return leaderboard
}

func printGeneralList(leaderboard map[string]any) {
	members := leaderboard["members"].(map[string]any)
	keys := make([]string, 0, len(members))
	for key := range members {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return members[keys[i]].(map[string]any)["local_score"].(float64) > members[keys[j]].(map[string]any)["local_score"].(float64)
	})
	fmt.Println("\nGeneral leaderboard ordered by score:")
	for i, key := range keys {
		var name string
		if members[key].(map[string]any)["name"] != nil {
			name = members[key].(map[string]any)["name"].(string)
		} else {
			name = key
		}
		score := int(members[key].(map[string]any)["local_score"].(float64))
		if score == 0 {
			break
		}
		fmt.Println(i+1, name, score)
	}
}

func getStarCompletionTime(leaderboard map[string]any, member string, starId int, partId int) float64 {
	members := leaderboard["members"].(map[string]any)
	dayList := members[member].(map[string]any)["completion_day_level"]
	if dayList == nil {
		return 0
	}
	partList := dayList.(map[string]any)[strconv.Itoa(starId)]
	if partList == nil {
		return 0
	}
	partData := partList.(map[string]any)[strconv.Itoa(partId)]
	if partData == nil {
		return 0
	}
	completionTime := partData.(map[string]any)["get_star_ts"].(float64)
	return completionTime
}

func printListForGivenStar(leaderboard map[string]any, starId int, partId int) {
	members := leaderboard["members"].(map[string]any)
	keys := make([]string, 0, len(members))
	for key := range members {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return getStarCompletionTime(leaderboard, keys[j], starId, partId) > getStarCompletionTime(leaderboard, keys[i], starId, partId)
	})
	counter := 0
	fmt.Println("\nLeaderboard for star ", starId, " part ", partId, " ordered by minutes it took:")
	for _, key := range keys {
		timeStamp := getStarCompletionTime(leaderboard, key, starId, partId)
		if timeStamp == 0 {
			continue
		}
		var name string
		if members[key].(map[string]any)["name"] != nil {
			name = members[key].(map[string]any)["name"].(string)
		} else {
			name = key
		}
		counter++
		timeInSeconds := timeStamp - leaderboard["day1_ts"].(float64) - (float64(starId)-1)*24*3600
		timeInMinutes := int(timeInSeconds / 60)
		fmt.Println(counter, name, timeInMinutes)
	}
}

func main() {
	leaderboard := getJsonData("apidata.json")
	fmt.Println("\nTotal number of members of the leaderboard: ", len(leaderboard["members"].(map[string]any)))
	printGeneralList(leaderboard)
	printListForGivenStar(leaderboard, 10, 1)
}
