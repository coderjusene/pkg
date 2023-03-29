package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig(configPath, configType, configName) *viper.Viper {
	v := viper.New()

	v.SetConfigType("json")
	v.SetConfigName("app")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("没有找到app.json,请确认在该目录下添加了app.json")
			os.Exit(1)
		}
	}

	return v
}
