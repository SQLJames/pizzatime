package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"log"
	"io/ioutil"
)

type loljson struct {
	data data
	meta meta
}

type data struct {
	id 					string
	date 				string
	home_team 			teams
	home_team_score 	int
	period 				int
	postseason 			bool
	season 				int
	status 				string
	time 				string
	vistor_team			teams
	visitor_team_score 	int
}

type meta struct {
	total_pages 		int
	current_page 		int
	next_page			int
	per_page			int
	total_count			int

}

type teams struct{
	id 					int
	abbreviation 		string
	city				string
	conference			string
	division 			string
	full_name			string
	name 				string
}

func pizzaurl()(pizzaurl string) {
	year, month, day := time.Now().Date()
	return fmt.Sprint("https://www.balldontlie.io/api/v1/games?seasons[]=%s&team_ids[]=16&dates[]=%s", year, fmt.Sprint("%s-%s-%s", year, month, day))
}

func getresult()(result string){
	resp, err := http.Get(pizzaurl())
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	jr := &loljson{}
	json.Unmarshal(body, &jr)
}