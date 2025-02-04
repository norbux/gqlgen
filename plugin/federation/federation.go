package federation

import (
	"fmt"
	"sort"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/norbux/gqlgen/codegen"
	"github.com/norbux/gqlgen/codegen/config"
	"github.com/norbux/gqlgen/codegen/templates"
	"github.com/norbux/gqlgen/plugin"
	"github.com/norbux/gqlgen/plugin/federation/fieldset"
)

type federation struct {
	Entities []*Entity
}

// New returns a federation plugin that injects
// federated directives and types into the schema
func New() plugin.Plugin {
	return &federation{}
}

// Name returns the plugin name
func (f *federation) Name() string {
	return "federation"
}

// MutateConfig mutates the configuration
func (f *federation) MutateConfig(cfg *config.Config) error {
	builtins := config.TypeMap{
		"_Service": {
			Model: config.StringList{
				"github.com/norbux/gqlgen/plugin/federation/fedruntime.Service",
			},
		},
		"_Entity": {
			Model: config.StringList{
				"github.com/norbux/gqlgen/plugin/federation/fedruntime.Entity",
			},
		},
		"Entity": {
			Model: config.StringList{
				"github.com/norbux/gqlgen/plugin/federation/fedruntime.Entity",
			},
		},
		"_Any": {
			Model: config.StringList{"github.com/norbux/gqlgen/graphql.Map"},
		},
	}
	for typeName, entry := range builtins {
		if cfg.Models.Exists(typeName) {
			return fmt.Errorf("%v already exists which must be reserved when Federation is enabled", typeName)
		}
		cfg.Models[typeName] = entry
	}
	cfg.Directives["external"] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives["requires"] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives["provides"] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives["key"] = config.DirectiveConfig{SkipRuntime: true}
	cfg.Directives["extends"] = config.DirectiveConfig{SkipRuntime: true}

	return nil
}

func (f *federation) InjectSourceEarly() *ast.Source {
	return &ast.Source{
		Name: "federation/directives.graphql",
		Input: `
scalar _Any
scalar _FieldSet

directive @external on FIELD_DEFINITION
directive @requires(fields: _FieldSet!) on FIELD_DEFINITION
directive @provides(fields: _FieldSet!) on FIELD_DEFINITION
directive @key(fields: _FieldSet!) repeatable on OBJECT | INTERFACE
directive @extends on OBJECT | INTERFACE
`,
		BuiltIn: true,
	}
}

// InjectSources creates a GraphQL Entity type with all
// the fields that had the @key directive
func (f *federation) InjectSourceLate(schema *ast.Schema) *ast.Source {
	f.setEntities(schema)

	entities := ""
	resolvers := ""
	entityResolverInputDefinitions := ""
	for i, e := range f.Entities {
		if i != 0 {
			entities += " | "
		}
		entities += e.Name

		if e.ResolverName != "" {
			if e.Multi {
				entityResolverInputDefinitions += "input " + e.InputType + " {\n"
				for _, keyField := range e.KeyFields {
					entityResolverInputDefinitions += fmt.Sprintf("\t%s: %s\n", keyField.Field.ToGo(), keyField.Definition.Type.String())
				}
				entityResolverInputDefinitions += "}\n"
				resolvers += fmt.Sprintf("\t%s(reps: [%s!]!): [%s]\n", e.ResolverName, e.InputType, e.Name)
			} else {
				resolverArgs := ""
				for _, keyField := range e.KeyFields {
					resolverArgs += fmt.Sprintf("%s: %s,", keyField.Field.ToGoPrivate(), keyField.Definition.Type.String())
				}
				resolvers += fmt.Sprintf("\t%s(%s): %s!\n", e.ResolverName, resolverArgs, e.Name)
			}
		}
	}

	if len(f.Entities) == 0 {
		// It's unusual for a service not to have any entities, but
		// possible if it only exports top-level queries and mutations.
		return nil
	}

	// resolvers can be empty if a service defines only "empty
	// extend" types.  This should be rare.
	if resolvers != "" {
		resolvers = entityResolverInputDefinitions + `
# fake type to build resolver interfaces for users to implement
type Entity {
	` + resolvers + `
}
`
	}

	return &ast.Source{
		Name:    "federation/entity.graphql",
		BuiltIn: true,
		Input: `
# a union of all types that use the @key directive
union _Entity = ` + entities + `
` + resolvers + `
type _Service {
  sdl: String
}

extend type Query {
  _entities(representations: [_Any!]!): [_Entity]!
  _service: _Service!
}
`,
	}
}

// Entity represents a federated type
// that was declared in the GQL schema.
type Entity struct {
	// TODO(miguel): encapsulate resolver information in its own
	// struct for future work to support multiple federated keys.

	Name         string      // The same name as the type declaration
	KeyFields    []*KeyField // The fields declared in @key.
	ResolverName string      // The resolver name, such as FindUserByID
	InputType    string      // The Go generated input type for multi entity resolvers
	Def          *ast.Definition
	Requires     []*Requires
	Multi        bool
}

type KeyField struct {
	Definition *ast.FieldDefinition
	Field      fieldset.Field        // len > 1 for nested fields
	Type       *config.TypeReference // The Go representation of that field type
}

// Requires represents an @requires clause
type Requires struct {
	Name  string                // the name of the field
	Field fieldset.Field        // source Field, len > 1 for nested fields
	Type  *config.TypeReference // The Go representation of that field type
}

