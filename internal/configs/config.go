package configs

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen struct {
		BindIp string `yaml:"bind_ip" env:"BIND_IP" env-description:"Bind IP"`
		Port   string `yaml:"port" env:"PORT" env-description:"Bind port"`
	} `yaml:"listen" env:"LISTEN" env-description:"Listen config"`
	Redis struct {
		Addr     string `yaml:"addr" env:"REDIS_ADDR" env-description:"Redis address"`
		Password string `yaml:"password" env:"REDIS_PASSWORD" env-description:"Redis password"`
	} `yaml:"redis" env:"REDIS" env-description:"Redis config"`
	Apitoken string `yaml:"apitoken" env:"APITOKEN" env-description:"Api token for access to data.egov.kz"`
}

var instance Config

func GetConfig(filepath string) *Config {
	log.Println("read appication config...")
	if err := cleanenv.ReadConfig(filepath, &instance); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		log.Println(help)
	}

	return &instance
}
