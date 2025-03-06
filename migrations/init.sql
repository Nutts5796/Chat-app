CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    chat_id INT REFERENCES chats(id),
    user_id INT REFERENCES users(id),
    content TEXT NOT NULL
);
