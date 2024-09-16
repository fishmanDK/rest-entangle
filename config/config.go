package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `json:"port" yaml:"port"`
	Url  string `json:"url" yaml:"url"`
}

func New(method string) *Config{
	cfg := new(Config)

	switch method{
	case "env":
		cfg = envParce()
	case "json":

	case "yml":
		cfg = ymlParce()

	default:
		cfg = ymlParce()
	}

	return cfg
}

func ymlParce() *Config{
	configPath := "/config/config.yml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

func envParce() *Config{
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}

	url := os.Getenv("URL")
	if url == "" {
		log.Fatal("URL is not set")
	}

	return &Config{
		Port: port,
		Url: url,
	}
}

func jsonParce() *Config{
	configPath := "/config/config.json"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil{
		log.Fatal(err.Error())
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed unmalshal data json: %v", err)
	}

	return &cfg
}