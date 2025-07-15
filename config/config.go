package config

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server      *Server
	Db          *Db
	TokenConfig *TokenConfig
}

type Server struct {
	Port int
}

type TokenConfig struct {
	Secret string
	Exp    time.Duration
	Iss    string
}

type Db struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()

		if err != nil {
			log.Println("Warning: .env file not found, relying on environment variables only")
		}

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		configInstance = &Config{
			TokenConfig: &TokenConfig{
				Secret: viper.GetString("JWT_SECRETE"),
				Exp:    time.Hour * 24 * 3,
				Iss:    "theApplication",
			},
			Server: &Server{
				Port: viper.GetInt("SERVER_PORT"),
			},
			Db: &Db{
				Host:     viper.GetString("DB_HOST"),
				Port:     viper.GetInt("DB_PORT"),
				User:     viper.GetString("DB_USER"),
				Password: viper.GetString("DB_PASSWORD"),
				DBName:   viper.GetString("DB_DBNAME"),
				SSLMode:  viper.GetString("DB_SSLMODE"),
				TimeZone: viper.GetString("DB_TIMEZONE"),
			},
		}
	})

	return configInstance
}