func (e *Entity) allFieldsAreExternal() bool {
	for _, field := range e.Def.Fields {
		if field.Directives.ForName("external") == nil {
			return false
		}
	}
	return true
}

func (f *federation) GenerateCode(data *codegen.Data) error {
	if len(f.Entities) > 0 {
		if data.Objects.ByName("Entity") != nil {
			data.Objects.ByName("Entity").Root = true
		}
		for _, e := range f.Entities {
			obj := data.Objects.ByName(e.Def.Name)

			// fill in types for key fields
			//
			for _, keyField := range e.KeyFields {
				if len(keyField.Field) == 0 {
					fmt.Println("skipping key field " + keyField.Definition.Name + " in " + e.Def.Name)
					continue
				}
				cgField := keyField.Field.TypeReference(obj, data.Objects)
				keyField.Type = cgField.TypeReference
			}

			// fill in types for requires fields
			//
			for _, reqField := range e.Requires {
				if len(reqField.Field) == 0 {
					fmt.Println("skipping requires field " + reqField.Name + " in " + e.Def.Name)
					continue
				}
				cgField := reqField.Field.TypeReference(obj, data.Objects)
				reqField.Type = cgField.TypeReference
			}
		}
	}

	return templates.Render(templates.Options{
		PackageName:     data.Config.Federation.Package,
		Filename:        data.Config.Federation.Filename,
		Data:            f,
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
	})
}

func (f *federation) setEntities(schema *ast.Schema) {
	for _, schemaType := range schema.Types {
		if schemaType.Kind == ast.Interface {
			// TODO: support @key and @extends for interfaces
			if dir := schemaType.Directives.ForName("key"); dir != nil {
				panic("@key directive is not currently supported for interfaces.")
			}
			if dir := schemaType.Directives.ForName("extends"); dir != nil {
				panic("@extends directive is not currently supported for interfaces.")
			}
			continue
		}
		if schemaType.Kind == ast.Object {
			keys := schemaType.Directives.ForNames("key")
			if len(keys) > 1 {
				// TODO: support multiple keys -- multiple resolvers per Entity
				panic("only one @key directive currently supported")
			}

			if len(keys) > 0 {
				dir := keys[0]
				if len(dir.Arguments) != 1 || dir.Arguments[0].Name != "fields" {
					panic("Exactly one `fields` argument needed for @key declaration.")
				}
				arg := dir.Arguments[0]
				keyFieldSet := fieldset.New(arg.Value.Raw, nil)

				// TODO: why is this nested inside the @key handling? -- because it's per-Entity, and we make one per @key
				requires := []*Requires{}
				for _, f := range schemaType.Fields {
					dir := f.Directives.ForName("requires")
					if dir == nil {
						continue
					}
					requiresFieldSet := fieldset.New(dir.Arguments[0].Value.Raw, nil)
					for _, field := range requiresFieldSet {
						requires = append(requires, &Requires{
							Name:  field.ToGoPrivate(),
							Field: field,
						})
					}
				}

				keyFields := make([]*KeyField, len(keyFieldSet))
				for i, field := range keyFieldSet {
					def := field.FieldDefinition(schemaType, schema)

					if def == nil {
						panic(fmt.Sprintf("no field for %v", field))
					}

					keyFields[i] = &KeyField{Definition: def, Field: field}
				}

				e := &Entity{
					Name:      schemaType.Name,
					KeyFields: keyFields,
					Def:       schemaType,
					Requires:  requires,
				}

				// Let's process custom entity resolver settings.
				dir = schemaType.Directives.ForName("entityResolver")
				if dir != nil {
					if dirArg := dir.Arguments.ForName("multi"); dirArg != nil {
						if dirVal, err := dirArg.Value.Value(nil); err == nil {
							e.Multi = dirVal.(bool)
						}
					}
				}

				// If our schema has a field with a type defined in
				// another service, then we need to define an "empty
				// extend" of that type in this service, so this service
				// knows what the type is like.  But the graphql-server
				// will never ask us to actually resolve this "empty
				// extend", so we don't require a resolver function for
				// it.  (Well, it will never ask in practice; it's
				// unclear whether the spec guarantees this.  See
				// https://github.com/apollographql/apollo-server/issues/3852
				// ).  Example:
				//    type MyType {
				//       myvar: TypeDefinedInOtherService
				//    }
				//    // Federation needs this type, but
				//    // it doesn't need a resolver for it!
				//    extend TypeDefinedInOtherService @key(fields: "id") {
				//       id: ID @external
				//    }
				if !e.allFieldsAreExternal() {
					resolverFields := []string{}
					for _, f := range e.KeyFields {
						resolverFields = append(resolverFields, f.Field.ToGo())
					}

					resolverFieldsToGo := schemaType.Name + "By" + strings.Join(resolverFields, "And")
					if e.Multi {
						resolverFieldsToGo += "s" // Pluralize for better API readability
						e.ResolverName = fmt.Sprintf("findMany%s", resolverFieldsToGo)
					} else {
						e.ResolverName = fmt.Sprintf("find%s", resolverFieldsToGo)
					}

					e.InputType = resolverFieldsToGo + "Input"
				}

				f.Entities = append(f.Entities, e)
			}
		}
	}

	// make sure order remains stable across multiple builds
	sort.Slice(f.Entities, func(i, j int) bool {
		return f.Entities[i].Name < f.Entities[j].Name
	})
}
