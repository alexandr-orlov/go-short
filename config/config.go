package config

import (
	"flag"
	"os"
)

var (
	ListenAddr string
	BaseURL    string
)

func GetConfig() {
	flag.StringVar(&ListenAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "Base URL on retrun Location header")
	flag.Parse()

	SrvAddr, exists := os.LookupEnv("SERVER_ADDRESS")
	if exists {
		ListenAddr = SrvAddr
	}

	EnvUrl, exists := os.LookupEnv("BASE_URL")
	if exists {
		BaseURL = EnvUrl
	}
}
