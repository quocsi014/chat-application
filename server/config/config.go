package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBConnStr     string
	JWTSecretKey  string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	ServerPort    string
}

func LoadConfig() (Config, error) {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0 // Giá trị mặc định nếu không thể chuyển đổi
	}

	return Config{
		DBConnStr:     os.Getenv("DB_CONN_STR"),
		JWTSecretKey:  os.Getenv("JWT_SECRET_KEY"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       redisDB,
		ServerPort:    os.Getenv("SERVER_PORT"),
	}, nil
}
