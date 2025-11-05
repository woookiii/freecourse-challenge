package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
}

func NewConfig(path string) *Config {
	c := new(Config)
	//empty Config type struct is bind to c

	//Open: open the path's toml file
	if file, err := os.Open(path); err != nil {
		panic(err)
		//panic will stop the program when error occur
	} else {
		defer file.Close()

		if err = toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
