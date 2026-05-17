package main

import (
	"GoLang_Tutorial/internal/app"
	"GoLang_Tutorial/internal/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

// @title           GoLang Tutorial API
// @version         1.0
// @description     Đây là một project GoLang Tutorial với cấu trúc chuẩn.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
