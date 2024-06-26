// Code generated from config_spec.toml. DO NOT EDIT.

package example

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type FoobarEnvironment string

const (
	FoobarEnvironmentLocal FoobarEnvironment = "local"
	FoobarEnvironmentDev   FoobarEnvironment = "dev"
	FoobarEnvironmentProd  FoobarEnvironment = "prod"
)

const (
	FoobarConfigKeyEnableSomething = "FOOBAR_ENABLE_SOMETHING"
	FoobarConfigKeyEnvironment     = "FOOBAR_DEPLOY_ENV"
	FoobarConfigKeyMaxConnections  = "FOOBAR_MAX_CONNECTIONS"
	FoobarConfigKeyPGPassword      = "FOOBAR_PG_PASSWORD"
	FoobarConfigKeyPGUsername      = "FOOBAR_PG_USERNAME"
)

type FoobarConfig struct {
	enableSomething bool
	environment     FoobarEnvironment
	maxConnections  int64
	pgpassword      string
	pgusername      string
}

func NewFoobarConfigFromEnv() (*FoobarConfig, error) {
	cfg := &FoobarConfig{}

	if enableSomething, ok := os.LookupEnv(FoobarConfigKeyEnableSomething); ok {
		cfg.enableSomething = enableSomething == "true"
	} else {
		return nil, errors.New("required option missing: FOOBAR_ENABLE_SOMETHING")
	}

	if environment, ok := os.LookupEnv(FoobarConfigKeyEnvironment); ok {
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

	if maxConnections, ok := os.LookupEnv(FoobarConfigKeyMaxConnections); ok {
		if converted, err := strconv.ParseInt(maxConnections, 10, 64); err == nil {
			cfg.maxConnections = converted
		} else {
			return nil, err
		}
	} else {
		cfg.maxConnections = 256 * 256
	}

	cfg.pgpassword = os.Getenv(FoobarConfigKeyPGPassword)

	if pgusername, ok := os.LookupEnv(FoobarConfigKeyPGUsername); ok {
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
func (c *FoobarConfig) MaxConnections() int64 {
	return c.maxConnections
}
func (c *FoobarConfig) PGPassword() string {
	return c.pgpassword
}
func (c *FoobarConfig) PGUsername() string {
	return c.pgusername
}
