package apiserver

import "github.com/gimaevra94/rest/storage"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	logLevel string `toml:"log_level"`
	Storage  *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		logLevel: "debug",
		Store:  store.NewConfig(),
	}
}
