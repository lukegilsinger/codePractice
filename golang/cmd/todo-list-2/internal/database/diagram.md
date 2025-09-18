erDiagram
    users {
        INTEGER id PK
        TEXT username UNIQUE
        TEXT email UNIQUE
        TEXT password
        DATETIME created_at
        DATETIME updated_at
    }

    categories {
        INTEGER id PK
        INTEGER user_id FK
        TEXT name UNIQUE
        TEXT description
        TEXT color
        DATETIME created_at
        DATETIME updated_at
    }

    todos {
        INTEGER id PK
        INTEGER user_id FK
        TEXT title
        TEXT description
        BOOLEAN completed
        INTEGER category_id FK
        DATETIME created_at
        DATETIME updated_at
    }

    users ||--o{ categories : "user_id"
    users ||--o{ todos : "user_id"
    categories ||--o{ todos : "category_id"
