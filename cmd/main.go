package main

import (
	"log"

	"github.com/John-Gatsby/go-DDoS/config"
	"github.com/John-Gatsby/go-DDoS/internal/service"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	log.Println(cfg)
	svc := service.New(cfg)

	svc.Run()

	return nil
}
