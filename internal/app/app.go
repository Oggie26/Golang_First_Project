package app

//go:generate go run github.com/swaggo/swag/cmd/swag@latest init -g ../../cmd/api/main.go -o ../../docs -d ./,../../internal/controller,../../cmd/api --parseDependency

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "GoLang_Tutorial/docs"
	"GoLang_Tutorial/internal/config"
	"GoLang_Tutorial/internal/controller"
	"GoLang_Tutorial/internal/middleware"
	"GoLang_Tutorial/internal/repository"
	"GoLang_Tutorial/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(cfg *config.Config) {
	log.Println("Đang khởi động hệ thống...")
	log.Println("Đã kết nối Database")
	log.Println("Đã tiêm Dependencies")
	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRepo := repository.NewMemoryUserRepository()
	userService := service.NewUserService(userRepo, cfg)
	controller.NewUserController(r, userService)
	srv := &http.Server{
		Addr:    cfg.HTTP.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Server đang chạy tại cổng %s\n", cfg.HTTP.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Lỗi khởi động server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Nhận lệnh tắt server. Đang dọn dẹp...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server tắt bất thường:", err)
	}

	log.Println("✅ Server đã thoát an toàn.")
}
