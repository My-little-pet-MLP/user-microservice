  CREATE TABLE IF NOT EXISTS users (
            ID VARCHAR(255) PRIMARY KEY,
            fullname VARCHAR(255),
            imageUrl VARCHAR(255),
            email VARCHAR(255),
            phone VARCHAR(12),
            created_at TIMESTAMP
        );