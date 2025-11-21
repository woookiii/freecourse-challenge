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
		URLS []string
	} `toml:"kafka"`

	Info struct {
		Port    string
		Service string
	} `toml:"info"`
}

func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	}

	return c
}
