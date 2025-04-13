package mongodb

import (
	"context"

	"github.com/blackhorseya/todolist/app/domain/entity"
	"github.com/blackhorseya/todolist/app/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoCategoryRepository struct {
	client *mongo.Client
	db     *mongo.Database
	col    *mongo.Collection
}

// NewMongoCategoryRepository 建立 MongoDB 分類儲存庫
func NewMongoCategoryRepository(client *mongo.Client) repository.CategoryRepository {
	db := client.Database("todolist")
	col := db.Collection("categories")
	return &mongoCategoryRepository{
		client: client,
		db:     db,
		col:    col,
	}
}

func (r *mongoCategoryRepository) Create(ctx context.Context, category *entity.Category) error {
	_, err := r.col.InsertOne(ctx, category)
	return err
}

func (r *mongoCategoryRepository) GetByID(ctx context.Context, id string) (*entity.Category, error) {
	var category entity.Category
	err := r.col.FindOne(ctx, bson.M{"id": id}).Decode(&category)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *mongoCategoryRepository) Update(ctx context.Context, category *entity.Category) error {
	_, err := r.col.ReplaceOne(ctx, bson.M{"id": category.ID}, category)
	return err
}

func (r *mongoCategoryRepository) Delete(ctx context.Context, id string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (r *mongoCategoryRepository) List(ctx context.Context) ([]*entity.Category, error) {
	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []*entity.Category
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}
