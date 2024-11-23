package replacer

import (
	"fmt"
	"github.com/artarts36/gds"
	"os"
	"strings"

	"golang.org/x/mod/modfile"

	"github.com/artarts36/gomodfinder"
)

func ReplaceModule(goMod *gomodfinder.ModFile, newModule string, replaceNestedPackages bool) error {
	oldModule := goMod.Module.Mod.Path

	err := goMod.AddModuleStmt(newModule)
	if err != nil {
		return fmt.Errorf("failed to set new module: %w", err)
	}

	if replaceNestedPackages {
		err = ReplaceNestedPackages(goMod, oldModule, newModule)
		if err != nil {
			return fmt.Errorf("failed to replace nested packages: %w", err)
		}
	}

	newContent := modfile.Format(goMod.Syntax)

	return os.WriteFile(goMod.Path, newContent, 0755)
}

func ReplaceNestedPackages(goMod *gomodfinder.ModFile, oldModule string, newModule string) error {
	requirementsMap := make(map[string]string) // map[old]new

	for _, req := range goMod.Require {
		path := strings.Trim(req.Mod.Path, `"`)
		if !strings.HasPrefix(path, oldModule) {
			continue
		}

		nestedPkg := gds.NewString(newModule).Append(strings.TrimPrefix(path, oldModule)).Value

		err := goMod.AddRequire(nestedPkg, req.Mod.Version)
		if err != nil {
			return fmt.Errorf("failed to add nested package: %w", err)
		}

		err = goMod.DropRequire(path)
		if err != nil {
			return fmt.Errorf("failed to drop old nested package: %w", err)
		}

		requirementsMap[path] = nestedPkg
	}

	for _, rep := range goMod.Replace {
		if newReplace, ok := requirementsMap[rep.Old.Path]; ok {
			err := goMod.DropReplace(rep.Old.Path, rep.Old.Version)
			if err != nil {
				return fmt.Errorf("failed to drop old nested replace: %w", err)
			}

			err = goMod.AddReplace(newReplace, rep.Old.Version, rep.New.Path, rep.New.Version)
			if err != nil {
				return fmt.Errorf("failed to add new replace: %w", err)
			}
		}
	}

	return nil
}
