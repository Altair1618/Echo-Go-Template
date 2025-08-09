package config

import (
	"log"

	"github.com/Altair1618/Echo-Go-Template/internal/utils/env"
	"github.com/joho/godotenv"
)

type Config struct {
	App *AppConfig
}

type ConfigBuilder interface {
	InitEnv() ConfigBuilder
	WithApp() ConfigBuilder
	Build() *Config
}

type configBuilder struct {
	config *Config
}

func NewConfigBuilder() ConfigBuilder {
	return &configBuilder{
		config: &Config{},
	}
}

func (b *configBuilder) InitEnv() ConfigBuilder {
	loadDotEnv := env.GetBool("LOAD_DOT_ENV", true)

	if loadDotEnv {
		if err := godotenv.Load(); err != nil {
			log.Fatalf(
				"Error loading .env file: %v\n"+
					"Hint: Either create a .env file or set LOAD_DOT_ENV=false in your environment variables.",
				err,
			)
		}
	}

	return b
}

func (b *configBuilder) WithApp() ConfigBuilder {
	b.config.App = newAppConfig()
	return b
}

func (b *configBuilder) Build() *Config {
	return b.config
}
