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
}

func ConfigCtxFromFile(filename string) *ConfigCtx {
	spec := &ConfigSpec{}
	toml.DecodeFile(filename, spec)
	return &ConfigCtx{
		Spec:         spec,
		SpecFilename: filename,
	}
}
