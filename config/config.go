package config

import (
	"fmt"
	viperlib "github.com/spf13/viper"
)

var Viper *viperlib.Viper

func LoadConfig() {
	Viper = viperlib.New()
	Viper.SetConfigType("yaml")
	Viper.SetConfigFile("./config.yaml")
	err := Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//viper.WatchConfig()

}
