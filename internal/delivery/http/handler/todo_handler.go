package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/blackhorseya/todolist/internal/domain/entity"
	"github.com/blackhorseya/todolist/internal/domain/repository"
	"github.com/blackhorseya/todolist/internal/usecase"
	"github.com/gin-gonic/gin"
)

// TodoHandler 處理待辦事項相關的 HTTP 請求
type TodoHandler struct {
	todoUC usecase.TodoUseCase
}

// NewTodoHandler 建立待辦事項處理器
func NewTodoHandler(todoUC usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUC: todoUC,
	}
}

// CreateTodoRequest 建立待辦事項的請求資料
type CreateTodoRequest struct {
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description"`
	Priority    entity.Priority `json:"priority" binding:"required"`
	DueDate     string          `json:"dueDate" binding:"required"`
	CategoryID  string          `json:"categoryId" binding:"required"`
}

// CreateTodo 處理建立待辦事項的請求
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dueDate, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的日期格式"})
		return
	}

	todo, err := h.todoUC.CreateTodo(c.Request.Context(), usecase.CreateTodoInput{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     dueDate,
		CategoryID:  req.CategoryID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// ListTodos 處理列出待辦事項的請求
func (h *TodoHandler) ListTodos(c *gin.Context) {
	categoryID := c.Query("categoryId")
	statusStr := c.Query("status")
	priorityStr := c.Query("priority")

	var filter repository.TodoFilter
	if categoryID != "" {
		filter.CategoryID = &categoryID
	}

	if statusStr != "" {
		status := entity.TodoStatus
		if s, err := strconv.Atoi(statusStr); err == nil {
			status = entity.Status(s)
		}
		filter.Status = &status
	}

	if priorityStr != "" {
		priority := entity.Low
		if p, err := strconv.Atoi(priorityStr); err == nil {
			priority = entity.Priority(p)
		}
		filter.Priority = &priority
	}

	todos, err := h.todoUC.ListTodos(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetTodo 處理獲取單一待辦事項的請求
func (h *TodoHandler) GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.todoUC.GetTodo(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if todo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到待辦事項"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateTodoRequest 更新待辦事項的請求資料
type UpdateTodoRequest struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Priority    entity.Priority `json:"priority"`
	Status      entity.Status   `json:"status"`
	DueDate     string          `json:"dueDate"`
	CategoryID  string          `json:"categoryId"`
}

// UpdateTodo 處理更新待辦事項的請求
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var req UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dueDate, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的日期格式"})
		return
	}

	todo, err := h.todoUC.UpdateTodo(c.Request.Context(), usecase.UpdateTodoInput{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      req.Status,
		DueDate:     dueDate,
		CategoryID:  req.CategoryID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo 處理刪除待辦事項的請求
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := h.todoUC.DeleteTodo(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
