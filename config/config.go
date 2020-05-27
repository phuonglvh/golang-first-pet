package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// Config contains environment variables
type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Chat struct {
		Message struct {
			LiveTime int `yaml:"live_time"`
		} `yaml:"message"`
	} `yaml:"chat"`
}

// Cfg contains application's configs
var Cfg *Config = &Config{}

func init() {
	readFile(Cfg)
	fmt.Printf("%+v", Cfg)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Printf("Load config file error: %s", err)
	os.Exit(2)
}
