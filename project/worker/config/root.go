package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	DB struct {
		Database string
		URL      string
	} `toml:"db"`

	Kafka struct {
		URLS    []string
		GroupID string
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
