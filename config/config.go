package config

type Config struct {
	DBConnectionString string
	RedisAddress       string
	RedisPassword      string
	RedisDB            int
}

func NewConfig() *Config {
	return &Config{
		DBConnectionString: "postgres://postgres:admin@localhost/todoapp?sslmode=disable",
		RedisAddress:       "localhost:6379",
		RedisPassword:      "",
		RedisDB:            0,
	}
}
