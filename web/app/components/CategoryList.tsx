import { Category } from "@/types/entity";

interface CategoryListProps {
  categories: Category[];
  selectedCategoryId?: string;
  onSelect: (categoryId?: string) => void;
  onCreateClick: () => void;
  onEditClick: (category: Category) => void;
  onDeleteClick: (categoryId: string) => void;
}

export default function CategoryList({
  categories,
  selectedCategoryId,
  onSelect,
  onCreateClick,
  onEditClick,
  onDeleteClick,
}: CategoryListProps) {
  return (
    <div className="bg-white dark:bg-gray-800 rounded-lg p-4 todo-shadow">
      <h2 className="text-lg font-semibold mb-4">分類</h2>
      <button
        onClick={onCreateClick}
        className="w-full bg-primary hover:bg-primary-hover text-white rounded-lg px-4 py-2 mb-4"
      >
        新增分類
      </button>
      <div className="space-y-2">
        <button
          onClick={() => onSelect(undefined)}
          className={`w-full text-left px-4 py-2 rounded-lg transition-colors ${
            !selectedCategoryId
              ? "bg-primary/10 text-primary"
              : "bg-secondary hover:bg-secondary-hover"
          }`}
        >
          所有事項
        </button>
        {categories.map((category) => (
          <div
            key={category.id}
            className={`group flex items-center gap-2 rounded-lg ${
              selectedCategoryId === category.id
                ? "bg-primary/10"
                : "bg-secondary hover:bg-secondary-hover"
            }`}
          >
            <button
              onClick={() => onSelect(category.id)}
              className="flex-1 text-left px-4 py-2"
            >
              {category.name}
            </button>
            <div className="hidden group-hover:flex items-center gap-1 px-2">
              <button
                onClick={() => onEditClick(category)}
                className="p-1 hover:bg-gray-200 dark:hover:bg-gray-600 rounded"
              >
                <span className="sr-only">編輯</span>
                ✏️
              </button>
              <button
                onClick={() => onDeleteClick(category.id)}
                className="p-1 hover:bg-gray-200 dark:hover:bg-gray-600 rounded"
              >
                <span className="sr-only">刪除</span>
                🗑️
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}