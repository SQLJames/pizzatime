package core

import (
	"context"
	"encoding/json"
	"fmt"
	balldontlie "pizzatime/balldontlie"
	"strconv"
	"strings"
	"time"
)

type games struct {
	config *Config
	runID  string
	TeamID int
}

func Games(c Config, team_id int) Action {
	return &games{
		config: &c,
		TeamID: team_id,
	}
}

func (r *games) Config() Config {
	fmt.Println(r.config)
	return *r.config
}
func (g *games) Run(ctx context.Context) error {
	url_params := make(map[string][]string, 0)
	var teams []string
	g.config.Logger.Debug(g.TeamID)
	if g.TeamID != -1 {
		teams = append(teams, strconv.Itoa(g.TeamID))
		url_params["team_ids[]"] = teams
	}
	var pages []string

	pages = append(pages, strconv.Itoa(100))
	url_params["per_page"] = pages
	var startdate []string
	date := time.Now().AddDate(0, 0, -7)
	date.Format("2006/01/02")
	datestring := strings.Split(date.String(), " ")
	startdate = append(startdate, datestring[0])
	g.config.Logger.Debug(startdate)
	url_params["start_date"] = startdate
	NBA := balldontlie.New(g.config.Logger)
	bytes := NBA.Get("https://www.balldontlie.io/api/v1/games", url_params)
	g.config.Logger.Debug(string(bytes))
	Games := &balldontlie.Games{}
	err := json.Unmarshal(bytes, Games)
	if err != nil {
		return err
	}

	Games.Eval()
	g.config.Logger.Debug(Games.Data)
	
	Games.Table()
	
	return nil
}
