package config

import "os"

const (
	defaultHTTPPort = "8080"
	defaultHost     = "localhost"
	Prod            = "prod"
	Dev             = "dev"
)

type (
	Config struct {
		Host       string
		Port       string
		ServerMode string
	}
)

func Init() *Config {
	serverMode, ok := os.LookupEnv("SERVER_MODE")
	if !ok {
		serverMode = Dev
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultHTTPPort
	}

	return &Config{host, port, serverMode}
}
