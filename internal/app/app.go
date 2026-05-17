package app

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
	"GoLang_Tutorial/internal/models"
	"GoLang_Tutorial/internal/repository"
	"GoLang_Tutorial/internal/service"
	"GoLang_Tutorial/pkg/database"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(cfg *config.Config) {
	log.Println("Đang khởi động hệ thống...")

	db, err := database.NewPostgresConnection(cfg.PG.URL)
	if err != nil {
		log.Fatalf("Lỗi kết nối Database: %v", err)
	}
	defer db.Close()

	err = db.MigrateAll(models.GetModels()...)
	if err != nil {
		log.Fatalf("Lỗi Migration: %v", err)
	}

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	if cfg.Swagger.Enabled {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 1. Module Account (Public)
	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo, cfg)
	controller.NewAccountController(r, accountService, cfg)

	// 2. Module Product (Protected)
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)

	productRoutes := r.Group("/api/v1/products")
	productRoutes.Use(middleware.AuthMiddleware(cfg))
	{
		controller.NewProductController(productRoutes, productService)
	}

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
