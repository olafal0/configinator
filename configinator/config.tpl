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

{{- /* Generate constants for var keys */}}

const (
  {{- range $varName, $varDef := .Spec.Vars}}
  {{$.Spec.Settings.Name}}ConfigKey{{$varName}} = "{{$varDef.Var}}"
  {{- end}}
)

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
  {{- if $varDef.Optional }}
  {{- if eq $varDef.Type "int64" }}
  if val, ok := os.LookupEnv({{$.Spec.Settings.Name}}ConfigKey{{$varName}}); ok {
    if converted, err := strconv.ParseInt(val, 10, 64); err == nil {
      cfg.{{unexport $varName}} = converted
    } else {
      return nil, err
    }
  } {{- if $varDef.Default}} else {
    cfg.{{unexport $varName}} = {{asInt $varDef.Default}}
  } {{- end}}
  {{- else if eq $varDef.Type "bool" }}
  if val, ok := os.LookupEnv({{$.Spec.Settings.Name}}ConfigKey{{$varName}}); ok {
    cfg.{{unexport $varName}} = val == "true"
  } {{- if $varDef.Default}} else {
    cfg.{{unexport $varName}} = {{$varDef.Default}}
  } {{- end}}
  {{- else}}
  cfg.{{unexport $varName}} = os.Getenv({{$.Spec.Settings.Name}}ConfigKey{{$varName}})
  {{- end}}
  {{- else}}
  if {{unexport $varName}}, ok := os.LookupEnv({{$.Spec.Settings.Name}}ConfigKey{{$varName}}); ok {
    {{- if eq $varDef.Type "enum"}}
    switch {{$.Spec.Settings.Name}}{{$varName}}({{unexport $varName}}) {
      {{- range $enumValue := $varDef.EnumValues}}
      case {{$.Spec.Settings.Name}}{{$varName}}{{export $enumValue}}:
        cfg.{{unexport $varName}} = {{$.Spec.Settings.Name}}{{$varName}}{{export $enumValue}}
      {{- end}}
      default:
        return nil, fmt.Errorf("unexpected {{$varDef.Var}} value: '%s'", {{unexport $varName}})
    }
    {{- else if eq $varDef.Type "int64"}}
    if converted, err := strconv.ParseInt({{unexport $varName}}, 10, 64); err == nil {
      cfg.{{unexport $varName}} = converted
    } else {
      return nil, err
    }
    {{- else if eq $varDef.Type "bool"}}
    cfg.{{unexport $varName}} = {{unexport $varName}} == "true"
    {{- else}}
    cfg.{{unexport $varName}} = {{unexport $varName}}
    {{- end}}
	} else {
  {{- if $varDef.Default}}
    {{- if eq $varDef.Type "string" }}
    cfg.{{unexport $varName}} = "{{$varDef.Default}}"
    {{- else if eq $varDef.Type "bool"}}
    cfg.{{unexport $varName}} = {{$varDef.Default}}
    {{- else}}
    cfg.{{unexport $varName}} = {{asInt $varDef.Default}}
    {{- end}}
  {{- else}}
		return nil, errors.New("required option missing: {{$varDef.Var}}")
  {{- end}}
  }
  {{- end}}
  {{end}}

  return cfg, nil
}

{{- /* Create booladic functions for checking enum equality */}}

{{range $varName, $varDef := .Spec.Vars}}
{{- if $varDef.Doc }}
// {{ $varDef.Doc}}
{{- end }}
{{- if eq $varDef.Type "enum"}}
func (c *{{$.Spec.Settings.Name}}Config) {{$.Spec.Settings.Name}}{{$varName}}() {{$.Spec.Settings.Name}}{{$varName}} {
  return c.{{unexport $varName}}
}
{{- range $enumValue := $varDef.EnumValues}}
func (c *{{$.Spec.Settings.Name}}Config) Is{{$varName}}{{export $enumValue}}() bool {
  return c.{{unexport $varName}} == {{$.Spec.Settings.Name}}{{$varName}}{{export $enumValue}}
}
{{- end}}
{{- else}}
func (c *{{$.Spec.Settings.Name}}Config) {{$varName}}() {{$varDef.Type}} {
  return c.{{unexport $varName}}
}
{{- end}}
{{- end}}
