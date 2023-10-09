package config

type EnvAppConfig struct {
	AppEnv  string `env:"APP_ENV"`
	AppHost string `env:"APP_HOST"`
	AppPort int    `env:"APP_PORT"`
	AppName string `env:"APP_NAME"`
}
