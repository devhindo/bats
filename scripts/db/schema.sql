-- name: db
CREATE DATABASE IF NOT EXISTS `bats`;

--name: use
USE `bats`;

--name: users
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    country VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--name: addtestusers
INSERT INTO users (username, email, password) VALUES
('test', 'test@mail.com', 'test'),
('test2', 'test2@mail.com', 'test2');

--name: posts
CREATE TABLE IF NOT EXISTS posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    author_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    content TEXT NOT NULL,
    likes INT DEFAULT 0,
    comments_number INT DEFAULT 0,
    FOREIGN KEY (author_id) REFERENCES users(id)
);