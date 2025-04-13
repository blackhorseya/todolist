"use client";

import { useState, useEffect, useCallback } from "react";
import { Category, Todo, Priority, TodoStatus } from "@/types/entity";
import { CreateCategoryRequest, CreateTodoRequest } from "@/types/api";
import CategoryList from "./components/CategoryList";
import CreateCategoryForm from "./components/CreateCategoryForm";
import CreateTodoForm from "./components/CreateTodoForm";
import TodoItem from "./components/TodoItem";
import { CategoryService } from "./services/categoryService";
import { TodoService } from "./services/todoService";

export default function Home() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [selectedCategoryId, setSelectedCategoryId] = useState<string>();
  const [showCreateCategory, setShowCreateCategory] = useState(false);
  const [showCreateTodo, setShowCreateTodo] = useState(false);

  const categoryService = new CategoryService();
  const todoService = new TodoService();

  // 載入分類列表
  const loadCategories = useCallback(async () => {
    const response = await categoryService.listCategories();
    if (response.data) {
      setCategories(response.data);
    }
  }, []);

  // 載入待辦事項列表
  const loadTodos = useCallback(async () => {
    const response = await todoService.listTodos(
      selectedCategoryId ? { categoryId: selectedCategoryId } : undefined
    );
    if (response.data) {
      setTodos(response.data);
    }
  }, [selectedCategoryId]);

  // 初始載入
  useEffect(() => {
    loadCategories();
  }, [loadCategories]);

  // 當選擇的分類改變時重新載入待辦事項
  useEffect(() => {
    loadTodos();
  }, [loadTodos, selectedCategoryId]);

  // 處理建立分類
  const handleCreateCategory = async (data: CreateCategoryRequest) => {
    const response = await categoryService.createCategory(data);
    if (response.data) {
      await loadCategories();
      setShowCreateCategory(false);
    }
  };

  // 處理刪除分類
  const handleDeleteCategory = async (categoryId: string) => {
    const response = await categoryService.deleteCategory(categoryId);
    if (!response.error) {
      await loadCategories();
      if (selectedCategoryId === categoryId) {
        setSelectedCategoryId(undefined);
      }
    }
  };

  // 處理編輯分類
  const handleEditCategory = async (category: Category) => {
    // TODO: 實作編輯分類的 UI
  };

  // 處理建立待辦事項
  const handleCreateTodo = async (data: CreateTodoRequest) => {
    const response = await todoService.createTodo(data);
    if (response.data) {
      await loadTodos();
      setShowCreateTodo(false);
    }
  };

  // 處理更新待辦事項狀態
  const handleUpdateTodoStatus = async (id: string, status: TodoStatus) => {
    // 先取得現有待辦事項
    const todoResponse = await todoService.getTodo(id);
    if (!todoResponse.data) {
      return;
    }

    // 更新狀態，保留其他欄位不變
    const response = await todoService.updateTodo(id, {
      ...todoResponse.data,
      status,
      dueDate: todoResponse.data.dueDate, // 保持原有的日期格式
    });
    
    if (response.data) {
      await loadTodos();
    }
  };

  // 處理刪除待辦事項
  const handleDeleteTodo = async (id: string) => {
    const response = await todoService.deleteTodo(id);
    if (!response.error) {
      await loadTodos();
    }
  };

  // 處理編輯待辦事項
  const handleEditTodo = async (todo: Todo) => {
    // TODO: 實作編輯待辦事項的 UI
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-blue-50 to-white dark:from-gray-900 dark:to-gray-800">
      <div className="container mx-auto px-4 py-8">
        <header className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">我的待辦事項</h1>
          <p className="mt-2 text-gray-600 dark:text-gray-300">管理您的日常任務</p>
        </header>

        <div className="grid gap-6 md:grid-cols-[300px_1fr]">
          {/* 分類側邊欄 */}
          <aside className="space-y-4">
            {showCreateCategory ? (
              <CreateCategoryForm
                onSubmit={handleCreateCategory}
                onCancel={() => setShowCreateCategory(false)}
              />
            ) : (
              <CategoryList
                categories={categories}
                selectedCategoryId={selectedCategoryId}
                onSelect={setSelectedCategoryId}
                onCreateClick={() => setShowCreateCategory(true)}
                onEditClick={handleEditCategory}
                onDeleteClick={handleDeleteCategory}
              />
            )}
          </aside>

          {/* 主要內容區 */}
          <main className="space-y-4">
            <div className="flex gap-4 mb-6">
              <input
                type="text"
                placeholder="搜尋待辦事項..."
                className="flex-1 px-4 py-2 rounded-lg border border-gray-200 dark:border-gray-700 dark:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-primary"
              />
              <button
                onClick={() => setShowCreateTodo(true)}
                className="bg-primary hover:bg-primary-hover text-white rounded-lg px-6 py-2"
              >
                新增待辦
              </button>
            </div>

            {showCreateTodo && (
              <div className="bg-white dark:bg-gray-800 rounded-lg p-4 todo-shadow">
                <CreateTodoForm
                  categories={categories}
                  onSubmit={handleCreateTodo}
                  onCancel={() => setShowCreateTodo(false)}
                />
              </div>
            )}

            {/* 待辦事項列表 */}
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4 todo-shadow">
              <div className="space-y-2">
                {todos.map((todo) => (
                  <TodoItem
                    key={todo.id}
                    todo={todo}
                    onStatusChange={handleUpdateTodoStatus}
                    onDelete={handleDeleteTodo}
                    onEdit={handleEditTodo}
                  />
                ))}
              </div>
            </div>
          </main>
        </div>
      </div>
    </div>
  );
}
