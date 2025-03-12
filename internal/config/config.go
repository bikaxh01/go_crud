package config

import (
	"flag"

	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string 
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HttpServer   `yaml:"http_server" env:"http_server" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
   
	// if the config path is not found then read from terminal
	if configPath == "" {

		flags := flag.String("CONFIG-PATH", "", "config path to be here")
		flag.Parse()

		
		configPath = *flags

		// if not passed via terminal the throw error
		if configPath == "" {
			log.Fatal("Config path not found")
		}
	}

	// check whether the file exists ot not
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
	
		log.Fatal("Config file not found")
	}
	
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
        
		log.Fatal("Unable to add config")
	}
     return &cfg
}
