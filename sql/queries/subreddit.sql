-- name: CreateSubreddit :many
INSERT INTO subreddit (
    title, name, creator_id
) VALUES (
    $1, $2, $3
) RETURNING id, name;

-- name: FindSubredditByName :many
SELECT id FROM subreddit
WHERE name = $1;

-- name: FindSubredditById :many
SELECT id, creator_id FROM subreddit
WHERE id = $1;

-- name: UpdateSubredditTitle :exec
UPDATE subreddit 
SET title = $2
WHERE id = $1;

-- name: UpdateSubredditAbout :exec
UPDATE subreddit 
SET about = $2
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
 
-- name: IsAlreadyJoined :many
SELECT user_id, subreddit_id FROM user_subreddit_join 
WHERE user_id = $1 AND subreddit_id = $2;

-- name: JoinSubreddit :many
INSERT INTO user_subreddit_join (user_id, subreddit_id)
VALUES ($1, $2) 
RETURNING user_id,subreddit_id;

-- name: LeaveSubreddit :exec
DELETE FROM user_subreddit_join 
WHERE user_id = $1 AND subreddit_id = $2;

-- name: FindSubredditByIDPublic :many
SELECT id, name, about, title, avatar, cover, is_verified, created_at, creator_id FROM subreddit 
WHERE id = $1;

-- name: FindSubredditByNamePublic :many
SELECT id, name, about, title, avatar, cover, is_verified, created_at, creator_id FROM subreddit 
WHERE name = $1;