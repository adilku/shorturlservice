package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/adilku/shorturlservice/internal/app/shorturlserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/walletserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	fmt.Println(configPath)
	config := shorturlserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
}


