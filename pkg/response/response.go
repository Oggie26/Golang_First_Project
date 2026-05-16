package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func Created(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusCreated, Response{
		Code:    http.StatusCreated,
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
