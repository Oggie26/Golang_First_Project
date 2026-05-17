package main

import (
	"GoLang_Tutorial/internal/app"
	"GoLang_Tutorial/internal/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	fmt.Println("Config:", cfg)
	app.Run(cfg)
}
