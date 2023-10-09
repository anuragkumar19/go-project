-- name: CreatePost :many
INSERT INTO posts (
    title, text, image, video, link, subreddit_id, creator_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id;

-- name: VotePost :exec
INSERT INTO vote_post (post_id, user_id, down)
VALUES ($1, $2, $3)
ON CONFLICT
DO UPDATE SET down = $3
WHERE vote_post.post_id = $1 AND vote_post.user_id = $2;

-- name: GetVote :many
SELECT post_id, user_id, down FROM vote_post
WHERE post_id = $1 AND user_id = $2;

-- name: RemoveVote :exec
DELETE FROM vote_post
WHERE vote_post.post_id = $1 AND vote_post.user_id = $2;

-- name: FindPostById :many
SELECT id, creator_id, subreddit_id FROM posts
WHERE id = $1;
