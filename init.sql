\c postgres;

CREATE TABLE IF NOT EXISTS Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password_hash VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS Tasks (
    task_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    priority INT NOT NULL CHECK(priority >= 0 AND priority <= 2),
    due_date timestamp,
    completed BOOLEAN NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);