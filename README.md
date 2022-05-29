# configinator

`configinator` is a small utility for generating configuration initializers from a TOML-based spec file. See `example/config_spec.toml` for an example spec, and `example/config.go` for an example of the code this tool generates.

## Installation

```bash
go install github.com/olafal0/configinator@latest
```

## Usage

```go
package config

//go:generate configinator -specfile config_spec.toml
```

Then run:

```bash
go generate
```

## Spec

The spec is written as a [TOML file](https://toml.io/) with two sections: `settings` and `vars`.

### `settings`

- `name`: The name of the service you're configuring. Used to name the generated struct types, enum types, constants, and functions.
- `package_name`: The name of the go package to use in the generated file(s).

### `vars`

`vars` is a map of variable names to variable definitions. For example,

```toml
[vars.PGUsername]
var = "FOOBAR_PG_USERNAME"
type = "string"
default = "postgres"
```

Here, `PGUsername` is the variable name, and the generated config struct will have a method `PGUsername() string`.

- `var`: The environment variable key to use to load this variable.
- `type`: The type of the loaded variable. Currently, only `string` and `enum` are supported.
- `default`: The value to set the variable to if the environment variable is not set. Not supported for `enum` type vars.
- `enum_values`: A list of potential string values that an `enum` type var must be set to. If the environment variable is not one of these values, config loading will fail.
