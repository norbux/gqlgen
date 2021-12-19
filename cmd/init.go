package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/norbux/gqlgen/api"
	"github.com/norbux/gqlgen/codegen/config"
	"github.com/norbux/gqlgen/internal/code"
	"github.com/norbux/gqlgen/plugin/servergen"
	"github.com/urfave/cli/v2"
)

// var configTemplate = template.Must(template.New("name").Parse(
var tmpl =	`# Where are all the schema files located? globs are supported eg  src/**/*.graphqls
schema:
  - [baseDirectory]graph/*.graphqls

# Where should the generated server code go?
exec:
  filename: [baseDirectory]graph/generated/generated.go
  package: generated

# Uncomment to enable federation
# federation:
#   filename: [baseDirectory]graph/generated/federation.go
#   package: generated

# Where should any generated models go?
model:
  filename: [baseDirectory]graph/model/models_gen.go
  package: model

# Where should the resolver implementations go?
resolver:
  layout: follow-schema
  dir: [baseDirectory]graph
  package: graph

# Optional: turn on use ` + "`" + `gqlgen:"fieldName"` + "`" + ` tags in your models
# struct_tag: json

# Optional: turn on to use []Thing instead of []*Thing
# omit_slice_element_pointers: false

# Optional: set to speed up generation time by not performing a final validation pass.
# skip_validation: true

# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.
autobind:
  - "{{.}}/[baseDirectory]graph/model"

# This section declares type mapping between the GraphQL and go type systems
#
# The first line in each type will be used as defaults for resolver arguments and
# modelgen, the others will be allowed when binding to fields. Configure them to
# your liking
models:
  ID:
    model:
      - github.com/norbux/gqlgen/graphql.ID
      - github.com/norbux/gqlgen/graphql.Int
      - github.com/norbux/gqlgen/graphql.Int64
      - github.com/norbux/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/norbux/gqlgen/graphql.Int
      - github.com/norbux/gqlgen/graphql.Int64
      - github.com/norbux/gqlgen/graphql.Int32
`

var configTemplate = template.Must(template.New("name").Parse(tmpl))

var schemaDefault = `# GraphQL schema example
#
# https://gqlgen.com/getting-started/

type Todo {
  id: ID!
  text: String!
  done: Boolean!
  user: User!
}

type User {
  id: ID!
  name: String!
}

type Query {
  todos: [Todo!]!
}

input NewTodo {
  text: String!
  userId: String!
}

type Mutation {
  createTodo(input: NewTodo!): Todo!
}
`

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "create a new gqlgen project",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
		&cli.StringFlag{Name: "config, c", Usage: "the config filename", Value: "gqlgen.yml"},
		&cli.StringFlag{Name: "server", Usage: "where to write the server stub to", Value: "server.go"},
		&cli.StringFlag{Name: "schema", Usage: "where to write the schema stub to", Value: "graph/schema.graphqls"},
		&cli.StringFlag{Name: "base-dir, bd", Usage: "Base directory for config file, server.go, schema and models", Value: ""},
	},
	Action: func(ctx *cli.Context) error {
		configFilename := ctx.String("config")
		serverFilename := ctx.String("server")
		schemaFilename := ctx.String("schema")
		baseDirectory  := ctx.String("base-dir")

		fmt.Printf("Base directory in init: %v", baseDirectory)

		tmpl = strings.Replace(tmpl, "[baseDirectory]", baseDirectory, 6)
		configTemplate.Parse(tmpl)

		pkgName := code.ImportPathForDir(".")
		if pkgName == "" {
			return fmt.Errorf("unable to determine import path for current directory, you probably need to run go mod init first")
		}

		// check schema and config don't already exist
		for _, filename := range []string{baseDirectory + configFilename, schemaFilename, serverFilename} {
			if fileExists(filename) {
				return fmt.Errorf("%s already exists", filename)
			}
		}
		_, err := config.LoadConfigFromDefaultLocations()
		if err == nil {
			return fmt.Errorf("gqlgen.yml already exists in a parent directory")
		}

		// create config
		fmt.Println("Creating", baseDirectory + configFilename)
		if err := initFile(baseDirectory + configFilename, executeConfigTemplate(pkgName)); err != nil {
			return err
		}

		// create schema
		fmt.Println("Creating", baseDirectory + schemaFilename)
		if err := initFile(baseDirectory + schemaFilename, schemaDefault); err != nil {
			return err
		}

		// create the package directory with a temporary file so that go recognises it as a package
		// and autobinding doesn't error out
		tmpPackageNameFile := baseDirectory + "graph/model/_tmp_gqlgen_init.go"
		if err := initFile(tmpPackageNameFile, "package model"); err != nil {
			return err
		}
		defer os.Remove(tmpPackageNameFile)

		var cfg *config.Config
		if cfg, err = config.LoadConfig(baseDirectory + configFilename); err != nil {
			panic(err)
		}

		fmt.Println("Creating", baseDirectory + serverFilename)
		fmt.Println("Generating...")
		if err := api.Generate(cfg, api.AddPlugin(servergen.New(baseDirectory + serverFilename))); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		fmt.Printf("\nExec \"go run ./%s\" to start GraphQL server\n", baseDirectory + serverFilename)
		return nil
	},
}

func executeConfigTemplate(pkgName string) string {
	var buf bytes.Buffer
	if err := configTemplate.Execute(&buf, pkgName); err != nil {
		panic(err)
	}

	return buf.String()
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, fs.ErrNotExist)
}

func initFile(filename, contents string) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		return fmt.Errorf("unable to create directory for file '%s': %w", filename, err)
	}
	if err := ioutil.WriteFile(filename, []byte(contents), 0o644); err != nil {
		return fmt.Errorf("unable to write file '%s': %w", filename, err)
	}

	return nil
}
