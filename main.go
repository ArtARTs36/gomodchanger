package main

import (
	"context"

	"github.com/artarts36/gomodchanger/pkg/cmd"
	cli "github.com/artarts36/singlecli"
)

func main() {
	application := &cli.App{
		BuildInfo: &cli.BuildInfo{
			Name:      "gomodchanger",
			Version:   "v0.1.0",
			BuildDate: "23-11-2024 15:30",
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
			{
				Name:        "nested",
				Description: "replace also requirements to nested modules",
				WithValue:   false,
			},
		},
		Action: run,
	}

	application.RunWithGlobalArgs(context.Background())
}

func run(ctx *cli.Context) error {
	command := cmd.Default()

	projectDir := "./"
	if pd, ok := ctx.GetOpt("project-dir"); ok {
		projectDir = pd
	}

	return command.Run(ctx.Context, cmd.Params{
		NewModule:            ctx.Args["new-module"],
		ProjectDir:           projectDir,
		ReplaceNestedModules: ctx.HasOpt("nested"),
	})
}
