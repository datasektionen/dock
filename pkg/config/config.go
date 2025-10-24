package config

import (
	"flag"
)

type Config struct {
	StonPort    string
	SpamPort    string
	RfingerPort string
	ConfigFile  string
}

var config Config

func init() {
	flag.StringVar(&config.StonPort, "ston-port", "9001", "Port for the ston api service")
	flag.StringVar(&config.SpamPort, "spam-port", "9002", "Port for the spam api service")
	flag.StringVar(&config.RfingerPort, "rfinger-port", "9003", "Port for the rfinger api service")
	flag.StringVar(&config.ConfigFile, "config-file", "config.yaml", "Path to a yaml config file")

}

func GetConfig() *Config {
	flag.Parse()

	return &config
}
