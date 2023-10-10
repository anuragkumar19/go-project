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
    replies.parent_reply_id,
    replies.post_id,
    replies.creator_id,
    replies.created_at,
    users.username AS creator_username,
    users.avatar AS creator_avatar,
    users.name AS creator_name,
    COALESCE(r.replies_count, 0) AS replies_count,
    COALESCE(up_votes.up_vote_count, 0) AS up_vote_count,
    COALESCE(down_votes.down_vote_count, 0) AS down_vote_count
FROM
    replies
JOIN
    users ON replies.creator_id = users.id
LEFT JOIN (
    SELECT parent_reply_id, COUNT(id) AS replies_count
    FROM replies as r
    GROUP BY parent_reply_id
) AS r ON replies.id = r.parent_reply_id
LEFT JOIN (
    SELECT reply_id, COUNT(user_id) AS up_vote_count
    FROM vote_reply
    WHERE down = FALSE
    GROUP BY reply_id
) AS up_votes ON replies.id = up_votes.reply_id
LEFT JOIN (
    SELECT reply_id, COUNT(user_id) AS down_vote_count
    FROM vote_reply
    WHERE down = TRUE
    GROUP BY reply_id
) AS down_votes ON replies.id = down_votes.reply_id
WHERE
    replies.id = $1;

-- name: GetUserReplyPublic :many
SELECT
    replies.id,
    replies.content,
    replies.parent_reply_id,
    replies.post_id,
    replies.creator_id,
    replies.created_at,
    users.username AS creator_username,
    users.avatar AS creator_avatar,
    users.name AS creator_name,
    COALESCE(r.replies_count, 0) AS replies_count,
    COALESCE(up_votes.up_vote_count, 0) AS up_vote_count,
    COALESCE(down_votes.down_vote_count, 0) AS down_vote_count
FROM
    replies
JOIN
    users ON replies.creator_id = users.id
LEFT JOIN (
    SELECT parent_reply_id, COUNT(id) AS replies_count
    FROM replies as r
    GROUP BY parent_reply_id
) AS r ON replies.id = r.parent_reply_id
LEFT JOIN (
    SELECT reply_id, COUNT(user_id) AS up_vote_count
    FROM vote_reply
    WHERE down = FALSE
    GROUP BY reply_id
) AS up_votes ON replies.id = up_votes.reply_id
LEFT JOIN (
    SELECT reply_id, COUNT(user_id) AS down_vote_count
    FROM vote_reply
    WHERE down = TRUE
    GROUP BY reply_id
) AS down_votes ON replies.id = down_votes.reply_id
WHERE
    replies.creator_id = $1
ORDER BY
    replies.created_at DESC
LIMIT
    $2
OFFSET
    $3;

-- name: GetPostReplyPublic :many
SELECT
    replies.id,
    replies.content,
    replies.parent_reply_id,
    replies.post_id,
    replies.creator_id,
    replies.created_at,
    users.username AS creator_username,
    users.avatar AS creator_avatar,
    users.name AS creator_name,
    COALESCE(r.replies_count, 0) AS replies_count,
    COALESCE(up_votes.up_vote_count, 0) AS up_vote_count,
    COALESCE(down_votes.down_vote_count, 0) AS down_vote_count
FROM
    replies
JOIN
    users ON replies.creator_id = users.id
LEFT JOIN (
    SELECT parent_reply_id, COUNT(id) AS replies_count
    FROM replies as r
    GROUP BY parent_reply_id
) AS r ON replies.id = r.parent_reply_id
LEFT JOIN (
    SELECT reply_id, COUNT(user_id) AS up_vote_count
    FROM vote_reply
    WHERE down = FALSE
    GROUP BY reply_id
) AS up_votes ON replies.id = up_votes.reply_id
LEFT JOIN (
    SELECT reply_id, COUNT(user_id) AS down_vote_count
    FROM vote_reply
    WHERE down = TRUE
    GROUP BY reply_id
) AS down_votes ON replies.id = down_votes.reply_id
WHERE
    replies.post_id = $1
ORDER BY
    replies.created_at DESC
LIMIT
    $2
OFFSET
    $3;

-- name: GetReplyReplies :many
SELECT
    replies.id,
    replies.content,
    replies.parent_reply_id,
    replies.post_id,
    replies.creator_id,
    replies.created_at,
    users.username AS creator_username,
    users.avatar AS creator_avatar,
    users.name AS creator_name,
    COALESCE(r.replies_count, 0) AS replies_count,
    COALESCE(up_votes.up_vote_count, 0) AS up_vote_count,
    COALESCE(down_votes.down_vote_count, 0) AS down_vote_count
FROM
    replies
JOIN
    users ON replies.creator_id = users.id
LEFT JOIN (
    SELECT parent_reply_id, COUNT(id) AS replies_count
    FROM replies as r
    GROUP BY parent_reply_id
) AS r ON replies.id = r.parent_reply_id
LEFT JOIN (
    SELECT reply_id, COUNT(user_id) AS up_vote_count
    FROM vote_reply
    WHERE down = FALSE
    GROUP BY reply_id
) AS up_votes ON replies.id = up_votes.reply_id
LEFT JOIN (
    SELECT reply_id, COUNT(user_id) AS down_vote_count
    FROM vote_reply
    WHERE down = TRUE
    GROUP BY reply_id
) AS down_votes ON replies.id = down_votes.reply_id
WHERE
    replies.parent_reply_id = $1
ORDER BY
    replies.created_at DESC
LIMIT
    $2
OFFSET
    $3;




