package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
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
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
