package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gimaevra94/rest/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml",
		"path to config file")
}

func main() {
	flag.Parse()
	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	newServer := server.New(config)
	if err := newServer.Start(); err != nil {
		log.Fatal(err)
	}

}
