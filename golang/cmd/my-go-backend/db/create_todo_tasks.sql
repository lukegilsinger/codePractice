CREATE TABLE IF NOT EXISTS todo_tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL,
    due_date DATE,
    cadence TEXT NOT NULL CHECK (cadence IN ('daily', 'weekly', 'monthly', 'yearly'))
);