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
    COALESCE(replies.replies_count, 0) AS replies_count,
    COALESCE(up_votes.up_vote_count, 0) AS up_vote_count,
    COALESCE(down_votes.down_vote_count, 0) AS down_vote_count,
    COALESCE(user_votes.vote, 0) AS vote
FROM
    posts
JOIN
    users ON posts.creator_id = users.id
JOIN
    subreddit ON posts.subreddit_id = subreddit.id
LEFT JOIN (
    SELECT post_id, COUNT(id) AS replies_count
    FROM replies
    GROUP BY post_id
) AS replies ON posts.id = replies.post_id
LEFT JOIN (
    SELECT post_id, COUNT(user_id) AS up_vote_count
    FROM vote_post
    WHERE down = FALSE
    GROUP BY post_id
) AS up_votes ON posts.id = up_votes.post_id
LEFT JOIN (
    SELECT post_id, COUNT(user_id) AS down_vote_count
    FROM vote_post
    WHERE down = TRUE
    GROUP BY post_id
) AS down_votes ON posts.id = down_votes.post_id
LEFT JOIN (
    SELECT post_id, MAX(CASE WHEN vp.user_id = $2 AND vp.down = FALSE THEN 1 WHEN vp.user_id = $2 AND vp.down = TRUE THEN -1 ELSE 0 END) AS vote
    FROM vote_post AS vp
    WHERE vp.user_id = $2
    GROUP BY vp.post_id
) AS user_votes ON posts.id = user_votes.post_id
WHERE
    posts.id = $1;




-- name: GetPostsOfUser :many
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
    COALESCE(replies.replies_count, 0) AS replies_count,
    COALESCE(up_votes.up_vote_count, 0) AS up_vote_count,
    COALESCE(down_votes.down_vote_count, 0) AS down_vote_count,
    COALESCE(user_votes.vote, 0) AS vote
FROM
    posts
JOIN
    users ON posts.creator_id = users.id
JOIN
    subreddit ON posts.subreddit_id = subreddit.id
LEFT JOIN (
    SELECT post_id, COUNT(id) AS replies_count
    FROM replies
    GROUP BY post_id
) AS replies ON posts.id = replies.post_id
LEFT JOIN (
    SELECT post_id, COUNT(user_id) AS up_vote_count
    FROM vote_post
    WHERE down = FALSE
    GROUP BY post_id
) AS up_votes ON posts.id = up_votes.post_id
LEFT JOIN (
    SELECT post_id, COUNT(user_id) AS down_vote_count
    FROM vote_post
    WHERE down = TRUE
    GROUP BY post_id
) AS down_votes ON posts.id = down_votes.post_id
LEFT JOIN (
    SELECT post_id, MAX(CASE WHEN vp.user_id = $4 AND vp.down = FALSE THEN 1 WHEN vp.user_id = $4 AND vp.down = TRUE THEN -1 ELSE 0 END) AS vote
    FROM vote_post AS vp
    WHERE vp.user_id = $4
    GROUP BY vp.post_id
) AS user_votes ON posts.id = user_votes.post_id
WHERE
    posts.creator_id = $1
ORDER BY
    posts.created_at DESC
LIMIT
    $2
OFFSET
    $3;

-- name: GetSubredditPosts :many
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
    COALESCE(replies.replies_count, 0) AS replies_count,
    COALESCE(up_votes.up_vote_count, 0) AS up_vote_count,
    COALESCE(down_votes.down_vote_count, 0) AS down_vote_count,
    COALESCE(user_votes.vote, 0) AS vote
FROM
    posts
JOIN
    users ON posts.creator_id = users.id
JOIN
    subreddit ON posts.subreddit_id = subreddit.id
LEFT JOIN (
    SELECT post_id, COUNT(id) AS replies_count
    FROM replies
    GROUP BY post_id
) AS replies ON posts.id = replies.post_id
LEFT JOIN (
    SELECT post_id, COUNT(user_id) AS up_vote_count
    FROM vote_post
    WHERE down = FALSE
    GROUP BY post_id
) AS up_votes ON posts.id = up_votes.post_id
LEFT JOIN (
    SELECT post_id, COUNT(user_id) AS down_vote_count
    FROM vote_post
    WHERE down = TRUE
    GROUP BY post_id
) AS down_votes ON posts.id = down_votes.post_id
LEFT JOIN (
    SELECT post_id, MAX(CASE WHEN vp.user_id = $4 AND vp.down = FALSE THEN 1 WHEN vp.user_id = $4 AND vp.down = TRUE THEN -1 ELSE 0 END) AS vote
    FROM vote_post AS vp
    WHERE vp.user_id = $4
    GROUP BY vp.post_id
) AS user_votes ON posts.id = user_votes.post_id
WHERE
    posts.subreddit_id = $1
ORDER BY
    posts.created_at DESC
LIMIT
    $2
OFFSET
    $3;

-- name: GetFeedPosts :many
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
    COALESCE(replies.replies_count, 0) AS replies_count,
    COALESCE(up_votes.up_vote_count, 0) AS up_vote_count,
    COALESCE(down_votes.down_vote_count, 0) AS down_vote_count,
    COALESCE(user_votes.vote, 0) AS vote
FROM
    posts
JOIN
    users ON posts.creator_id = users.id
JOIN
    subreddit ON posts.subreddit_id = subreddit.id
LEFT JOIN (
    SELECT post_id, COUNT(id) AS replies_count
    FROM replies
    GROUP BY post_id
) AS replies ON posts.id = replies.post_id
LEFT JOIN (
    SELECT post_id, COUNT(user_id) AS up_vote_count
    FROM vote_post
    WHERE down = FALSE
    GROUP BY post_id
) AS up_votes ON posts.id = up_votes.post_id
LEFT JOIN (
    SELECT post_id, COUNT(user_id) AS down_vote_count
    FROM vote_post
    WHERE down = TRUE
    GROUP BY post_id
) AS down_votes ON posts.id = down_votes.post_id
LEFT JOIN (
    SELECT post_id, MAX(CASE WHEN vp.user_id = $4 AND vp.down = FALSE THEN 1 WHEN vp.user_id = $4 AND vp.down = TRUE THEN -1 ELSE 0 END) AS vote
    FROM vote_post AS vp
    WHERE vp.user_id = $4
    GROUP BY vp.post_id
) AS user_votes ON posts.id = user_votes.post_id
WHERE
    posts.subreddit_id = ANY($1::int[])
ORDER BY
    posts.created_at DESC
LIMIT
    $2
OFFSET
    $3;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;