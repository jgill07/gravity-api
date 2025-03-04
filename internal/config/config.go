package config

type ApiConfig struct {
	Port    int
	Service string
}

type Config struct {
	ApiConfig   ApiConfig
	Environment string
}

func LoadConfig() (*Config, error) {

	var cfg Config

	cfg.Environment = getEnv("environment")
	cfg.ApiConfig.Port = getIntEnv("api_port", 8080)
	cfg.ApiConfig.Service = getEnv("api_service")

	return &cfg, nil
}
