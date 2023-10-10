-- name: CreatePost :many
INSERT INTO posts (
    title, text, image, video, link, subreddit_id, creator_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id;

-- name: VotePost :exec
INSERT INTO vote_post (post_id, user_id, down)
VALUES ($1, $2, $3)
ON CONFLICT (post_id, user_id)
DO UPDATE SET down = $3
WHERE vote_post.post_id = $1 AND vote_post.user_id = $2;

-- name: GetPostVote :many
SELECT post_id, user_id, down FROM vote_post
WHERE post_id = $1 AND user_id = $2;

-- name: RemovePostVote :exec
DELETE FROM vote_post
WHERE vote_post.post_id = $1 AND vote_post.user_id = $2;

-- name: FindPostById :many
SELECT id, creator_id, subreddit_id FROM posts
WHERE id = $1;

-- name: GetPostByIDPublic :many
SELECT
    posts.id,
    posts.title,
    posts.text,
    posts.image,
    posts.video,
    posts.link,
    posts.subreddit_id,
    posts.creator_id,
    posts.created_at,
    users.username AS creator_username,
    users.avatar AS creator_avatar,
    users.name AS creator_name,
    subreddit.name AS subreddit_name,
    subreddit.avatar AS subreddit_avatar,
    subreddit.is_verified AS subreddit_is_verified,
    subreddit.title AS subreddit_title,
    count(replies.id) AS replies_count
FROM
    posts
JOIN
    users ON posts.creator_id = users.id
JOIN
    subreddit ON posts.subreddit_id = subreddit.id
JOIN
    replies ON posts.id = replies.post_id
WHERE
    posts.id = $1;
