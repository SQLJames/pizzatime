package balldontlie

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Teams struct {
	Data []Team `json:"data"`
	Meta any    `json:"meta"`
}

//Team defines a team response.
type Team struct {
	Id           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
	Full_name    string `json:"full_name"`
	Name         string `json:"name"`
}

func (t *Teams) Table() {
	tw := tabwriter.NewWriter(os.Stdout, 4, 8, 1, '\t', 0)
	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "Id", "Abbreviation", "City", "Conference", "Division", "Full_name", "Name")
	for _, v := range t.Data {
		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\t%s\t%s\t%s\n", v.Id, v.Abbreviation, v.City, v.Conference, v.Division, v.Full_name, v.Name)
	}
	tw.Flush()
}
