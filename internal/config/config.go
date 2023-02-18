package config

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		Name string `mapstructure:"name"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"pass"`
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$GOPATH/src/github.com/weehongayden/finance-flow-api/internal/config")
	viper.AutomaticEnv()

	var config Config

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
