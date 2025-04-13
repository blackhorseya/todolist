// 待辦事項狀態列舉
export enum TodoStatus {
  Todo = 1,
  InProgress = 2,
  Done = 3,
}

// 優先級別列舉
export enum Priority {
  Low = 1,
  Medium = 2,
  High = 3,
}

// 分類實體介面
export interface Category {
  id: string;
  name: string;
  description?: string;
  createdAt: string;
  updatedAt: string;
}

// 待辦事項實體介面
export interface Todo {
  id: string;
  title: string;
  description?: string;
  status: TodoStatus;
  priority: Priority;
  categoryID: string;
  dueDate: string;
  createdAt: string;
  updatedAt: string;
}