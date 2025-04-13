import { Category, Priority } from "@/types/entity";
import { CreateTodoRequest } from "@/types/api";
import { FormEvent, useState } from "react";

interface CreateTodoFormProps {
  categories: Category[];
  onSubmit: (data: CreateTodoRequest) => void;
  onCancel: () => void;
}

export default function CreateTodoForm({
  categories,
  onSubmit,
  onCancel,
}: CreateTodoFormProps) {
  const [formData, setFormData] = useState<CreateTodoRequest>({
    title: "",
    description: "",
    categoryId: categories[0]?.id || "",
    priority: Priority.Medium,
    dueDate: new Date().toISOString().split("T")[0],
  });

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    // 將日期轉換為 RFC3339 格式
    const rfc3339Date = new Date(formData.dueDate).toISOString();
    onSubmit({
      ...formData,
      dueDate: rfc3339Date,
    });
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label htmlFor="title" className="block text-sm font-medium mb-1">
          標題
        </label>
        <input
          type="text"
          id="title"
          required
          value={formData.title}
          onChange={(e) =>
            setFormData((prev) => ({ ...prev, title: e.target.value }))
          }
          className="w-full px-4 py-2 rounded-lg border border-gray-200 dark:border-gray-700 dark:bg-gray-800"
        />
      </div>

      <div>
        <label htmlFor="description" className="block text-sm font-medium mb-1">
          描述
        </label>
        <textarea
          id="description"
          value={formData.description}
          onChange={(e) =>
            setFormData((prev) => ({ ...prev, description: e.target.value }))
          }
          className="w-full px-4 py-2 rounded-lg border border-gray-200 dark:border-gray-700 dark:bg-gray-800"
          rows={3}
        />
      </div>

      <div className="grid grid-cols-2 gap-4">
        <div>
          <label htmlFor="category" className="block text-sm font-medium mb-1">
            分類
          </label>
          <select
            id="category"
            required
            value={formData.categoryId}
            onChange={(e) =>
              setFormData((prev) => ({ ...prev, categoryId: e.target.value }))
            }
            className="w-full px-4 py-2 rounded-lg border border-gray-200 dark:border-gray-700 dark:bg-gray-800"
          >
            {categories.map((category) => (
              <option key={category.id} value={category.id}>
                {category.name}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label htmlFor="priority" className="block text-sm font-medium mb-1">
            優先級別
          </label>
          <select
            id="priority"
            required
            value={formData.priority}
            onChange={(e) =>
              setFormData((prev) => ({
                ...prev,
                priority: Number(e.target.value) as Priority,
              }))
            }
            className="w-full px-4 py-2 rounded-lg border border-gray-200 dark:border-gray-700 dark:bg-gray-800"
          >
            <option value={Priority.Low}>低</option>
            <option value={Priority.Medium}>中</option>
            <option value={Priority.High}>高</option>
          </select>
        </div>
      </div>

      <div>
        <label htmlFor="dueDate" className="block text-sm font-medium mb-1">
          截止日期
        </label>
        <input
          type="date"
          id="dueDate"
          required
          value={formData.dueDate}
          onChange={(e) =>
            setFormData((prev) => ({ ...prev, dueDate: e.target.value }))
          }
          className="w-full px-4 py-2 rounded-lg border border-gray-200 dark:border-gray-700 dark:bg-gray-800"
        />
      </div>

      <div className="flex justify-end gap-4 mt-6">
        <button
          type="button"
          onClick={onCancel}
          className="px-4 py-2 rounded-lg bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600"
        >
          取消
        </button>
        <button
          type="submit"
          className="px-4 py-2 rounded-lg bg-primary hover:bg-primary-hover text-white"
        >
          建立
        </button>
      </div>
    </form>
  );
}