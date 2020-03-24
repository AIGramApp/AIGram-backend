package config

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfig inits the configuration
func LoadConfig(logger *logrus.Logger) *AppConfiguration {
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
	currentConfig.Print(logger)
	return currentConfig
}
