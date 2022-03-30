/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 27/03/22, 17:17
 * Copyright (c) 2022
 */

package conf

import (
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	Name   string `mapstructure:"name"`
	Assets struct {
		DefaultPath string `mapstructure:"default_path"`
	} `mapstructure:"assets"`
	JWT struct {
		Key             string `mapstructure:"key"`
		Issuer          string `mapstructure:"issuer"`
		Duration        string `mapstructure:"duration"`
		RefreshDuration string `mapstructure:"refresh_duration"`
	} `mapstructure:"jwt"`
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Mysql struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	} `mapstructure:"mysql"`
	GoogleAuth struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"google_auth"`
}

func LoadConfig() (app AppConfig, err error) {
	path := os.Getenv("CONFIG_PATH")
	env := os.Getenv("CONFIG_ENV")

	viper.AddConfigPath(path)
	viper.SetConfigName("app-" + env)
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&app)
	return
}
