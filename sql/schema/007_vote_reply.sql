-- +goose Up
CREATE TABLE vote_reply (
  user_id INT NOT NULL,
  reply_id INT NOT NULL,
  down BOOLEAN NOT NULL,
  PRIMARY KEY (user_id, reply_id),
  CONSTRAINT vote_reply_user_fk FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT vote_reply_reply_fk FOREIGN KEY (reply_id) REFERENCES replies(id) 
);

-- +goose Down
DROP TABLE vote_reply;

