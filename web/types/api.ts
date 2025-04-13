import { Priority, TodoStatus } from './entity';

// 分類相關請求
export interface CreateCategoryRequest {
  name: string;
  description?: string;
}

export interface UpdateCategoryRequest {
  name?: string;
  description?: string;
}

// 待辦事項相關請求
export interface CreateTodoRequest {
  title: string;
  description?: string;
  categoryId: string;
  priority: Priority;
  dueDate: string;
}

export interface UpdateTodoRequest {
  title?: string;
  description?: string;
  categoryId?: string;
  status?: TodoStatus;
  priority?: Priority;
  dueDate?: string;
}

// API 查詢參數
export interface TodoQueryParams {
  categoryId?: string;
  status?: TodoStatus;
  priority?: Priority;
}

// API 錯誤回應
export interface ApiError {
  [key: string]: string;
}