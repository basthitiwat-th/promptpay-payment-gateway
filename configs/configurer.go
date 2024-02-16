package configs

import (
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	setupConfigEnv()

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func setupConfigEnv() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
