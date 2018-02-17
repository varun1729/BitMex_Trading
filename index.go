package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type LeaderboardResponse struct {
	Profit float64 `json:"profit"`
	IsRealName bool `json:"isRealName"`
	Name string `json:"name"`
}

func main() {
	response, errorResponse := http.Get("https://www.bitmex.com/api/v1/leaderboard")
	if response.Status == "200 OK"{
		if errorResponse != nil {
		fmt.Printf("The HTTP request failed with error %s\n",errorResponse)
		} else {
			body, errorBody := ioutil.ReadAll(response.Body)
			if errorBody != nil {
				fmt.Println("There was an error retreiving the body", errorBody)
			} else {
				leaderboard := []LeaderboardResponse{}
				json.Unmarshal([]byte(body),&leaderboard)
				if leaderboard != nil {
					for i:= 0; i < len(leaderboard); i++ {
						fmt.Printf("The user %s made %f\n", leaderboard[i].Name, leaderboard[i].Profit)
					}
				} else {
					fmt.Println("leaderboard array is undefined")
				}
				defer response.Body.Close()
			}
		}
	} else {
		fmt.Println("Response recieved with status ", string(response.Status))
	}
}