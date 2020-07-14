package main

import (
	"log"
	"os"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/battlemidget/columbo-go/internal/rules"
)

func main() {
	app := &cli.App{
		Name:  "columbo",
		Usage: "he got clues for dayzzz",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "rules",
				Aliases: []string{"r"},
				Usage:   "Load rules spec from `YAML` file",
			},
		},
		Action: func(c *cli.Context) error {
			var rulesSpec rules.RulesSpec
			rules := c.String("rules")
			output := rulesSpec.Parse(rules)
			fmt.Printf(output.Rules[0].Id)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
