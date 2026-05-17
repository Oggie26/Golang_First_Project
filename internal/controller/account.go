package controller

import (
	"GoLang_Tutorial/internal/config"
	"GoLang_Tutorial/internal/dto"
	"GoLang_Tutorial/internal/middleware"
	"GoLang_Tutorial/internal/service"
	"GoLang_Tutorial/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AccountController struct {
	service  service.AccountService
	validate *validator.Validate
	cfg      *config.Config
}

func NewAccountController(r *gin.Engine, s service.AccountService, cfg *config.Config) {
	c := &AccountController{
		service:  s,
		validate: validator.New(),
		cfg:      cfg,
	}

	// Public Routes
	api := r.Group("/api/v1/accounts")
	{
		api.POST("/register", c.Register)
		api.POST("/login", c.Login)
	}

	// Private Routes (Cần đăng nhập)
	auth := r.Group("/api/v1/accounts")
	auth.Use(middleware.AuthMiddleware(cfg))
	{
		auth.GET("/profile", c.GetProfile)
		auth.PUT("/profile", c.UpdateProfile)
	}
}

func (c *AccountController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Dữ liệu đầu vào không hợp lệ")
		return
	}

	if err := c.validate.Struct(req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Lỗi kiểm tra dữ liệu: "+err.Error())
		return
	}

	res, err := c.service.Register(ctx.Request.Context(), &req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(ctx, "Đăng ký thành công!", res)
}

func (c *AccountController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Dữ liệu đầu vào không hợp lệ")
		return
	}

	if err := c.validate.Struct(req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Lỗi kiểm tra dữ liệu: "+err.Error())
		return
	}

	res, err := c.service.Login(ctx.Request.Context(), &req)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(ctx, "Đăng nhập thành công!", res)
}

func (c *AccountController) GetProfile(ctx *gin.Context) {
	// Lấy userID từ context (do AuthMiddleware lưu vào)
	accountID, exists := ctx.Get("userID")
	if !exists {
		response.Error(ctx, http.StatusUnauthorized, "Không tìm thấy thông tin đăng nhập")
		return
	}

	user, err := c.service.GetProfile(ctx.Request.Context(), accountID.(string))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Lỗi lấy thông tin cá nhân")
		return
	}

	response.Success(ctx, "Lấy thông tin thành công", user)
}

func (c *AccountController) UpdateProfile(ctx *gin.Context) {
	accountID, _ := ctx.Get("userID")

	var req struct {
		FullName string `json:"full_name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Dữ liệu không hợp lệ")
		return
	}

	err := c.service.UpdateProfile(ctx.Request.Context(), accountID.(string), req.FullName, req.Email)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Lỗi cập nhật thông tin")
		return
	}

	response.Success(ctx, "Cập nhật thành công", nil)
}
