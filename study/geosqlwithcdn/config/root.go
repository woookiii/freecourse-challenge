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

	Info struct {
		Port    string
		Service string
	} `toml:"info"`

	Aws struct {
		Key       string
		SecretKey string
		Region    string
		Bucket    string
	} `toml:"aws"`
}

func NewConfig(path string) *Config {
	c := new(Config)

	//open toml file from path
	if f, err := os.Open(path); err != nil {
		panic(err)
		//read toml by reader and decode it and put it to c
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
