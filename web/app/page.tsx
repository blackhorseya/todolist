"use client";

import { useState } from "react";
import { Category, Todo, Priority, TodoStatus } from "@/types/entity";
import CategoryList from "./components/CategoryList";
import CreateCategoryForm from "./components/CreateCategoryForm";
import CreateTodoForm from "./components/CreateTodoForm";
import TodoItem from "./components/TodoItem";

export default function Home() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [selectedCategoryId, setSelectedCategoryId] = useState<string>();
  const [showCreateCategory, setShowCreateCategory] = useState(false);
  const [showCreateTodo, setShowCreateTodo] = useState(false);

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
                onSubmit={(data) => {
                  // TODO: 實作建立分類
                  setShowCreateCategory(false);
                }}
                onCancel={() => setShowCreateCategory(false)}
              />
            ) : (
              <CategoryList
                categories={categories}
                selectedCategoryId={selectedCategoryId}
                onSelect={setSelectedCategoryId}
                onCreateClick={() => setShowCreateCategory(true)}
                onEditClick={(category) => {
                  // TODO: 實作編輯分類
                }}
                onDeleteClick={(categoryId) => {
                  // TODO: 實作刪除分類
                }}
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
                  onSubmit={(data) => {
                    // TODO: 實作建立待辦事項
                    setShowCreateTodo(false);
                  }}
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
                    onStatusChange={(id, status) => {
                      // TODO: 實作狀態更新
                    }}
                    onDelete={(id) => {
                      // TODO: 實作刪除
                    }}
                    onEdit={(todo) => {
                      // TODO: 實作編輯
                    }}
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
