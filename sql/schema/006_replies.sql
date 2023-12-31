-- +goose Up
CREATE TABLE replies (
    id SERIAL PRIMARY KEY,
    creator_id INT NOT NULL,
    post_id INT,
    parent_reply_id INT,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
    updated_at TIMESTAMP DEFAULT current_timestamp NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (creator_id) REFERENCES users(id),
    CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_parent_reply FOREIGN KEY (parent_reply_id) REFERENCES replies(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE replies;