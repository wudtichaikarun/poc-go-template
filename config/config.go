package config

import (
	"os"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

var CurrentConfig *EnvConfig

func InitEnvConfig() {
	CurrentConfig = Load()
}

type EnvConfig struct {
	EnvAppConfig
	EnvMongoConfig
}

type Loader interface {
	Load(f string) error
}

type Parser interface {
	Parse(s any) error
}

func LoadFromFile(fileName string, l Loader, p Parser) *EnvConfig {
	// even if I add test cases for below 2 panic cases
	// the go coverage still unconver
	// only thing can do now is wait for go:unreahable or go:coverage features to ignore this mess
	config := &EnvConfig{}
	if _, err := os.Stat(fileName); err == nil {
		if err := l.Load(fileName); err != nil {
			panic(err)
		}
	}
	if err := p.Parse(config); err != nil {
		panic(err)
	}
	return config
}

type GoDotEnvLoader struct{}

func (l *GoDotEnvLoader) Load(fileName string) error {
	return godotenv.Load(fileName)
}

type EnvParser struct{}

func (p *EnvParser) Parse(s any) error {
	return env.Parse(s)
}

func Load(args ...string) *EnvConfig {
	configFile := ""
	if len(args) > 0 {
		configFile = args[0]
	} else {
		configFile = ".env"
	}
	return LoadFromFile(configFile, &GoDotEnvLoader{}, &EnvParser{})
}
