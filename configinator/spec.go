package configinator

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
)

type ConfigVarDef struct {
	Var        string   `toml:"var"`
	Doc        string   `toml:"doc"`
	Type       string   `toml:"type"`
	Default    string   `toml:"default"`
	Optional   bool     `toml:"optional"`
	EnumValues []string `toml:"enum_values"`
}

type ConfigSettings struct {
	Name        string `toml:"name"`
	PackageName string `toml:"package_name"`
}

type ConfigSpec struct {
	Settings ConfigSettings           `toml:"settings"`
	Vars     map[string]*ConfigVarDef `toml:"vars"`
}

type ConfigCtx struct {
	Spec         *ConfigSpec
	SpecFilename string
	Imports      []string
}

func ConfigCtxFromFile(filename string) (*ConfigCtx, error) {
	spec := &ConfigSpec{}
	_, err := toml.DecodeFile(filename, spec)
	if err != nil {
		return nil, err
	}
	imports := []string{"os"}
	// fmt is used for enums
	// errors is used if any vars without defaults are present
	importErrors, importFmt, importStrconv := false, false, false
	for varKey, varDef := range spec.Vars {
		// Vars should be type string by default
		if varDef.Type == "" {
			varDef.Type = "string"
		}
		// Check for valid types
		if varDef.Type != "string" && varDef.Type != "int64" && varDef.Type != "enum" && varDef.Type != "bool" {
			return nil, fmt.Errorf("invalid type %s for var %s", varDef.Type, varKey)
		}
		if varDef.Type == "int64" {
			importStrconv = true
		}
		// Multiline doc comments need to have comment markers inserted at every line break
		if strings.ContainsRune(varDef.Doc, '\n') {
			varDef.Doc = strings.Join(strings.Split(varDef.Doc, "\n"), "\n// ")
		}
		// Errors should be imported for any vars that can fail if they're missing
		if varDef.Default == "" && !varDef.Optional {
			importErrors = true
		}
		if varDef.Default != "" && varDef.Optional {
			return nil, fmt.Errorf("optional vars cannot have defaults (var %s). Including a default without optional = true will have the same effect", varKey)
		}
		// fmt.Errorf is used for enum types
		if varDef.Type == "enum" {
			importFmt = true
			if varDef.Optional {
				return nil, fmt.Errorf("optional enum types not allowed (var %s)", varKey)
			}
		}
	}

	if importErrors {
		imports = append(imports, "errors")
	}
	if importFmt {
		imports = append(imports, "fmt")
	}
	if importStrconv {
		imports = append(imports, "strconv")
	}

	return &ConfigCtx{
		Spec:         spec,
		SpecFilename: filename,
		Imports:      imports,
	}, nil
}
