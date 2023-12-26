package config

type Config struct {
	DataBase DB
}

type DB struct {
	DbType         string `env:"DB_TYPE" envDefault:"postgres://"`
	DbUser         string `env:"POSTGRES_USER" envDefault:"userL0:"`
	DbUserPassword string `env:"POSTGRES_PASSWORD" envDefault:"123456"`
	DbPort         string `env:"DB_PORT" envDefault:"@0.0.0.0:5432/"`
	DbName         string `env:"POSTGRES_DB" envDefault:"postgres"`
}

func GetConfig() *Config {
	var appConfig Config

	if err := env
}