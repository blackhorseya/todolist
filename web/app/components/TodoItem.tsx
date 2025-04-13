import { Priority, Todo, TodoStatus } from "@/types/entity";

interface TodoItemProps {
  todo: Todo;
  onStatusChange?: (id: string, status: TodoStatus) => void;
  onDelete?: (id: string) => void;
  onEdit?: (todo: Todo) => void;
}

const priorityLabels: Record<Priority, { text: string; className: string }> = {
  [Priority.Low]: {
    text: "低優先級",
    className: "bg-success/10 text-success",
  },
  [Priority.Medium]: {
    text: "中等優先級",
    className: "bg-warning/10 text-warning",
  },
  [Priority.High]: {
    text: "高優先級",
    className: "bg-danger/10 text-danger",
  },
};

export default function TodoItem({ todo, onStatusChange, onDelete, onEdit }: TodoItemProps) {
  const priorityLabel = priorityLabels[todo.priority];
  
  return (
    <div className="todo-item p-4 rounded-lg bg-secondary dark:bg-gray-700 flex items-center justify-between gap-4">
      <div className="flex items-center gap-4">
        <input
          type="checkbox"
          checked={todo.status === TodoStatus.Done}
          onChange={(e) => onStatusChange?.(todo.id, e.target.checked ? TodoStatus.Done : TodoStatus.Todo)}
          className="w-5 h-5 rounded border-gray-300"
        />
        <div>
          <h3 className={`font-medium ${todo.status === TodoStatus.Done ? 'line-through text-gray-400' : ''}`}>
            {todo.title}
          </h3>
          {todo.description && (
            <p className="text-sm text-gray-600 dark:text-gray-300">{todo.description}</p>
          )}
          <div className="flex gap-2 mt-1 text-xs text-gray-500">
            <span>截止日期: {new Date(todo.dueDate).toLocaleDateString('zh-TW')}</span>
          </div>
        </div>
      </div>
      <div className="flex items-center gap-2">
        <span className={`px-2 py-1 text-xs rounded-full ${priorityLabel.className}`}>
          {priorityLabel.text}
        </span>
        <button
          onClick={() => onEdit?.(todo)}
          className="p-2 hover:bg-gray-100 dark:hover:bg-gray-600 rounded-lg"
        >
          <span className="sr-only">編輯</span>
          ✏️
        </button>
        <button
          onClick={() => onDelete?.(todo.id)}
          className="p-2 hover:bg-gray-100 dark:hover:bg-gray-600 rounded-lg"
        >
          <span className="sr-only">刪除</span>
          🗑️
        </button>
      </div>
    </div>
  );
}