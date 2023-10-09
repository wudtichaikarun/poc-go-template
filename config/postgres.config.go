package config

type EnvPGConfig struct {
	PGHost     string `env:"PG_DB_HOST"`
	PGPort     int    `env:"PG_DB_PORT"`
	PGUser     string `env:"PG_DB_USER"`
	PGPassword string `env:"PG_DB_PASSWORD"`
	PGDbName   string `env:"PG_DB_NAME"`
	PGSslMode  string `env:"PG_DB_SSLMODE"`
	PGTimezone string `env:"PG_DB_TIMEZONE"`
}
