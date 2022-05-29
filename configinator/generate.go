package configinator

import (
	"bytes"
	_ "embed"
	"html/template"
	"strings"
)

//go:embed config.tpl
var configSrc string
var configTemplate = template.Must(template.New("config").Funcs(
	map[string]interface{}{
		"export":   export,
		"unexport": unexport,
	},
).Parse(configSrc))

func export(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func unexport(s string) string {
	uppers := 0
	for i := 0; i < len(s); i++ {
		if strings.ToUpper(s[0:i]) != s[0:i] {
			break
		}
		uppers++
	}
	return strings.ToLower(s[0:uppers]) + s[uppers:]
}

func ExecuteTemplate(buf *bytes.Buffer, configCtx *ConfigCtx) error {
	return configTemplate.Execute(buf, configCtx)
}
