package middleware

import (
	"MyBlog/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTAuthMiddleware 校验 JWT 并把用户信息写入上下文
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求未携带 token"})
			c.Abort()
			return
		}

		// 一般格式: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 格式错误"})
			c.Abort()
			return
		}

		claims, err := utils.ParseJWT(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 无效"})
			c.Abort()
			return
		}

		// 把用户信息写入上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
