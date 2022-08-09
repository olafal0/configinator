// Code generated from config_spec.toml. DO NOT EDIT.

package example

import (
	"errors"
	"fmt"
	"os"
)

type FoobarEnvironment string

const (
	FoobarEnvironmentLocal FoobarEnvironment = "local"
	FoobarEnvironmentDev   FoobarEnvironment = "dev"
	FoobarEnvironmentProd  FoobarEnvironment = "prod"
)

type FoobarConfig struct {
	enableSomething bool
	environment     FoobarEnvironment
	pgpassword      string
	pgusername      string
}

func NewFoobarConfigFromEnv() (*FoobarConfig, error) {
	cfg := &FoobarConfig{}

	if enableSomething, ok := os.LookupEnv("FOOBAR_ENABLE_SOMETHING"); ok {
		cfg.enableSomething = enableSomething == "true"
	} else {
		return nil, errors.New("required option missing: FOOBAR_ENABLE_SOMETHING")
	}

	if environment, ok := os.LookupEnv("FOOBAR_DEPLOY_ENV"); ok {
		switch FoobarEnvironment(environment) {
		case FoobarEnvironmentLocal:
			cfg.environment = FoobarEnvironmentLocal
		case FoobarEnvironmentDev:
			cfg.environment = FoobarEnvironmentDev
		case FoobarEnvironmentProd:
			cfg.environment = FoobarEnvironmentProd
		default:
			return nil, fmt.Errorf("unexpected FOOBAR_DEPLOY_ENV value: '%s'", environment)
		}
	} else {
		return nil, errors.New("required option missing: FOOBAR_DEPLOY_ENV")
	}

	if pgpassword, ok := os.LookupEnv("FOOBAR_PG_PASSWORD"); ok {
		cfg.pgpassword = pgpassword
	} else {
		return nil, errors.New("required option missing: FOOBAR_PG_PASSWORD")
	}

	if pgusername, ok := os.LookupEnv("FOOBAR_PG_USERNAME"); ok {
		cfg.pgusername = pgusername
	} else {
		cfg.pgusername = "postgres"
	}

	return cfg, nil
}

func (c *FoobarConfig) EnableSomething() bool {
	return c.enableSomething
}
func (c *FoobarConfig) FoobarEnvironment() FoobarEnvironment {
	return c.environment
}
func (c *FoobarConfig) IsEnvironmentLocal() bool {
	return c.environment == FoobarEnvironmentLocal
}
func (c *FoobarConfig) IsEnvironmentDev() bool {
	return c.environment == FoobarEnvironmentDev
}
func (c *FoobarConfig) IsEnvironmentProd() bool {
	return c.environment == FoobarEnvironmentProd
}
func (c *FoobarConfig) PGPassword() string {
	return c.pgpassword
}
func (c *FoobarConfig) PGUsername() string {
	return c.pgusername
}
