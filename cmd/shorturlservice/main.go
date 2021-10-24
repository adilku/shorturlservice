package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/adilku/shorturlservice/internal/app/shorturlserver"
	"log"
)

var (
	configPath string
	databaseOption string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/shorturlservice.toml", "path to config file")
	flag.StringVar(&databaseOption, "db-option", "postgres", "database option")
}

func main() {
	flag.Parse()
	config := shorturlserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	if err := shorturlserver.Start(config, databaseOption); err != nil {
		log.Fatal(err)
	}
}


