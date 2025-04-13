package mongodb

import (
	"context"

	"github.com/blackhorseya/todolist/app/domain/entity"
	"github.com/blackhorseya/todolist/app/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoTodoRepository struct {
	client *mongo.Client
	db     *mongo.Database
	col    *mongo.Collection
}

// NewMongoTodoRepository 建立 MongoDB 待辦事項儲存庫
func NewMongoTodoRepository(client *mongo.Client) repository.TodoRepository {
	db := client.Database("todolist")
	col := db.Collection("todos")
	return &mongoTodoRepository{
		client: client,
		db:     db,
		col:    col,
	}
}

func (r *mongoTodoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	_, err := r.col.InsertOne(ctx, todo)
	return err
}

func (r *mongoTodoRepository) GetByID(ctx context.Context, id string) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.col.FindOne(ctx, bson.M{"id": id}).Decode(&todo)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *mongoTodoRepository) Update(ctx context.Context, todo *entity.Todo) error {
	_, err := r.col.ReplaceOne(ctx, bson.M{"id": todo.ID}, todo)
	return err
}

func (r *mongoTodoRepository) Delete(ctx context.Context, id string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (r *mongoTodoRepository) List(ctx context.Context, filter repository.TodoFilter) ([]*entity.Todo, error) {
	bsonFilter := bson.M{}
	if filter.CategoryID != nil {
		bsonFilter["categoryid"] = *filter.CategoryID
	}
	if filter.Status != nil {
		bsonFilter["status"] = *filter.Status
	}
	if filter.Priority != nil {
		bsonFilter["priority"] = *filter.Priority
	}

	cursor, err := r.col.Find(ctx, bsonFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []*entity.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}
