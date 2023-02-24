package config

import (
	"github.com/spf13/viper"
)

type Server struct {
	Address string `mapstructure:"address"`
	Port    uint16 `mapstructure:"port"`
}

type Config struct {
	Server   Server `mapstructure:"server"`
	Database struct {
		Host string `mapstructure:"host"`
		Port int16  `mapstructure:"port"`
		Name string `mapstructure:"name"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
	}
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$GOPATH/src/github.com/weehongayden/finance-flow-api/internal/config")
	viper.AutomaticEnv()

	var config Config

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
