-- +goose Up
CREATE TABLE subreddit (
    id SERIAL PRIMARY KEY,
    title TEXT DEFAULT '' NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    avatar TEXT DEFAULT 'https://api.dicebear.com/7.x/bottts/png' NOT NULL,
    cover TEXT DEFAULT '' NOT NULL,
    creator_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
    updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL,

    CONSTRAINT creator FOREIGN KEY(creator_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE subreddit;