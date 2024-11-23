## gomodchanger

gomodchanger - console tool for changing go module path and imports in code files.

Install as: `go install github.com/artarts36/gomodchanger@v0.1.0`

```
Usage
  gomodchanger new-module [--project-dir=<value>] [--nested]

Arguments
  new-module   new module, required

Options
  project-dir  path to project directory
  nested       replace also requirements to nested modules
```

## Usage example

/usr/project1/go.mod:
```
module /usr/project1
```

run `gomodchanger /usr/project2` or `gomodchanger /usr/project2 --project-dir=/usr/project1`
