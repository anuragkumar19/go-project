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
ON CONFLICT (reply_id, user_id)
DO UPDATE SET down = $3
WHERE vote_reply.reply_id = $1 AND vote_reply.user_id = $2;

-- name: GetReplyVote :many
SELECT reply_id, user_id, down FROM vote_reply
WHERE reply_id = $1 AND user_id = $2;

-- name: RemoveReplyVote :exec
DELETE FROM vote_reply
WHERE vote_reply.reply_id = $1 AND vote_reply.user_id = $2;

-- name: GetReplyByIdPublic :many
SELECT
    replies.id,
    replies.content,
    replies.creator_id,
    replies.created_at,
    replies.post_id,
    replies.parent_reply_id,
    users.username AS creator_username,
    users.avatar AS creator_avatar,
    users.name AS creator_name,
    COUNT(r.id) AS replies_count
FROM
    replies
JOIN
    users ON replies.creator_id = users.id
JOIN 
    replies AS r on r.parent_reply_id = replies.id
WHERE
    replies.id = $1;
