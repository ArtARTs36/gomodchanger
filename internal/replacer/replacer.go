package replacer

import (
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/artarts36/gds"

	"github.com/artarts36/gomodchanger/internal/file"
)

type Replacer struct {
}

func NewReplacer() *Replacer {
	return &Replacer{}
}

func (r *Replacer) Replace(goFile *file.File, oldModule, newModule string) error {
	fset := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fset, goFile.Path, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	modified := false

	for i, goImport := range parsedFile.Imports {
		path := strings.Trim(goImport.Path.Value, `"`)
		if !strings.HasPrefix(path, oldModule) {
			continue
		}

		parsedFile.Imports[i].Path.Value = gds.
			NewString(newModule).
			Append(strings.TrimPrefix(path, oldModule)).
			Wrap(`"`).
			Value
		modified = true
	}

	if !modified {
		return nil
	}

	f, err := os.OpenFile(goFile.Path, os.O_RDWR, 0755)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", goFile.Path, err)
	}

	return format.Node(f, fset, parsedFile)
}
