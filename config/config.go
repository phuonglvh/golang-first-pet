package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// EnvConfiguration contains environment variables
type EnvConfiguration struct {
	Server struct {
		Port int32  `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Chat struct {
		Message struct {
			Lifetime int `yaml:"lifetime", envconfig:"CHAT_MESSAGE_LIFETIME"`
		} `yaml:"message"`
	} `yaml:"chat"`
}

// Env contains application's configs
var Env *EnvConfiguration = &EnvConfiguration{}

func init() {
	mode := os.Getenv("MODE")
	fmt.Printf("Server is running in %s mode\n", mode)
	if mode == "PRODUCTION" {
		readEnv(Env)
	} else {
		readFile(Env)
	}
	fmt.Printf("%+v\n", Env)
}

func readFile(cfg *EnvConfiguration) {
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

func readEnv(cfg *EnvConfiguration) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Printf("Load config file error: %s", err)
	os.Exit(2)
}
