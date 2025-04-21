import { useEffect, useState } from 'react';
import { Pomodoro, PomodoroStatus, Todo } from '@/types/entity';

interface PomodoroProps {
  todo: Todo;
  onPomodoroComplete?: () => void;
}

export default function PomodoroTimer({ todo, onPomodoroComplete }: PomodoroProps) {
  const [pomodoro, setPomodoro] = useState<Pomodoro>({
    todoId: todo.id,
    status: PomodoroStatus.Idle,
    timeLeft: 25 * 60, // 25 分鐘
    totalTime: 25 * 60,
    breakTime: 5 * 60, // 5 分鐘休息
  });

  useEffect(() => {
    let timer: NodeJS.Timeout;
    
    if (pomodoro.status === PomodoroStatus.Running) {
      timer = setInterval(() => {
        setPomodoro(prev => {
          if (prev.timeLeft <= 0) {
            if (prev.status === PomodoroStatus.Running) {
              // 工作時段結束，進入休息時間
              return {
                ...prev,
                status: PomodoroStatus.Break,
                timeLeft: prev.breakTime,
              };
            } else if (prev.status === PomodoroStatus.Break) {
              // 休息時間結束，重置番茄鐘
              onPomodoroComplete?.();
              return {
                ...prev,
                status: PomodoroStatus.Idle,
                timeLeft: prev.totalTime,
              };
            }
          }
          return {
            ...prev,
            timeLeft: prev.timeLeft - 1,
          };
        });
      }, 1000);
    }

    return () => clearInterval(timer);
  }, [pomodoro.status, onPomodoroComplete]);

  const formatTime = (seconds: number) => {
    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;
    return `${minutes.toString().padStart(2, '0')}:${remainingSeconds.toString().padStart(2, '0')}`;
  };

  const handleStart = () => {
    setPomodoro(prev => ({
      ...prev,
      status: PomodoroStatus.Running,
    }));
  };

  const handlePause = () => {
    setPomodoro(prev => ({
      ...prev,
      status: PomodoroStatus.Paused,
    }));
  };

  const handleResume = () => {
    setPomodoro(prev => ({
      ...prev,
      status: PomodoroStatus.Running,
    }));
  };

  const handleReset = () => {
    setPomodoro(prev => ({
      ...prev,
      status: PomodoroStatus.Idle,
      timeLeft: prev.totalTime,
    }));
  };

  return (
    <div className="pomodoro-timer bg-white dark:bg-gray-800 p-4 rounded-lg shadow-sm">
      <div className="text-center mb-4">
        <h3 className="text-lg font-semibold mb-2">
          {pomodoro.status === PomodoroStatus.Break ? '休息時間' : todo.title}
        </h3>
        <div className="text-3xl font-mono">
          {formatTime(pomodoro.timeLeft)}
        </div>
        <div className="text-sm text-gray-500 mt-1">
          {pomodoro.status === PomodoroStatus.Break ? '休息中' : 
           pomodoro.status === PomodoroStatus.Running ? '專注中' :
           pomodoro.status === PomodoroStatus.Paused ? '已暫停' : '準備開始'}
        </div>
      </div>
      <div className="flex justify-center gap-2">
        {pomodoro.status === PomodoroStatus.Idle && (
          <button
            onClick={handleStart}
            className="px-4 py-2 bg-primary text-white rounded-lg hover:bg-primary/90"
          >
            開始
          </button>
        )}
        {pomodoro.status === PomodoroStatus.Running && (
          <button
            onClick={handlePause}
            className="px-4 py-2 bg-warning text-white rounded-lg hover:bg-warning/90"
          >
            暫停
          </button>
        )}
        {pomodoro.status === PomodoroStatus.Paused && (
          <>
            <button
              onClick={handleResume}
              className="px-4 py-2 bg-primary text-white rounded-lg hover:bg-primary/90"
            >
              繼續
            </button>
            <button
              onClick={handleReset}
              className="px-4 py-2 bg-danger text-white rounded-lg hover:bg-danger/90"
            >
              重置
            </button>
          </>
        )}
      </div>
    </div>
  );
}