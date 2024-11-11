package file

import (
	"fmt"
	"github.com/artarts36/gds"
	"os"
	"strings"
)

var collectSkipDirs = gds.NewSet[string](
	".git",
	"vendor",
)

func Collect(dir string) ([]*File, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(fmt.Sprintf("failed to read directory %q: %s", dir, err))
	}

	files := make([]*File, 0)

	for _, entry := range entries {
		path := fmt.Sprintf("%s%s%s", dir, string(os.PathSeparator), entry.Name())

		if entry.IsDir() {
			if collectSkipDirs.Has(entry.Name()) {
				continue
			}

			children, childrenErr := Collect(path)
			if childrenErr != nil {
				return nil, childrenErr
			}
			files = append(files, children...)

			continue
		}

		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".pb.go") {
			continue
		}

		files = append(files, &File{
			Name: entry.Name(),
			Path: path,
		})
	}

	return files, nil
}
