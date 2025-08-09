package config

import (
	"log"
	"strconv"

	"github.com/Altair1618/Echo-Go-Template/internal/utils/env"
)

type AppEnv string

const (
	EnvDevelopment AppEnv = "development"
	EnvTesting     AppEnv = "testing"
	EnvProduction  AppEnv = "production"
)

type AppConfig struct {
	Name string
	Port string
	port int
	Env  AppEnv
}

func newAppConfig() *AppConfig {
	appCfg := &AppConfig{}
	appCfg.Load()
	appCfg.Validate()

	return appCfg
}

func (c *AppConfig) Load() {
	c.Name = env.GetString("APP_NAME", "MyApp")
	c.port = env.GetInt("APP_PORT", 8080)
	c.Port = ":" + strconv.Itoa(c.port)
	c.Env = AppEnv(env.GetString("APP_ENV", string(EnvDevelopment)))
}

func (c *AppConfig) Validate() {
	if c.port <= 0 || c.port > 65535 {
		log.Fatalf("Invalid port number: %d. Must be between 1 and 65535.", c.port)
	}

	switch c.Env {
	case EnvDevelopment, EnvTesting, EnvProduction:
		// Valid environment
	default:
		log.Fatalf("Invalid environment: %s. Must be one of: development, testing, production.", c.Env)
	}
}
