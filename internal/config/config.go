package config

import (
	"time"
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"` // для сопоставления параметров в yaml файле с структурой конфига
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct{
	Address string `yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config { // те функции, которые должны паниковать - называем Must
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == ""{
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file does not exists: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}