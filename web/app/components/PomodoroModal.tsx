import { Todo } from '@/types/entity';
import PomodoroTimer from './Pomodoro';

interface PomodoroModalProps {
  todo: Todo | null;
  onClose: () => void;
}

export default function PomodoroModal({ todo, onClose }: PomodoroModalProps) {
  if (!todo) return null;

  return (
    <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div className="bg-white dark:bg-gray-800 rounded-lg p-4 max-w-md w-full mx-4">
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold">番茄鐘</h2>
          <button
            onClick={onClose}
            className="p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg"
          >
            ✕
          </button>
        </div>
        <PomodoroTimer todo={todo} onPomodoroComplete={onClose} />
      </div>
    </div>
  );
}