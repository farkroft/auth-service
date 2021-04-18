package config

import (
	"github.com/spf13/viper"
	"gitlab.com/farkroft/auth-service/external/constants"
)

// Repository repository
type Repository interface {
	GetString(str string) string
	GetInt(str string) int
}

// Config return struct of viper
type Config struct {
	cfg *viper.Viper
}

// NewConfig return new instance of config
func NewConfig(configPath string) *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigType("yaml")
	v.SetConfigName("application")
	v.AddConfigPath(configPath)
	v.AddConfigPath(constants.EnvConfigPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err) // don't change to log
	}

	return &Config{
		cfg: v,
	}
}

// GetString return string from env var
func (c *Config) GetString(str string) string {
	return c.cfg.GetString(str)
}

func (c *Config) GetInt(str string) int {
	return c.cfg.GetInt(str)
}
