package core

import (
	"context"
	"encoding/json"
	"fmt"
	balldontlie "pizzatime/balldontlie"
	"strconv"
)

type teams struct {
	config *Config
	runID  string
}

func Teams(c Config) Action {
	return &teams{
		config: &c,
	}
}

func (y *teams) Config() Config {
	fmt.Println(y.config)
	return *y.config
}
func (r *teams) Run(ctx context.Context) error {
	url_params := make(map[string][]string, 0)
	var params []string

	params = append(params, strconv.Itoa(100))
	url_params["per_page"] = params

	NBA := balldontlie.New(r.config.Logger)
	bytes := NBA.Get("https://www.balldontlie.io/api/v1/teams", url_params)
	r.config.Logger.Debug(string(bytes))
	Teams := &balldontlie.Teams{}
	err := json.Unmarshal(bytes, Teams)
	if err != nil {
		r.config.Logger.Error(err)
		return err
	}
	Teams.Table()
	//Teams.TableOutput2()
	return nil
}
