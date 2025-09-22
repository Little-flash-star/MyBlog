package controllers

import (
	"MyBlog/config"
	"MyBlog/internal/models"
	"MyBlog/internal/utils" // 假设你有一个 utils 包处理 JWT 之类
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// 请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 响应结构体
type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Expires  int64  `json:"expires"` // unix 时间戳
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误", "detail": err.Error()})
		return
	}

	// 查询用户
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		// 可以选择不区分「不存在」或「密码错误」来防止用户名探测
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码不正确"})
		return
	}

	// 比较密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码不正确"})
		return
	}

	// 如果用 JWT
	tokenString, err := utils.GenerateJWT(user.ID, user.Username) // 假定 GenerateJWT 返回 token 和 error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 失败"})
		return
	}

	// 返回成功
	resp := LoginResponse{
		Username: user.Username,
		Token:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24).Unix(), // 假设 token 有效 24 小时
	}
	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "data": resp})
}
