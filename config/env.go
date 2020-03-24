package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Init inits the configuration
func Init() *AppConfiguration {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	var currentConfig *AppConfiguration
	config.ReadInConfig()
	err := config.Unmarshal(&currentConfig)
	if err != nil {
		panic(fmt.Errorf("Cannot load the configuration file %s", err.Error()))
	}
	fmt.Println("Current configuration ", currentConfig)
	return currentConfig
}
