package constants

import "github.com/spf13/viper"

type constant struct {
	config *viper.Viper
}

func NewConstants(config *viper.Viper) Constants {
	return &constant{
		config: config,
	}
}

type Constants interface {
	// Implement a new constant below
	constantsDb
}
