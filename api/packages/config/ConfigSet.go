package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	replacer := strings.NewReplacer("_", ".")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	// viper.Debug()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading the configs!")
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("Unable to decode into struct!")
	}
}
