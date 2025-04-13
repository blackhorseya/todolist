package usecase

import (
	"context"

	"github.com/blackhorseya/todolist/internal/domain/entity"
	"github.com/blackhorseya/todolist/internal/domain/repository"
)

// CategoryUseCase 定義分類的使用案例介面
type CategoryUseCase interface {
	CreateCategory(ctx context.Context, input CreateCategoryInput) (*entity.Category, error)
	UpdateCategory(ctx context.Context, input UpdateCategoryInput) (*entity.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	GetCategory(ctx context.Context, id string) (*entity.Category, error)
	ListCategories(ctx context.Context) ([]*entity.Category, error)
}

type categoryUseCase struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryUseCase 建立分類使用案例實例
func NewCategoryUseCase(categoryRepo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
	}
}

// CreateCategoryInput 建立分類的輸入資料
type CreateCategoryInput struct {
	Name        string
	Description string
}

// UpdateCategoryInput 更新分類的輸入資料
type UpdateCategoryInput struct {
	ID          string
	Name        string
	Description string
}

func (uc *categoryUseCase) CreateCategory(ctx context.Context, input CreateCategoryInput) (*entity.Category, error) {
	category := entity.NewCategory(input.Name, input.Description)

	if err := uc.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (uc *categoryUseCase) UpdateCategory(ctx context.Context, input UpdateCategoryInput) (*entity.Category, error) {
	category, err := uc.categoryRepo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	category.Update(input.Name, input.Description)

	if err := uc.categoryRepo.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (uc *categoryUseCase) DeleteCategory(ctx context.Context, id string) error {
	return uc.categoryRepo.Delete(ctx, id)
}

func (uc *categoryUseCase) GetCategory(ctx context.Context, id string) (*entity.Category, error) {
	return uc.categoryRepo.GetByID(ctx, id)
}

func (uc *categoryUseCase) ListCategories(ctx context.Context) ([]*entity.Category, error) {
	return uc.categoryRepo.List(ctx)
}
