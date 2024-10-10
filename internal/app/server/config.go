package server

type Config struct {
	Addr     string `toml:"addr"`
	logLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		Addr:     ":8080",
		logLevel: "debug",
	}
}
