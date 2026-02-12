package config

import "os"

const Version = "0.1.0"

type Config struct {
	ListenAddr    string
	EndpointsFile string
}

func Load() *Config {
	return &Config{
		ListenAddr:    envOrDefault("LISTEN_ADDR", ":4322"),
		EndpointsFile: envOrDefault("ENDPOINTS_FILE", "endpoints.json"),
	}
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
