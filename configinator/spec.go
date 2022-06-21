package configinator

import "github.com/BurntSushi/toml"

type ConfigVarDef struct {
	Var        string   `toml:"var"`
	Type       string   `toml:"type"`
	Default    string   `toml:"default"`
	EnumValues []string `toml:"enum_values"`
}

type ConfigSettings struct {
	Name        string `toml:"name"`
	PackageName string `toml:"package_name"`
}

type ConfigSpec struct {
	Settings ConfigSettings          `toml:"settings"`
	Vars     map[string]ConfigVarDef `toml:"vars"`
}

type ConfigCtx struct {
	Spec         *ConfigSpec
	SpecFilename string
	Imports      []string
}

func ConfigCtxFromFile(filename string) *ConfigCtx {
	spec := &ConfigSpec{}
	toml.DecodeFile(filename, spec)
	imports := []string{"os"}
	// fmt is used for enums
	// errors is used if any vars without defaults are present
	importErrors, importFmt := false, false
	for _, varDef := range spec.Vars {
		if varDef.Default == "" {
			importErrors = true
		}
		if varDef.Type == "enum" {
			importFmt = true
		}
	}

	if importErrors {
		imports = append(imports, "errors")
	}
	if importFmt {
		imports = append(imports, "fmt")
	}

	return &ConfigCtx{
		Spec:         spec,
		SpecFilename: filename,
		Imports:      imports,
	}
}
