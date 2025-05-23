package db

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	DBName   string `mapstructure:"dbname" yaml:"dbname"`
	Schema   string `mapstructure:"schema" yaml:"schema"`
	SSLMode  string `mapstructure:"sslmode" yaml:"sslmode"`
	Debug    bool   `mapstructure:"debug" yaml:"debug"`
	Driver   string `mapstructure:"driver" yaml:"driver"`
}

func NewConfig() *Config {
	return &Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		DBName:   getEnv("DB_NAME", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		Schema:   getEnv("DB_SCHEMA", "public"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		Debug:    getEnvAsBool("DB_DEBUG_MODE", true),
		Driver:   getEnv("DRIVER", "postgres"),
	}
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.ParseBool(valStr)
	if err != nil {
		log.Printf("Invalid value for %s: %s. Using default: %v", key, valStr, defaultVal)
		return defaultVal
	}
	return val
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig(configs ...*Config) *Config {
	if len(configs) > 0 && configs[0] != nil {
		return configs[0]
	}
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.dbname", "postgres")
	viper.SetDefault("database.schema", "public")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.debug", true)
	viper.SetDefault("database.driver", "postgres")
	return &Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
		Schema:   viper.GetString("database.schema"),
		SSLMode:  viper.GetString("database.sslmode"),
		Debug:    viper.GetBool("database.debug"),
		Driver:   viper.GetString("database.driver"),
	}
}
