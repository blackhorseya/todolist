package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/blackhorseya/todolist/app/domain/entity"
	"github.com/blackhorseya/todolist/app/domain/repository"
	"github.com/blackhorseya/todolist/app/usecase"
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

// @Summary     建立待辦事項
// @Description 建立新的待辦事項
// @Tags        待辦事項
// @Accept      json
// @Produce     json
// @Param       request body     CreateTodoRequest true "待辦事項資訊"
// @Success     201    {object}  entity.Todo
// @Failure     400    {object}  map[string]string
// @Failure     500    {object}  map[string]string
// @Router      /v1/todos [post]
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

// @Summary     列出待辦事項
// @Description 取得所有待辦事項列表，可依照分類、狀態和優先級別進行過濾
// @Tags        待辦事項
// @Accept      json
// @Produce     json
// @Param       categoryId query    string  false "分類 ID"
// @Param       status     query    int     false "狀態"
// @Param       priority   query    int     false "優先級別"
// @Success     200       {array}   entity.Todo
// @Failure     500       {object}  map[string]string
// @Router      /v1/todos [get]
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

// @Summary     取得待辦事項
// @Description 透過 ID 取得特定待辦事項
// @Tags        待辦事項
// @Accept      json
// @Produce     json
// @Param       id  path     string true "待辦事項 ID"
// @Success     200 {object} entity.Todo
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /v1/todos/{id} [get]
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

// @Summary     更新待辦事項
// @Description 更新特定待辦事項的資訊
// @Tags        待辦事項
// @Accept      json
// @Produce     json
// @Param       id      path     string           true "待辦事項 ID"
// @Param       request body     UpdateTodoRequest true "更新資訊"
// @Success     200     {object} entity.Todo
// @Failure     400     {object} map[string]string
// @Failure     500     {object} map[string]string
// @Router      /v1/todos/{id} [put]
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

// @Summary     刪除待辦事項
// @Description 刪除特定待辦事項
// @Tags        待辦事項
// @Accept      json
// @Produce     json
// @Param       id  path     string true "待辦事項 ID"
// @Success     204 {string} string "No Content"
// @Failure     500 {object} map[string]string
// @Router      /v1/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := h.todoUC.DeleteTodo(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
