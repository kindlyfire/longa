package main

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseUri  string `mapstructure:"database_uri"`
	ListenAddr   string `mapstructure:"listen_addr"`
	HomeRedirect string `mapstructure:"home_redirect"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
}

func (c Config) Print() {
	fmt.Printf("DatabaseUri:   %v\n", c.DatabaseUri)
	fmt.Printf("ListenAddr:    %v\n", c.ListenAddr)
	fmt.Printf("HomeRedirect:  %v\n", c.HomeRedirect)
	fmt.Printf("Username:      %v\n", c.Username)
	fmt.Printf("Password:      %v\n", c.Password)
}

func loadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.BindEnv("DatabaseUri", "L_DATABASE_URI")
	viper.BindEnv("ListenAddr", "L_LISTEN_ADDR")
	viper.BindEnv("HomeRedirect", "L_HOME_REDIRECT")
	viper.BindEnv("Username", "L_USERNAME")
	viper.BindEnv("Password", "L_PASSWORD")

	viper.SetDefault("database_uri", "file:data.db?cache=shared&_fk=1")
	viper.SetDefault("listen_addr", ":3000")
	viper.SetDefault("home_redirect", "https://github.com/kindlyfire/longa")
	viper.SetDefault("username", "admin")
	viper.SetDefault("password", "admin")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn("no config file found, using defaults or env variables")
		} else {
			return err
		}
	}

	return viper.Unmarshal(&config)
}
