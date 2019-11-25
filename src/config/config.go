package config

var ConfigName string

type config struct {
	ForestAppLocation string `mapstructure:"forest_app_location"`
}

var Cfg config
