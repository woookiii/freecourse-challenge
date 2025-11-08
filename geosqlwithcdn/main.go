package main

import (
	"flag"
	"fmt"
	"geosqlwithcdn/config"
)

// go run . -cfg={flexible path}
var cfgPath = flag.String("cfg", "./config.toml", "config path")

func main() {
	flag.Parse()
	config.NewConfig(*cfgPath)
	fmt.Println("start")
}
