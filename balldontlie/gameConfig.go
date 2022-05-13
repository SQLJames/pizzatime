package balldontlie

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

type Games struct {
	Data []Game `json:"data"`
	Meta any    `json:"meta"`
}

type Game struct {
	Id                 int       `json:"id"`
	Date               time.Time `json:"date"`
	Home_team_score    int       `json:"home_team_score"`
	Visitor_team_score int       `json:"visitor_team_score"`
	Season             int       `json:"season"`
	Period             int       `json:"period"`
	Status             string    `json:"status"`
	Time               string    `json:"time"`
	Postseason         bool      `json:"postseason"`
	Home_team          Team      `json:"home_team"`
	Visitor_team       Team      `json:"visitor_team"`
	Winner             Team      `json:"winner"`
	Pizzatime          string    `json:"orderPizza"`
}

func (g *Games) Table() {
	tw := tabwriter.NewWriter(os.Stdout, 10, 4, 4, ' ', tabwriter.TabIndent)
	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "Date", "Home_score", "Visitor_score", "Status", "Home_team", "Visitor_team", "Winner", "Get_Pizza?")

	for _, v := range g.Data {
		date := strings.Split(v.Date.String(), " ")
		fmt.Fprintf(tw, "%s\t%d\t%d\t%s\t%s\t%s\t%s\t%s\n", date[0], v.Home_team_score, v.Visitor_team_score, v.Status, v.Home_team.Name, v.Visitor_team.Name, v.Winner.Name, v.Pizzatime)

	}

	tw.Flush()
}

func (g *Games) Eval() {
	g.sort()
	g.determineWinner()
	g.orderPizza()
}

func (g *Games) sort() {
	sort.Slice(g.Data, func(i, j int) bool {
		return g.Data[i].Date.Before(g.Data[j].Date)
	})
}

func (g *Games) determineWinner() {
	for i, game := range g.Data {
		if game.Home_team_score > game.Visitor_team_score {
			g.Data[i].Winner = game.Home_team
		} else {
			g.Data[i].Winner = game.Visitor_team
		}
	}
}

func (g *Games) orderPizza() {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	for i, game := range g.Data {
		gameDate := strings.Split(game.Date.String(), " ")
		if gameDate[0] != yesterday {
			continue
		}
		if game.Winner.Name == "Heat" {
			g.Data[i].Pizzatime = "yes"
		}
	}
	for i, game := range g.Data {
		if game.Pizzatime == "" {
			g.Data[i].Pizzatime = "no"
		}
	}
}
