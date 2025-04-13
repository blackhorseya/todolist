import Image from "next/image";

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-blue-50 to-white dark:from-gray-900 dark:to-gray-800">
      <div className="container mx-auto px-4 py-8">
        <header className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">我的待辦事項</h1>
          <p className="mt-2 text-gray-600 dark:text-gray-300">管理您的日常任務</p>
        </header>

        <div className="grid gap-6 md:grid-cols-[300px_1fr]">
          {/* 分類側邊欄 */}
          <aside className="space-y-4">
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4 todo-shadow">
              <h2 className="text-lg font-semibold mb-4">分類</h2>
              <button className="w-full bg-primary hover:bg-primary-hover text-white rounded-lg px-4 py-2 mb-4">
                新增分類
              </button>
              <div className="space-y-2">
                <button className="w-full text-left px-4 py-2 rounded-lg bg-secondary hover:bg-secondary-hover transition-colors">
                  所有事項
                </button>
                {/* 這裡將動態載入分類列表 */}
              </div>
            </div>
          </aside>

          {/* 主要內容區 */}
          <main className="space-y-4">
            <div className="flex gap-4 mb-6">
              <input
                type="text"
                placeholder="搜尋待辦事項..."
                className="flex-1 px-4 py-2 rounded-lg border border-gray-200 dark:border-gray-700 dark:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-primary"
              />
              <button className="bg-primary hover:bg-primary-hover text-white rounded-lg px-6 py-2">
                新增待辦
              </button>
            </div>

            {/* 待辦事項列表 */}
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4 todo-shadow">
              <div className="space-y-2">
                {/* 待辦事項卡片範例 */}
                <div className="todo-item p-4 rounded-lg bg-secondary dark:bg-gray-700 flex items-center justify-between gap-4">
                  <div className="flex items-center gap-4">
                    <input type="checkbox" className="w-5 h-5 rounded border-gray-300" />
                    <div>
                      <h3 className="font-medium">完成待辦事項清單設計</h3>
                      <p className="text-sm text-gray-600 dark:text-gray-300">前端介面開發</p>
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <span className="px-2 py-1 text-xs rounded-full bg-warning/10 text-warning">
                      中等優先級
                    </span>
                    <button className="p-2 hover:bg-gray-100 dark:hover:bg-gray-600 rounded-lg">
                      <span className="sr-only">編輯</span>
                      ✏️
                    </button>
                    <button className="p-2 hover:bg-gray-100 dark:hover:bg-gray-600 rounded-lg">
                      <span className="sr-only">刪除</span>
                      🗑️
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </main>
        </div>
      </div>
    </div>
  );
}
