-- +goose Up
CREATE TABLE vote_post (
  user_id INT NOT NULL,
  post_id INT NOT NULL,
  down BOOLEAN NOT NULL,
  PRIMARY KEY (user_id, post_id),
  CONSTRAINT vote_post_user_fk FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT vote_post_post_fk FOREIGN KEY (post_id) REFERENCES posts(id)
);

-- +goose Down
DROP TABLE vote_post;