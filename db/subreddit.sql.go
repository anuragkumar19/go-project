// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: subreddit.sql

package database

import (
	"context"
	"time"
)

const createSubreddit = `-- name: CreateSubreddit :many
INSERT INTO subreddit (
    title, name, creator_id
) VALUES (
    $1, $2, $3
) RETURNING id, name
`

type CreateSubredditParams struct {
	Title     string `json:"title"`
	Name      string `json:"name"`
	CreatorID int32  `json:"creator_id"`
}

type CreateSubredditRow struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) CreateSubreddit(ctx context.Context, arg CreateSubredditParams) ([]CreateSubredditRow, error) {
	rows, err := q.db.QueryContext(ctx, createSubreddit, arg.Title, arg.Name, arg.CreatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreateSubredditRow
	for rows.Next() {
		var i CreateSubredditRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findSubredditByIDPublic = `-- name: FindSubredditByIDPublic :many
SELECT id, name, about, title, avatar, cover, is_verified, created_at, creator_id FROM subreddit 
WHERE id = $1
`

type FindSubredditByIDPublicRow struct {
	ID         int32     `json:"id"`
	Name       string    `json:"name"`
	About      string    `json:"about"`
	Title      string    `json:"title"`
	Avatar     string    `json:"avatar"`
	Cover      string    `json:"cover"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	CreatorID  int32     `json:"creator_id"`
}

func (q *Queries) FindSubredditByIDPublic(ctx context.Context, id int32) ([]FindSubredditByIDPublicRow, error) {
	rows, err := q.db.QueryContext(ctx, findSubredditByIDPublic, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindSubredditByIDPublicRow
	for rows.Next() {
		var i FindSubredditByIDPublicRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.About,
			&i.Title,
			&i.Avatar,
			&i.Cover,
			&i.IsVerified,
			&i.CreatedAt,
			&i.CreatorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findSubredditById = `-- name: FindSubredditById :many
SELECT id, creator_id FROM subreddit
WHERE id = $1
`

type FindSubredditByIdRow struct {
	ID        int32 `json:"id"`
	CreatorID int32 `json:"creator_id"`
}

func (q *Queries) FindSubredditById(ctx context.Context, id int32) ([]FindSubredditByIdRow, error) {
	rows, err := q.db.QueryContext(ctx, findSubredditById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindSubredditByIdRow
	for rows.Next() {
		var i FindSubredditByIdRow
		if err := rows.Scan(&i.ID, &i.CreatorID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findSubredditByName = `-- name: FindSubredditByName :many
SELECT id FROM subreddit
WHERE name = $1
`

func (q *Queries) FindSubredditByName(ctx context.Context, name string) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, findSubredditByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findSubredditByNamePublic = `-- name: FindSubredditByNamePublic :many
SELECT id, name, about, title, avatar, cover, is_verified, created_at, creator_id FROM subreddit 
WHERE name = $1
`

type FindSubredditByNamePublicRow struct {
	ID         int32     `json:"id"`
	Name       string    `json:"name"`
	About      string    `json:"about"`
	Title      string    `json:"title"`
	Avatar     string    `json:"avatar"`
	Cover      string    `json:"cover"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	CreatorID  int32     `json:"creator_id"`
}

func (q *Queries) FindSubredditByNamePublic(ctx context.Context, name string) ([]FindSubredditByNamePublicRow, error) {
	rows, err := q.db.QueryContext(ctx, findSubredditByNamePublic, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindSubredditByNamePublicRow
	for rows.Next() {
		var i FindSubredditByNamePublicRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.About,
			&i.Title,
			&i.Avatar,
			&i.Cover,
			&i.IsVerified,
			&i.CreatedAt,
			&i.CreatorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isAlreadyJoined = `-- name: IsAlreadyJoined :many
SELECT user_id, subreddit_id FROM user_subreddit_join 
WHERE user_id = $1 AND subreddit_id = $2
`

type IsAlreadyJoinedParams struct {
	UserID      int32 `json:"user_id"`
	SubredditID int32 `json:"subreddit_id"`
}

func (q *Queries) IsAlreadyJoined(ctx context.Context, arg IsAlreadyJoinedParams) ([]UserSubredditJoin, error) {
	rows, err := q.db.QueryContext(ctx, isAlreadyJoined, arg.UserID, arg.SubredditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserSubredditJoin
	for rows.Next() {
		var i UserSubredditJoin
		if err := rows.Scan(&i.UserID, &i.SubredditID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const joinSubreddit = `-- name: JoinSubreddit :many
INSERT INTO user_subreddit_join (user_id, subreddit_id)
VALUES ($1, $2) 
RETURNING user_id,subreddit_id
`

type JoinSubredditParams struct {
	UserID      int32 `json:"user_id"`
	SubredditID int32 `json:"subreddit_id"`
}

func (q *Queries) JoinSubreddit(ctx context.Context, arg JoinSubredditParams) ([]UserSubredditJoin, error) {
	rows, err := q.db.QueryContext(ctx, joinSubreddit, arg.UserID, arg.SubredditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserSubredditJoin
	for rows.Next() {
		var i UserSubredditJoin
		if err := rows.Scan(&i.UserID, &i.SubredditID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const leaveSubreddit = `-- name: LeaveSubreddit :exec
DELETE FROM user_subreddit_join 
WHERE user_id = $1 AND subreddit_id = $2
`

type LeaveSubredditParams struct {
	UserID      int32 `json:"user_id"`
	SubredditID int32 `json:"subreddit_id"`
}

func (q *Queries) LeaveSubreddit(ctx context.Context, arg LeaveSubredditParams) error {
	_, err := q.db.ExecContext(ctx, leaveSubreddit, arg.UserID, arg.SubredditID)
	return err
}

const updateSubredditAbout = `-- name: UpdateSubredditAbout :exec
UPDATE subreddit 
SET about = $2
WHERE id = $1
`

type UpdateSubredditAboutParams struct {
	ID    int32  `json:"id"`
	About string `json:"about"`
}

func (q *Queries) UpdateSubredditAbout(ctx context.Context, arg UpdateSubredditAboutParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditAbout, arg.ID, arg.About)
	return err
}

const updateSubredditAvatar = `-- name: UpdateSubredditAvatar :exec
UPDATE subreddit 
SET avatar = $2
WHERE id = $1
`

type UpdateSubredditAvatarParams struct {
	ID     int32  `json:"id"`
	Avatar string `json:"avatar"`
}

func (q *Queries) UpdateSubredditAvatar(ctx context.Context, arg UpdateSubredditAvatarParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditAvatar, arg.ID, arg.Avatar)
	return err
}

const updateSubredditCover = `-- name: UpdateSubredditCover :exec
UPDATE subreddit 
SET cover = $2
WHERE id = $1
`

type UpdateSubredditCoverParams struct {
	ID    int32  `json:"id"`
	Cover string `json:"cover"`
}

func (q *Queries) UpdateSubredditCover(ctx context.Context, arg UpdateSubredditCoverParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditCover, arg.ID, arg.Cover)
	return err
}

const updateSubredditName = `-- name: UpdateSubredditName :exec
UPDATE subreddit 
SET name = $2
WHERE id = $1
`

type UpdateSubredditNameParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateSubredditName(ctx context.Context, arg UpdateSubredditNameParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditName, arg.ID, arg.Name)
	return err
}

const updateSubredditTitle = `-- name: UpdateSubredditTitle :exec
UPDATE subreddit 
SET title = $2
WHERE id = $1
`

type UpdateSubredditTitleParams struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

func (q *Queries) UpdateSubredditTitle(ctx context.Context, arg UpdateSubredditTitleParams) error {
	_, err := q.db.ExecContext(ctx, updateSubredditTitle, arg.ID, arg.Title)
	return err
}
