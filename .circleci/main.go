package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/google/go-github/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	app := &cli.App{
		Name:     "prebuild-gitpod-templates",
		Compiled: time.Now(),
		Version:  "v0.0.0",
		HelpName: "prebuild-gitpod-templates",
		Usage:    "rebuilds prebuilds for all gitpod-io/*",
		Authors: []*cli.Author{
			{
				Name:  "Geoffrey Huntley",
				Email: "ghuntley@ghuntley.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Aliases:  []string{"t"},
				Usage:    "GitHub API token",
				EnvVars:  []string{"GITHUB_API_TOKEN"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			var token = c.String("token")

			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc := oauth2.NewClient(ctx, ts)
		
			client := github.NewClient(tc)
			opts := &github.RepositoryListByOrgOptions{Type: "public"}
		
			for {
				repos, response, err := client.Repositories.ListByOrg(context.Background(), "gitpod-io", opts)
				if err != nil {
					return err
				}
		
				log.Trace("GitHub Rate Limit", response.Rate)
		
				log.Trace("Current Page: ", opts.Page)
				opts.Page = response.NextPage
		
				log.Trace("Next Page: ", response.NextPage)
				if response.NextPage == 0 {
					log.Trace("Final page!")
		
					break
				}
		
				for i, repo := range repos {
					i = i + 1;

					var htmlUrl = *repo.HTMLURL;

					if strings.Contains(htmlUrl, "template-") {
						var prebuildUrl = "https://gitpod.io/#prebuild/" + htmlUrl;
						fmt.Println(prebuildUrl);
					}
				}
			}
		
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}