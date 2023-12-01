package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig(configPath, configType, configName string) *viper.Viper {
	v := viper.New()

	v.SetConfigType(configType)
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("没有找到app.json")
			os.Exit(1)
		}
	}

	return v
}
