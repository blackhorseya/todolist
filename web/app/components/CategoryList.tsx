import { Category } from "@/types/entity";

interface CategoryListProps {
  categories: Category[];
  selectedCategoryId?: string;
  onSelect: (categoryId?: string) => void;
  onCreateClick: () => void;
  onEditClick: (category: Category) => void;
  onDeleteClick: (categoryId: string) => void;
}

interface CategoryItem {
  id: string;
  name: string;
  onClick: () => void;
  isSelected: boolean;
  category?: Category;
}

export default function CategoryList({
  categories,
  selectedCategoryId,
  onSelect,
  onCreateClick,
  onEditClick,
  onDeleteClick,
}: CategoryListProps) {
  const categoryItems: CategoryItem[] = [
    {
      id: "all",
      name: "所有事項",
      onClick: () => onSelect(undefined),
      isSelected: !selectedCategoryId,
    },
    ...categories.map((category) => ({
      id: category.id,
      name: category.name,
      onClick: () => onSelect(category.id),
      isSelected: selectedCategoryId === category.id,
      category,
    })),
  ];

  const handleEditClick = (item: CategoryItem) => {
    if (item.category) {
      onEditClick(item.category);
    }
  };

  const handleDeleteClick = (item: CategoryItem) => {
    if (item.category) {
      onDeleteClick(item.category.id);
    }
  };

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
        {categoryItems.map((item) => (
          <div
            key={item.id}
            className={`group flex items-center gap-2 rounded-lg ${
              item.isSelected
                ? "bg-primary/10"
                : "bg-secondary hover:bg-secondary-hover"
            }`}
          >
            <button
              onClick={item.onClick}
              className="flex-1 text-left px-4 py-2"
            >
              {item.name}
            </button>
            {item.category && (
              <div className="hidden group-hover:flex items-center gap-1 px-2">
                <button
                  onClick={() => handleEditClick(item)}
                  className="p-1 hover:bg-gray-200 dark:hover:bg-gray-600 rounded"
                >
                  <span className="sr-only">編輯</span>
                  ✏️
                </button>
                <button
                  onClick={() => handleDeleteClick(item)}
                  className="p-1 hover:bg-gray-200 dark:hover:bg-gray-600 rounded"
                >
                  <span className="sr-only">刪除</span>
                  🗑️
                </button>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}