package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// LoadConfig inits the configuration
func LoadConfig(logger *logrus.Logger) *AppConfiguration {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")

	var currentConfig *AppConfiguration
	config.ReadInConfig()
	err := config.Unmarshal(&currentConfig)
	if os.Getenv("CONFIG") != "" {
		yaml.Unmarshal([]byte(os.Getenv("CONFIG")), currentConfig)
	}
	if err != nil {
		panic(fmt.Errorf("Cannot load the configuration file %s", err.Error()))
	}
	currentConfig.Print(logger)
	return currentConfig
}
