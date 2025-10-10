-- Create todo_history table (sqlite)
CREATE TABLE IF NOT EXISTS todo_history (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
    todo_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    completed_at DATE DEFAULT NOW(),
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    notes TEXT,
    FOREIGN KEY (todo_id) REFERENCES todos (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_task_completions_todo_id ON todo_history(todo_id);
CREATE INDEX IF NOT EXISTS idx_task_completions_user_id ON todo_history(user_id);
CREATE INDEX IF NOT EXISTS idx_task_completions_completed_at ON todo_history(completed_at);

-- Add constraint to prevent duplicate completions on same day
CREATE UNIQUE INDEX IF NOT EXISTS idx_task_completions_unique_daily 
    ON todo_history(todo_id, DATE(completed_at));