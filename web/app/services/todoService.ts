import { CreateTodoRequest, TodoQueryParams, UpdateTodoRequest } from "@/types/api";
import { Todo } from "@/types/entity";
import { ApiResponse, ITodoService } from "@/types/service";

export class TodoService implements ITodoService {
  private baseUrl: string;

  constructor() {
    this.baseUrl = `${process.env.NEXT_PUBLIC_API_URL}/v1/todos`;
  }

  async listTodos(params?: TodoQueryParams): Promise<ApiResponse<Todo[]>> {
    try {
      const queryString = params
        ? `?${new URLSearchParams(
            Object.entries(params)
              .filter(([_, value]) => value !== undefined)
              .reduce((acc, [key, value]) => ({ ...acc, [key]: value }), {})
          ).toString()}`
        : "";

      const response = await fetch(`${this.baseUrl}${queryString}`);
      const data = await response.json();
      return { data };
    } catch (error) {
      return { error: { message: "取得待辦事項列表失敗" } };
    }
  }

  async getTodo(id: string): Promise<ApiResponse<Todo>> {
    try {
      const response = await fetch(`${this.baseUrl}/${id}`);
      const data = await response.json();
      return { data };
    } catch (error) {
      return { error: { message: "取得待辦事項失敗" } };
    }
  }

  async createTodo(req: CreateTodoRequest): Promise<ApiResponse<Todo>> {
    try {
      const response = await fetch(this.baseUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(req),
      });
      const data = await response.json();
      return { data };
    } catch (error) {
      return { error: { message: "建立待辦事項失敗" } };
    }
  }

  async updateTodo(id: string, req: UpdateTodoRequest): Promise<ApiResponse<Todo>> {
    try {
      const response = await fetch(`${this.baseUrl}/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(req),
      });
      const data = await response.json();
      return { data };
    } catch (error) {
      return { error: { message: "更新待辦事項失敗" } };
    }
  }

  async deleteTodo(id: string): Promise<ApiResponse<void>> {
    try {
      await fetch(`${this.baseUrl}/${id}`, {
        method: "DELETE",
      });
      return {};
    } catch (error) {
      return { error: { message: "刪除待辦事項失敗" } };
    }
  }
}