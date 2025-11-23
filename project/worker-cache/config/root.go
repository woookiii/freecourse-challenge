package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Kafka struct {
		URLS     []string
		APIKey   string
		Secret   string
		ClientId string
	} `toml:"kafka"`
}

func NewConfig(path string) *Config {
	c := new(Config)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	err = toml.NewDecoder(f).Decode(c)
	if err != nil {
		panic(err)
	}

	return c
}
