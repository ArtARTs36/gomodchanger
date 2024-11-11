package app

import (
	"context"
	"fmt"
	"github.com/artarts36/gomodchanger/internal/file"
	"github.com/artarts36/gomodchanger/internal/replacer"
	"github.com/artarts36/gomodfinder"
	"log/slog"
)

type Command struct {
	replacer *replacer.Replacer
}

type CommandRunParams struct {
	NewModule string

	ProjectDir string
}

func NewCommand(replacer *replacer.Replacer) *Command {
	return &Command{
		replacer: replacer,
	}
}

func (c *Command) Run(ctx context.Context, params CommandRunParams) error {
	gomod, err := gomodfinder.Find(params.ProjectDir, 1)
	if err != nil {
		return fmt.Errorf("failed to find go.mod: %w", err)
	}

	slog.InfoContext(ctx, "[cmd] finding go files")

	files, err := file.Collect(params.ProjectDir)
	if err != nil {
		return fmt.Errorf("failed to collect go files in %q: %w", params.ProjectDir, err)
	}

	if len(files) == 0 {
		slog.InfoContext(ctx, "[cmd] go files not found")
		return nil
	}

	for _, f := range files {
		slog.InfoContext(ctx, "[cmd] changing module in file", slog.String("file", f.Path))

		err = c.replacer.Replace(f, gomod.Module.Mod.Path, params.NewModule)
		if err != nil {
			return fmt.Errorf("failed to replace go module in file %q: %w", f.Path, err)
		}
	}

	return nil
}
