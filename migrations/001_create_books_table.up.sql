CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       author VARCHAR(255) NOT NULL,
                       publisher VARCHAR(255),
                       isbn VARCHAR(20) UNIQUE NOT NULL,
                       genre VARCHAR(100),
                       stock INT NOT NULL DEFAULT 0,
                       total_pages INT,
                       language VARCHAR(50),
                       published_at TIMESTAMP,
                       description TEXT,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
