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
SELECT
    subreddit.id,
    subreddit.name,
    subreddit.about,
    subreddit.title,
    subreddit.avatar,
    subreddit.cover,
    subreddit.is_verified,
    subreddit.created_at,
    subreddit.creator_id,
    COALESCE(member.member_count, 0) AS member_count,
    CASE
        WHEN usj.user_id IS NOT NULL THEN true
        ELSE false
    END AS is_joined
FROM
    subreddit
LEFT JOIN (
    SELECT subreddit_id, COUNT(user_id) AS member_count
    FROM user_subreddit_join
    GROUP BY subreddit_id
) AS member ON subreddit.id = member.subreddit_id
LEFT JOIN user_subreddit_join AS usj ON subreddit.id = usj.subreddit_id AND usj.user_id = $2
WHERE
    subreddit.id = $1;

-- name: FindSubredditByNamePublic :many
SELECT
    subreddit.id,
    subreddit.name,
    subreddit.about,
    subreddit.title,
    subreddit.avatar,
    subreddit.cover,
    subreddit.is_verified,
    subreddit.created_at,
    subreddit.creator_id,
    COALESCE(member.member_count, 0) AS member_count,
    CASE
        WHEN usj.user_id IS NOT NULL THEN true
        ELSE false
    END AS is_joined
FROM
    subreddit
LEFT JOIN (
    SELECT subreddit_id, COUNT(user_id) AS member_count
    FROM user_subreddit_join
    GROUP BY subreddit_id
) AS member ON subreddit.id = member.subreddit_id
LEFT JOIN user_subreddit_join AS usj ON subreddit.id = usj.subreddit_id AND usj.user_id = $2
WHERE
    subreddit.name = $1;


-- name: SearchSubredditPublic :many
SELECT
    subreddit.id,
    subreddit.name,
    subreddit.title,
    subreddit.avatar,
    subreddit.is_verified,
    COALESCE(member.member_count, 0) AS member_count,
    CASE
        WHEN usj.user_id IS NOT NULL THEN true
        ELSE false
    END AS is_joined
FROM
    subreddit
LEFT JOIN (
    SELECT subreddit_id, COUNT(user_id) AS member_count
    FROM user_subreddit_join
    GROUP BY subreddit_id
) AS member ON subreddit.id = member.subreddit_id
LEFT JOIN user_subreddit_join AS usj ON subreddit.id = usj.subreddit_id AND usj.user_id = $4
WHERE
    LOWER(subreddit.name) LIKE $1 OR LOWER(subreddit.title) LIKE $1 OR LOWER(subreddit.about) LIKE $1
ORDER BY
    subreddit.created_at DESC
LIMIT
    $2
OFFSET
    $3;

-- name: GetTopSubredditPublic :many
SELECT
    subreddit.id,
    subreddit.name,
    subreddit.title,
    subreddit.avatar,
    subreddit.is_verified,
    COALESCE(member.member_count, 0) AS member_count,
    CASE
        WHEN usj.user_id IS NOT NULL THEN true
        ELSE false
    END AS is_joined
FROM
    subreddit
LEFT JOIN (
    SELECT subreddit_id, COUNT(user_id) AS member_count
    FROM user_subreddit_join
    GROUP BY subreddit_id
) AS member ON subreddit.id = member.subreddit_id
LEFT JOIN user_subreddit_join AS usj ON subreddit.id = usj.subreddit_id AND usj.user_id = $1
ORDER BY
    member_count DESC
LIMIT
    10;

-- name: DeleteSubreddit :exec
DELETE FROM subreddit WHERE id = $1;

-- name: GetJoinedSubreddit :many
SELECT subreddit_id FROM user_subreddit_join 
WHERE user_id = $1;
