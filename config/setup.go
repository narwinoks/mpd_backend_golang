package config

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	App          AppConfig      `mapstructure:"app"`
	Database     DatabaseConfig `mapstructure:"database"`
	DatabaseTest DatabaseConfig `mapstructure:"database_test"`
	Log          LogConfig      `mapstructure:"log"`
	JWT          JWTConfig      `mapstructure:"jwt"`
	Redis        RedisConfig    `mapstructure:"redis"`
	CORS         CORSConfig     `mapstructure:"cors"`
}

type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
	AllowedHeaders []string `mapstructure:"allowed_headers"`
}

type RedisConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Port         int    `mapstructure:"port"`
	Env          string `mapstructure:"env"`
	AppName      string `mapstructure:"appname"`
	ExternalCode string `mapstructure:"external_code"`
}

type JWTConfig struct {
	Secret                 string `mapstructure:"secret"`
	AccessTokenExpiration  int    `mapstructure:"access_token_expiration"`  // in minutes
	RefreshTokenExpiration int    `mapstructure:"refresh_token_expiration"` // in days
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

func SetupLog(cfg *LogConfig) {
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05Z",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})

	// Create storage directory if not exists
	if _, err := os.Stat("./storage"); os.IsNotExist(err) {
		_ = os.Mkdir("./storage", 0755)
	}

	// Setup Lumberjack for log rotation
	logWriter := &lumberjack.Logger{
		Filename:   "./storage/mpd.log",
		MaxSize:    50, // megabytes
		MaxBackups: 30,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	// Output to both stdout and file
	multiWriter := io.MultiWriter(os.Stdout, logWriter)
	logrus.SetOutput(multiWriter)
}
