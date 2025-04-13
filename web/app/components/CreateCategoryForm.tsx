import { CreateCategoryRequest } from "@/types/api";
import { FormEvent, useState } from "react";

interface CreateCategoryFormProps {
  initialData?: CreateCategoryRequest;
  onSubmit: (data: CreateCategoryRequest) => void;
  onCancel: () => void;
}

export default function CreateCategoryForm({
  initialData,
  onSubmit,
  onCancel,
}: CreateCategoryFormProps) {
  const [formData, setFormData] = useState<CreateCategoryRequest>(
    initialData || {
      name: "",
      description: "",
    }
  );

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    onSubmit(formData);
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label htmlFor="name" className="block text-sm font-medium mb-1">
          名稱
        </label>
        <input
          type="text"
          id="name"
          required
          value={formData.name}
          onChange={(e) =>
            setFormData((prev) => ({ ...prev, name: e.target.value }))
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
          {initialData ? "更新" : "建立"}
        </button>
      </div>
    </form>
  );
}