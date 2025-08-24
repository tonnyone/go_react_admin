package config

import (
	"github.com/spf13/viper"
)

func LoadConfigWithPath(configPath string) (*AppConfig, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if configPath != "" {
		v.AddConfigPath(configPath)
	}
	v.AddConfigPath(".")
	v.AddConfigPath("..")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var c AppConfig
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
