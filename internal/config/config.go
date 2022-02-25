package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() *viper.Viper {
	viper.AddConfigPath("../../")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return viper.GetViper()
}
