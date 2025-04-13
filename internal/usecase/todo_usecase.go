package usecase

import (
	"context"
	"time"

	"github.com/blackhorseya/todolist/internal/domain/entity"
	"github.com/blackhorseya/todolist/internal/domain/repository"
)

// TodoUseCase 定義待辦事項的使用案例介面
type TodoUseCase interface {
	CreateTodo(ctx context.Context, input CreateTodoInput) (*entity.Todo, error)
	UpdateTodo(ctx context.Context, input UpdateTodoInput) (*entity.Todo, error)
	DeleteTodo(ctx context.Context, id string) error
	GetTodo(ctx context.Context, id string) (*entity.Todo, error)
	ListTodos(ctx context.Context, filter repository.TodoFilter) ([]*entity.Todo, error)
}

type todoUseCase struct {
	todoRepo     repository.TodoRepository
	categoryRepo repository.CategoryRepository
}

// NewTodoUseCase 建立待辦事項使用案例實例
func NewTodoUseCase(todoRepo repository.TodoRepository, categoryRepo repository.CategoryRepository) TodoUseCase {
	return &todoUseCase{
		todoRepo:     todoRepo,
		categoryRepo: categoryRepo,
	}
}

// CreateTodoInput 建立待辦事項的輸入資料
type CreateTodoInput struct {
	Title       string
	Description string
	Priority    entity.Priority
	DueDate     time.Time
	CategoryID  string
}

// UpdateTodoInput 更新待辦事項的輸入資料
type UpdateTodoInput struct {
	ID          string
	Title       string
	Description string
	Priority    entity.Priority
	Status      entity.Status
	DueDate     time.Time
	CategoryID  string
}

func (uc *todoUseCase) CreateTodo(ctx context.Context, input CreateTodoInput) (*entity.Todo, error) {
	// 驗證分類是否存在
	_, err := uc.categoryRepo.GetByID(ctx, input.CategoryID)
	if err != nil {
		return nil, err
	}

	todo := entity.NewTodo(
		input.Title,
		input.Description,
		input.Priority,
		input.DueDate,
		input.CategoryID,
	)

	if err := uc.todoRepo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (uc *todoUseCase) UpdateTodo(ctx context.Context, input UpdateTodoInput) (*entity.Todo, error) {
	todo, err := uc.todoRepo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	todo.Title = input.Title
	todo.Description = input.Description
	todo.Priority = input.Priority
	todo.Status = input.Status
	todo.DueDate = input.DueDate
	todo.CategoryID = input.CategoryID
	todo.UpdatedAt = time.Now()

	if err := uc.todoRepo.Update(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (uc *todoUseCase) DeleteTodo(ctx context.Context, id string) error {
	return uc.todoRepo.Delete(ctx, id)
}

func (uc *todoUseCase) GetTodo(ctx context.Context, id string) (*entity.Todo, error) {
	return uc.todoRepo.GetByID(ctx, id)
}

func (uc *todoUseCase) ListTodos(ctx context.Context, filter repository.TodoFilter) ([]*entity.Todo, error) {
	return uc.todoRepo.List(ctx, filter)
}
