package controllers

import (
	"MyBlog/config"
	"MyBlog/internal/models"
	"MyBlog/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// 创建文章请求体
type CreateArticleRequest struct {
	Title      string `json:"title" binding:"required"`
	Slug       string `json:"slug" binding:"required"`
	Summary    string `json:"summary"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint   `json:"category_id"`
	Tags       []uint `json:"tags"`
}

// 更新文章请求体
type UpdateArticleRequest struct {
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	CategoryID uint   `json:"category_id"`
	Tags       []uint `json:"tags"`
}

// 创建文章
func CreateArticle(c *gin.Context) {
	var req CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO 写用户昵称功能

	// 检查 slug 是否已存在
	var existArticle models.Article
	if err := config.DB.Where("slug = ?", req.Slug).First(&existArticle).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该 slug 已存在"})
		return
	}
	// 1️⃣ 获取请求头的 token
	tokenString := c.GetHeader("Authorization")
	if len(tokenString) < 8 || tokenString[:7] != "Bearer " {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少或格式错误的 token"})
		return
	}
	tokenString = tokenString[7:]

	// 2️⃣ 解析 JWT
	claims, err := utils.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}
	article := models.Article{
		Title:      req.Title,
		Slug:       req.Slug,
		Summary:    req.Summary,
		Content:    req.Content,
		AuthorID:   claims.UserID,
		CategoryID: req.CategoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := config.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}

	// 关联标签
	if len(req.Tags) > 0 {
		var tags []models.Tag
		config.DB.Where("id IN ?", req.Tags).Find(&tags)
		config.DB.Model(&article).Association("Tags").Append(&tags)
	}

	// 重新加载文章，确保返回 JSON 时 Category 和 Tags 都有
	config.DB.Preload("Category").Preload("Tags").First(&article, article.ID)

	c.JSON(http.StatusOK, gin.H{"message": "文章创建成功", "article": article})
}

// 获取文章列表（分页）
func GetArticles(c *gin.Context) {
	var articles []models.Article
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	offset := 0
	if p, err := strconv.Atoi(page); err == nil {
		offset = (p - 1) * 10
	}

	l, _ := strconv.Atoi(limit)

	// Preload 加载 Category 和 Tags
	if err := config.DB.Preload("Category").Preload("Tags").Limit(l).Offset(offset).Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"articles": articles})
}

// 获取文章详情
func GetArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	// 使用 Preload 加载关联的 Category 和 Tags
	if err := config.DB.Preload("Category").Preload("Tags").First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"article": article})
}

// 更新文章
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var req UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var article models.Article
	if err := config.DB.Preload("Tags").First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 从 JWT 获取当前用户ID
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少 token"})
		return
	}
	claims, err := utils.ParseJWT(tokenString[len("Bearer "):])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token 无效"})
		return
	}

	// 权限校验
	if article.AuthorID != claims.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限更新此文章"})
		return
	}

	// 更新字段
	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Slug != "" && req.Slug != article.Slug {
		// 检查 slug 唯一性
		var count int64
		config.DB.Model(&models.Article{}).Where("slug = ? AND id != ?", req.Slug, article.ID).Count(&count)
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "slug 已存在"})
			return
		}
		article.Slug = req.Slug
	}
	if req.Summary != "" {
		article.Summary = req.Summary
	}
	if req.Content != "" {
		article.Content = req.Content
	}
	if req.CategoryID != 0 && req.CategoryID != article.CategoryID {
		// 检查分类是否存在
		var category models.Category
		if err := config.DB.First(&category, req.CategoryID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "分类不存在"})
			return
		}
		article.CategoryID = req.CategoryID
	}
	article.UpdatedAt = time.Now()

	// 保存文章
	if err := config.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 更新标签关联
	if req.Tags != nil {
		var tags []models.Tag
		if err := config.DB.Where("id IN ?", req.Tags).Find(&tags).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "标签查询失败"})
			return
		}
		if err := config.DB.Model(&article).Association("Tags").Replace(&tags); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "标签更新失败"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章更新成功", "article": article})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	userID := c.GetUint("userID")
	if article.AuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限删除此文章"})
		return
	}

	if err := config.DB.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文章失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}
