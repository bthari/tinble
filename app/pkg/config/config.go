package config

import (
	yaml "gopkg.in/yaml.v3"
	"log"
	"os"
)

type Deployment struct {
	HttpPort int `yaml:"httpport"`
}

type Auth struct {
	JWTSecret string `yaml:"jwtsecret"`
}
type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type Config struct {
	Deployment *Deployment `yaml:"deployment"`
	Auth       *Auth       `yaml:"auth"`
	Mongo      *Database   `yaml:"mongodb"`
}

func GetConfig() *Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &config
}
