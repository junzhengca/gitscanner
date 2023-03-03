package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Version: "1.0.0",
		Name:    "gitscanner",
		Usage:   "Scan public git repositories for vulnerabilities",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Aliases:  []string{"t"},
				Usage:    "GitHub token, you can generate one on GitHub settings page",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "org",
				Aliases:  []string{"o"},
				Usage:    "GitHub organization name",
				Required: true,
			},
		},
		Action: ScanGitHubOrganization,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
