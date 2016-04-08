package util

import (
	"github.com/spf13/viper"
	"fmt"
)

func ReadConfig(filename string) {
	viper.SetConfigName(filename)
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
