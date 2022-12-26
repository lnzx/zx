package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Addr        string `yaml:"addr"`
	DatabaseUrl string `yaml:"database_url"`
}

var Conf = Config{}

func init() {
	if data, err := os.ReadFile("config.yaml"); err != nil {
		log.Fatalln("read config file error", err)
	} else if err = yaml.Unmarshal(data, &Conf); err != nil {
		log.Fatalln("parse config file error", err)
	}
}
