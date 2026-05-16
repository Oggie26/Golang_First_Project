package middleware

import (
	"GoLang_Tutorial/pkg/errors"
	"GoLang_Tutorial/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			if appErr, ok := err.(*errors.AppError); ok {
				response.Error(ctx, appErr.Code, appErr.Message)
				return
			}
			response.Error(ctx, http.StatusInternalServerError, err.Error())
		}
	}
}
