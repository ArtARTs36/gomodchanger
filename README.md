## gomodchanger

gomodchanger - console tool for changing go module path and imports in code files.

install as: `go install github.com/artarts36/gomodchanger`

```
Usage
  gomodchanger new-module [--project-dir=<value>]

Arguments
  new-module   new module, required

Options
  project-dir  path to project directory
```

## Usage example

/usr/project1/go.mod:
```
module /usr/project1
```

run `gomodchanger /usr/project2` or `gomodchanger /usr/project2 --project-dir=/usr/project1`





