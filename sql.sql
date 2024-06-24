CREATE TABLE categories (
    uuid TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at DATETIME NOT NULL
) AUTO_INCREMENT=1;

CREATE TABLE topics (
    uuid TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY(category_id) REFERENCES categories(id)
) AUTO_INCREMENT;

CREATE TABLE threads (
    uuid TEXT NOT NULL PRIMARY KEY,
    topic_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY(topic_id) REFERENCES topics(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
) AUTO_INCREMENT;

CREATE TABLE posts (
    uuid TEXT NOT NULL PRIMARY KEY,
    body TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    thread_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(thread_id) REFERENCES threads(id)
) AUTO_INCREMENT;

CREATE TABLE users (
    uuid VARCHAR(255) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL
) AUTO_INCREMEN;

CREATE TABLE sessions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    uuid VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
) AUTO_INCREMENT=1;