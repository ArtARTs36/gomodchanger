package cmd

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/artarts36/gomodfinder"

	"github.com/artarts36/gomodchanger/internal/file"
	"github.com/artarts36/gomodchanger/internal/replacer"
)

type Command struct {
	importsReplacer ImportsReplacer
	modReplacer     ModReplacer
}

type Params struct {
	NewModule string

	ProjectDir string
}

type ImportsReplacer func(goFile *file.File, oldModule, newModule string) error
type ModReplacer func(modFile *gomodfinder.ModFile, newModule string) error

func NewCommand(
	importsReplacer ImportsReplacer,
	modReplacer ModReplacer,
) *Command {
	return &Command{
		importsReplacer: importsReplacer,
		modReplacer:     modReplacer,
	}
}

func Default() *Command {
	return NewCommand(replacer.ReplaceImports, replacer.ReplaceModule)
}

func (c *Command) Run(ctx context.Context, params Params) error {
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

		err = c.importsReplacer(f, gomod.Module.Mod.Path, params.NewModule)
		if err != nil {
			return fmt.Errorf("failed to replace go module in file %q: %w", f.Path, err)
		}
	}

	slog.InfoContext(ctx, "[cmd] changing go.mod")
	err = c.modReplacer(gomod, params.NewModule)
	if err != nil {
		return fmt.Errorf("failed to replace go.mod: %w", err)
	}

	return nil
}
