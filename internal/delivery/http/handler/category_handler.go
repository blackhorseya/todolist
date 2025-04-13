package handler

import (
	"net/http"

	"github.com/blackhorseya/todolist/internal/usecase"
	"github.com/gin-gonic/gin"
)

// CategoryHandler 處理分類相關的 HTTP 請求
type CategoryHandler struct {
	categoryUC usecase.CategoryUseCase
}

// NewCategoryHandler 建立分類處理器
func NewCategoryHandler(categoryUC usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		categoryUC: categoryUC,
	}
}

// CreateCategoryRequest 建立分類的請求資料
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CreateCategory 處理建立分類的請求
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryUC.CreateCategory(c.Request.Context(), usecase.CreateCategoryInput{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// ListCategories 處理列出所有分類的請求
func (h *CategoryHandler) ListCategories(c *gin.Context) {
	categories, err := h.categoryUC.ListCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategory 處理獲取單一分類的請求
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := h.categoryUC.GetCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到分類"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategoryRequest 更新分類的請求資料
type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateCategory 處理更新分類的請求
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryUC.UpdateCategory(c.Request.Context(), usecase.UpdateCategoryInput{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory 處理刪除分類的請求
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := h.categoryUC.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
