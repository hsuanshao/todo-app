package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config configure application parameters
type Config struct {
	RunMode string  `yaml:"runMode"`
	Swagger Swagger `yaml:"swagger"`
	HTTP    HTTP    `yaml:"http"`
	Log     Log     `yaml:"log"`
}

// NewConfig is a constructor of config with config path
func NewConfig(path string) (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// HTTP configure HTTP server parameters
type HTTP struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// GetAddress combine host and port, format is `host:port`
func (h HTTP) GetAddress() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

// Swagger configure swagger parameters
type Swagger struct {
	Enabled bool `yaml:"enabled"`
}

// Log configure log parameters
type Log struct {
	Level string `yaml:"level"`
}
