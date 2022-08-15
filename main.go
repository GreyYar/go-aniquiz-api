package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"go-aniquiz-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/apiserver.toml", "path to apiserver config file")
}

func main() {
	flag.Parse()

	fmt.Println("Hello World!")

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
