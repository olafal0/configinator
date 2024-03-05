# configinator

`configinator` is a small utility for generating configuration initializers from a TOML-based spec file. See `example/config_spec.toml` for an example spec, and `example/config.go` for an example of the code this tool generates.

⚠️ Please note that this library is not injection-safe! Don't use this on config specs that you didn't write or haven't reviewed.

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

- `name`: The name of the service you're configuring. Used to name the generated struct types, enum types, constants, and functions. For example, if `name` is `Foobar`, the generated config struct type will be named `FoobarConfig`.
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
- `type`: The type of the loaded variable. Currently, only `string`, `enum`, `bool`, and `int64` are supported.
- `default`: (string) The expression to set the variable to if the environment variable is not set. Default values can be set for enums, but not naturally - see below. Conflicts with `optional`.
- `enum_values`: A list of potential string values that an `enum` type var must be set to. If the environment variable is not one of these values, config loading will fail.
- `optional`: A boolean where a true value means an unset environment variable will not cause an error. Conficts with `default`.

`default` and `optional` conflict because setting a default implies that the variable is optional, and setting optional implies that there is no default value needed. If both are specified, generation will return an error.

### Enums

Enum types have some special handling. In the example case of `Foobar`'s `vars.Environment`, there are three values: `local`, `dev`, and `production`. These are title-cased and used to generate a string enum type:

```go
type FoobarEnvironment string

const (
	FoobarEnvironmentLocal FoobarEnvironment = "local"
	FoobarEnvironmentDev   FoobarEnvironment = "dev"
	FoobarEnvironmentProd  FoobarEnvironment = "prod"
)
```

These enum types can be compared to. In addition, there are generated methods on the config struct type for each value. For example, `IsEnvironmentLocal() bool` returns true if the value is equal to `local`.

- If `FOOBAR_DEPLOY_ENV` is set, the generated loader will set `Environment` to `FoobarEnvironmentLocal` if the value is `local`, and so on; and will return an error if the value is not one of the expected ones.
- If the `FOOBAR_DEPLOY_ENV` is not set, the environment will be set to the literal default value. This means, if you try to set a default value of `local`, the generated code will look like this:

```go
} else {
  cfg.environment = local
}
```

...which will cause a compilation error, because local is an undefined symbol. Instead, you can configure the default like this:

```toml
[vars.Environment]
var = "FOOBAR_DEPLOY_ENV"
type = "enum"
enum_values = ["local", "dev", "prod"]
default = "FoobarEnvironmentLocal"
```

This also means that, for `default` specifically, you can use any string that's a valid Go expression:

```toml
[vars.Environment]
var = "FOOBAR_DEPLOY_ENV"
type = "enum"
enum_values = ["local", "dev", "prod"]
default = "FoobarEnvironmentLocal // Set to local by default"
```
