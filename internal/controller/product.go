package controller

import (
	"GoLang_Tutorial/internal/dto"
	"GoLang_Tutorial/internal/middleware"
	"GoLang_Tutorial/internal/service"
	"GoLang_Tutorial/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	service  service.ProductService
	validate *validator.Validate
}

func NewProductController(r gin.IRouter, s service.ProductService) {
	c := &ProductController{
		service:  s,
		validate: validator.New(),
	}

	api := r
	{
		// Chỉ Admin mới được tạo sản phẩm
		api.POST("", middleware.RoleMiddleware("admin"), c.Create)

		// Mọi người (đã đăng nhập) đều có thể xem chi tiết
		api.GET("/:id", c.GetByID)
	}
}

func (c *ProductController) Create(ctx *gin.Context) {
	var req dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Dữ liệu không hợp lệ")
		return
	}

	if err := c.validate.Struct(req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Lỗi kiểm tra: "+err.Error())
		return
	}

	res, err := c.service.CreateProduct(ctx.Request.Context(), &req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(ctx, "Tạo sản phẩm thành công!", res)
}

func (c *ProductController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	res, err := c.service.GetProduct(ctx.Request.Context(), uint(id))
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Không tìm thấy sản phẩm")
		return
	}

	response.Success(ctx, "Lấy thông tin thành công!", res)
}
