package middleware

import (
	"GoLang_Tutorial/internal/config"
	"GoLang_Tutorial/pkg/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Yêu cầu đăng nhập")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(c, http.StatusUnauthorized, "Định dạng Token không hợp lệ")
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("phương thức ký không hợp lệ: %v", token.Header["alg"])
			}
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Token hết hạn hoặc không hợp lệ")
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("userID", claims["sub"])
		c.Set("userRole", claims["role"])

		c.Next()
	}
}

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists {
			response.Error(c, http.StatusForbidden, "Không tìm thấy thông tin quyền hạn")
			c.Abort()
			return
		}

		isAllowed := false
		for _, r := range allowedRoles {
			if role == r {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			response.Error(c, http.StatusForbidden, "Bạn không có quyền thực hiện hành động này")
			c.Abort()
			return
		}

		c.Next()
	}
}
