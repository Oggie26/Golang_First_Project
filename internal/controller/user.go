package controller

import (
	"GoLang_Tutorial/internal/dto"
	"GoLang_Tutorial/internal/service"
	"GoLang_Tutorial/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(r *gin.Engine, userService service.UserService) {
	controller := &UserController{
		userService: userService,
	}

	api := r.Group("/api/v1")
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
	}
}


func (c *UserController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Dữ liệu đầu vào không hợp lệ")
		return
	}

	user, err := c.userService.Register(ctx.Request.Context(), &req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(ctx, "Đăng ký thành công!", user)
}


func (c *UserController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Dữ liệu đầu vào không hợp lệ")
		return
	}

	res, err := c.userService.Login(ctx.Request.Context(), &req)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(ctx, "Đăng nhập thành công!", res)
}
