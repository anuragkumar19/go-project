-- +goose Up
CREATE TABLE user_subreddit_join (
  user_id INT NOT NULL,
  subreddit_id INT NOT NULL,
  PRIMARY KEY (user_id, subreddit_id),
  CONSTRAINT user_subreddit_join_user_fk FOREIGN KEY (user_id) REFERENCES users(id), --
  CONSTRAINT user_subreddit_join_subreddit_fk FOREIGN KEY (subreddit_id) REFERENCES subreddit(id) 
);

-- +goose Down
DROP TABLE user_subreddit_join;
