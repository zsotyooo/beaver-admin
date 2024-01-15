package config

import "api/config"

func Get() config.Config {
	return configurations
}
