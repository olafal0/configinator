// Code generated from {{.SpecFilename}}. DO NOT EDIT.

package {{.Spec.Settings.PackageName}}

import (
  {{range $import := .Imports}}
    "{{$import}}"
  {{- end}}
)

{{/* For all enum vars, generate string enum types and const values for them */}}

{{- range $varName, $varDef := .Spec.Vars}}
{{- if eq $varDef.Type "enum"}}
type {{$.Spec.Settings.Name}}{{$varName}} string

const (
  {{- range $enumValue := $varDef.EnumValues}}
  {{$.Spec.Settings.Name}}{{$varName}}{{export $enumValue}} {{$.Spec.Settings.Name}}{{$varName}} = "{{$enumValue}}"
  {{- end}}
)
{{- end}}
{{- end}}

{{- /* Generate the main config struct type */}}

type {{$.Spec.Settings.Name}}Config struct {
  {{- range $varName, $varDef := .Spec.Vars}}
  {{- if eq $varDef.Type "enum"}}
  {{unexport $varName}} {{$.Spec.Settings.Name}}{{$varName}}
  {{- else}}
  {{unexport $varName}} {{$varDef.Type}}
  {{- end}}
  {{- end}}
}

func New{{.Spec.Settings.Name}}ConfigFromEnv() (*{{.Spec.Settings.Name}}Config, error) {
  cfg := &{{.Spec.Settings.Name}}Config{}

  {{range $varName, $varDef := .Spec.Vars}}
  if {{unexport $varName}}, ok := os.LookupEnv("{{$varDef.Var}}"); ok {
    {{- if eq $varDef.Type "enum"}}
    switch {{$.Spec.Settings.Name}}{{$varName}}({{unexport $varName}}) {
      {{- range $enumValue := $varDef.EnumValues}}
      case {{$.Spec.Settings.Name}}{{$varName}}{{export $enumValue}}:
        cfg.{{unexport $varName}} = {{$.Spec.Settings.Name}}{{$varName}}{{export $enumValue}}
      {{- end}}
      default:
        return nil, fmt.Errorf("unexpected {{$varDef.Var}} value: '%s'", {{unexport $varName}})
    }
    {{- else}}
		cfg.{{unexport $varName}} = {{unexport $varName}}
    {{- end}}
	} else {
  {{- if $varDef.Default}}
    cfg.{{unexport $varName}} = "{{$varDef.Default}}"
  {{- else}}
		return nil, errors.New("required option missing: {{$varDef.Var}}")
  {{- end}}
  }
  {{end}}

  return cfg, nil
}

{{- /* Create booladic functions for checking enum equality */}}

{{range $varName, $varDef := .Spec.Vars}}
{{- if eq $varDef.Type "enum"}}
{{- range $enumValue := $varDef.EnumValues}}
func (c *{{$.Spec.Settings.Name}}Config) Is{{export $enumValue}}{{$varName}}() bool {
  return c.{{unexport $varName}} == {{$.Spec.Settings.Name}}{{$varName}}{{export $enumValue}}
}
{{- end}}
{{- else}}
func (c *{{$.Spec.Settings.Name}}Config) {{$varName}}() {{$varDef.Type}} {
  return c.{{unexport $varName}}
}
{{- end}}
{{- end}}
