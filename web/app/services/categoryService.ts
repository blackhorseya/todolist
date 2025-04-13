import { CreateCategoryRequest, UpdateCategoryRequest } from "@/types/api";
import { Category } from "@/types/entity";
import { ApiResponse, ICategoryService } from "@/types/service";

export class CategoryService implements ICategoryService {
  private baseUrl: string;

  constructor() {
    this.baseUrl = `${process.env.NEXT_PUBLIC_API_URL}/v1/categories`;
  }

  async listCategories(): Promise<ApiResponse<Category[]>> {
    try {
      const response = await fetch(this.baseUrl);
      const data = await response.json();
      return { data };
    } catch (error) {
      return { error: { message: "取得分類列表失敗" } };
    }
  }

  async getCategory(id: string): Promise<ApiResponse<Category>> {
    try {
      const response = await fetch(`${this.baseUrl}/${id}`);
      const data = await response.json();
      return { data };
    } catch (error) {
      return { error: { message: "取得分類失敗" } };
    }
  }

  async createCategory(req: CreateCategoryRequest): Promise<ApiResponse<Category>> {
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
      return { error: { message: "建立分類失敗" } };
    }
  }

  async updateCategory(
    id: string,
    req: UpdateCategoryRequest
  ): Promise<ApiResponse<Category>> {
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
      return { error: { message: "更新分類失敗" } };
    }
  }

  async deleteCategory(id: string): Promise<ApiResponse<void>> {
    try {
      await fetch(`${this.baseUrl}/${id}`, {
        method: "DELETE",
      });
      return {};
    } catch (error) {
      return { error: { message: "刪除分類失敗" } };
    }
  }
}