package config

import "flag"

var (
	ListenAddr string
	BaseURL    string
)

func ParseFlags() {
	flag.StringVar(&ListenAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&BaseURL, "b", "http://localhost:8080", "Base URL on retrun Location header")

	flag.Parse()
}
