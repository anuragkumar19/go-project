-- name: CreateSubreddit :many
INSERT INTO subreddit (
    title, name, creator_id
) VALUES (
    $1, $2, $3
) RETURNING id, name;

-- name: FindSubreddit :many
SELECT id FROM subreddit
WHERE name = $1;

-- name: FindSubredditById :many
SELECT id, creator_id FROM subreddit
WHERE id = $1;

-- name: UpdateSubredditTitle :exec
UPDATE subreddit 
SET title = $2
WHERE id = $1;

-- name: UpdateSubredditAvatar :exec
UPDATE subreddit 
SET avatar = $2
WHERE id = $1;

-- name: UpdateSubredditCover :exec
UPDATE subreddit 
SET cover = $2
WHERE id = $1;

-- name: UpdateSubredditName :exec
UPDATE subreddit 
SET name = $2
WHERE id = $1;

