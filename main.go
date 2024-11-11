package main

import (
	"context"
	"github.com/artarts36/gomodchanger/internal/app"
	"github.com/artarts36/gomodchanger/internal/replacer"
	cli "github.com/artarts36/singlecli"
)

func main() {
	application := &cli.App{
		BuildInfo: &cli.BuildInfo{
			Name: "gomodchanger",
		},
		Args: []*cli.ArgDefinition{
			{
				Name:        "new-module",
				Description: "new module",
				Required:    true,
			},
		},
		Opts: []*cli.OptDefinition{
			{
				Name:        "project-dir",
				Description: "path to project directory",
				WithValue:   true,
			},
		},
		Action: run,
	}

	application.RunWithGlobalArgs(context.Background())
}

func run(ctx *cli.Context) error {
	cmd := app.NewCommand(replacer.NewReplacer())

	projectDir := "./"
	if pd, ok := ctx.GetOpt("project-dir"); ok {
		projectDir = pd
	}

	return cmd.Run(ctx.Context, app.CommandRunParams{
		NewModule:  ctx.Args["new-module"],
		ProjectDir: projectDir,
	})
}
