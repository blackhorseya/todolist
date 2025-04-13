import { ApiError, CreateCategoryRequest, CreateTodoRequest, TodoQueryParams, UpdateCategoryRequest, UpdateTodoRequest } from './api';
import { Category, Todo } from './entity';

// API 回應介面
export interface ApiResponse<T> {
  data?: T;
  error?: ApiError;
}

// 分類服務介面
export interface ICategoryService {
  listCategories(): Promise<ApiResponse<Category[]>>;
  getCategory(id: string): Promise<ApiResponse<Category>>;
  createCategory(req: CreateCategoryRequest): Promise<ApiResponse<Category>>;
  updateCategory(id: string, req: UpdateCategoryRequest): Promise<ApiResponse<Category>>;
  deleteCategory(id: string): Promise<ApiResponse<void>>;
}

// 待辦事項服務介面
export interface ITodoService {
  listTodos(params?: TodoQueryParams): Promise<ApiResponse<Todo[]>>;
  getTodo(id: string): Promise<ApiResponse<Todo>>;
  createTodo(req: CreateTodoRequest): Promise<ApiResponse<Todo>>;
  updateTodo(id: string, req: UpdateTodoRequest): Promise<ApiResponse<Todo>>;
  deleteTodo(id: string): Promise<ApiResponse<void>>;
}