package main

import (
	"context"
	"fmt"
	"log"
	"pizzatime/core"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	teamidKey = "teamid"
)

var (
	cmdListGames = &cobra.Command{
		Use: "games",
		//ValidArgs: []string{"team_id"},
		Short: "Lists the Games in the database",
		Run: func(cmd *cobra.Command, args []string) {

			team_id, err := cmd.Flags().GetInt("team_id")
			if err != nil {
				fmt.Println(err)
			}

			cmdListGames := core.Games(
				core.Config{
					Logger: logger.WithField("command", "Games"),
				},
				team_id)

			if err := cmdListGames.Run(context.TODO()); err != nil {
				log.Fatal(err.Error())
			}
		},
	}
)

func init() {
	v := viper.New()
	v.SetEnvPrefix("PIZZA")
	v.BindEnv("TEAM_ID")
	v.SetDefault("TEAM_ID", -1)
	var team_id int = v.GetInt("TEAM_ID")
	cmdListGames.Flags().IntP("team_id", "t", team_id, "Specifies the Team Id to get in the api call")
	root.AddCommand(cmdListGames)
}
