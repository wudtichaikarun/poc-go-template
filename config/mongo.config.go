package config

type EnvMongoConfig struct {
	MongoHost     string `env:"MG_DB_HOST"`
	MongoPort     int    `env:"MG_DB_PORT"`
	MongoUser     string `env:"MG_DB_USER"`
	MongoPassword string `env:"MG_DB_PASSWORD"`
	MongoDbName   string `env:"MG_DB_NAME"`
}
