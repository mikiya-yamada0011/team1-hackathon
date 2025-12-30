package config

import (
"log/slog"
"net/url"
"sync"
)

const DefaultConfigPath = "config/config.yaml"

var (
once           sync.Once
configInstance *Config
configLoader   ConfigLoader
)

type Config struct {
Database DatabaseConfig `yaml:"database"`
Server   ServerConfig   `yaml:"server"`
CORS     CorsConfig     `yaml:"cors"`
}

type DatabaseConfig struct {
Port     string `yaml:"port" env:"DATABASE_PORT" envDefault:"5432"`
Host     string `yaml:"host" env:"DATABASE_HOST" envDefault:"db"`
User     string `yaml:"user" env:"DATABASE_USER" envDefault:"myuser"`
Password string `yaml:"password" env:"DATABASE_PASSWORD" envDefault:"mypassword"`
Name     string `yaml:"name" env:"DATABASE_NAME" envDefault:"mydb"`
}

type ServerConfig struct {
Port string `yaml:"port" env:"SERVER_PORT" envDefault:"8080"`
}

type CorsConfig struct {
AllowedOrigins   []string `yaml:"allowedOrigins" env:"CORS_ALLOWED_ORIGINS" envSeparator:"," envDefault:"http://localhost:3000"`
AllowedMethods   []string `yaml:"allowedMethods" env:"CORS_ALLOWED_METHODS" envSeparator:"," envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
AllowedHeaders   []string `yaml:"allowedHeaders" env:"CORS_ALLOWED_HEADERS" envSeparator:"," envDefault:"Content-Type,Authorization"`
AllowCredentials bool     `yaml:"allowCredentials" env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
}

func (c DatabaseConfig) GetDSN() string {
var password string
if c.Password != "" {
password = url.QueryEscape(c.Password)
}
return "postgres://" + c.User + ":" + password + "@" + c.Host + ":" + c.Port + "/" + c.Name + "?sslmode=disable"
}

func GetConfig() (*Config, error) {
var loadErr error
once.Do(func() {
if configLoader == nil {
configLoader = NewYAMLConfigLoader()
}

config, err := configLoader.LoadWithEnv(DefaultConfigPath)
if err != nil {
slog.Error("設定の読み込みに失敗しました", "error", err, "path", DefaultConfigPath)
loadErr = err
return
}
configInstance = config
slog.Info("設定の読み込みに成功しました")
})

if loadErr != nil {
return nil, loadErr
}

return configInstance, nil
}
