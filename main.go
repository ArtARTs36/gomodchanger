package main

import (
	"context"
	"github.com/artarts36/gomodchanger/pkg/cmd"
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
	command := cmd.NewCommand()

	projectDir := "./"
	if pd, ok := ctx.GetOpt("project-dir"); ok {
		projectDir = pd
	}

	return command.Run(ctx.Context, cmd.Params{
		NewModule:  ctx.Args["new-module"],
		ProjectDir: projectDir,
	})
}
