CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       author TEXT NOT NULL

);

CREATE TABLE user_books (
                            user_id INT REFERENCES users(id) ON DELETE CASCADE,
                            book_id INT REFERENCES books(id) ON DELETE CASCADE,
                            PRIMARY KEY (user_id, book_id)
);
