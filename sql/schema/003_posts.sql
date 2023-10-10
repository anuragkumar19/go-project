-- +goose Up
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    text TEXT DEFAULT '' NOT NULL,
    image TEXT DEFAULT '' NOT NULL,
    video TEXT DEFAULT '' NOT NULL,
    link TEXT DEFAULT '' NOT NULL,
    subreddit_id INT NOT NULL,
    creator_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
    updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
    CONSTRAINT creator FOREIGN KEY(creator_id) REFERENCES users(id),
    CONSTRAINT subreddit FOREIGN KEY(subreddit_id) REFERENCES subreddit(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;