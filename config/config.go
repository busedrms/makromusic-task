package config

type Config struct {
	DBConnectionString string
	RedisAddress       string
	RedisPassword      string
	RedisDB            int
}

func NewConfig() *Config {
	return &Config{
		DBConnectionString: "postgres://username:password@localhost:5432/database_name",
		RedisAddress:       "localhost:6379",
		RedisPassword:      "",
		RedisDB:            0,
	}
}
