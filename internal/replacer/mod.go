package replacer

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"

	"github.com/artarts36/gomodfinder"
)

// @todo add replace requirements
func ReplaceModule(goMod *gomodfinder.ModFile, newModule string) error {
	err := goMod.AddModuleStmt(newModule)
	if err != nil {
		return fmt.Errorf("failed to set new module: %w", err)
	}

	newContent := modfile.Format(goMod.Syntax)

	return os.WriteFile(goMod.Path, newContent, 0755)
}
