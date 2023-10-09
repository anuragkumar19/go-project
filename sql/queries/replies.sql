-- name: CreateReply :many
INSERT INTO replies (creator_id, post_id, parent_reply_id, content) 
VALUES($1, $2, $3, $4) RETURNING id;

-- name: FindReplyById :many
SELECT id, creator_id, post_id, parent_reply_id FROM replies 
WHERE id = $1;

-- name: DeleteReply :exec
DELETE FROM replies WHERE id = $1;

-- name: VoteReply :exec
INSERT INTO vote_reply (reply_id, user_id, down)
VALUES ($1, $2, $3)
ON CONFLICT
DO UPDATE SET down = $3
WHERE vote_reply.reply_id = $1 AND vote_reply.user_id = $2;

-- name: GetReplyVote :many
SELECT reply_id, user_id, down FROM vote_reply
WHERE reply_id = $1 AND user_id = $2;

-- name: RemoveReplyVote :exec
DELETE FROM vote_reply
WHERE vote_reply.reply_id = $1 AND vote_reply.user_id = $2;
