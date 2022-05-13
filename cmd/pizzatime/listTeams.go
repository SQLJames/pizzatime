package main

import (
	"context"
	"log"
	"pizzatime/core"

	"github.com/spf13/cobra"
)

var (

	//TODO: toggle operations

	cmdListTeams = &cobra.Command{
		Use:   "teams",
		Short: "Lists the teams in the database",
		Args:  cobra.NoArgs,
		Run: func(_ *cobra.Command, args []string) {
			listTeams := core.Teams(
				core.Config{
					Logger: logger.WithField("command", "Teams"),
				},
			)
			if err := listTeams.Run(context.TODO()); err != nil {
				log.Fatal(err.Error())
			}
		},
	}
)

func init() {
	root.AddCommand(cmdListTeams)
}
