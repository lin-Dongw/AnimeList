// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: anime.sql

package db

import (
	"context"
)

const createAnime = `-- name: CreateAnime :one
INSERT INTO animes (
    user_id,
    name,
    desc,
    status
) VALUES (
    ?, ?, ?, ?
)
RETURNING id, user_id, name, "desc", status, created_at, updated_at
`

type CreateAnimeParams struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Status int64  `json:"status"`
}

func (q *Queries) CreateAnime(ctx context.Context, arg CreateAnimeParams) (Anime, error) {
	row := q.db.QueryRowContext(ctx, createAnime,
		arg.UserID,
		arg.Name,
		arg.Desc,
		arg.Status,
	)
	var i Anime
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Desc,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAnime = `-- name: DeleteAnime :exec
DELETE FROM animes
WHERE id = ?
`

func (q *Queries) DeleteAnime(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAnime, id)
	return err
}

const getAnime = `-- name: GetAnime :one
SELECT id, user_id, name, "desc", status, created_at, updated_at FROM animes
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetAnime(ctx context.Context, id int64) (Anime, error) {
	row := q.db.QueryRowContext(ctx, getAnime, id)
	var i Anime
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Desc,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAnime = `-- name: ListAnime :many
SELECT id, user_id, name, "desc", status, created_at, updated_at FROM animes
ORDER BY id
LIMIT ?
OFFSET ?
`

type ListAnimeParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) ListAnime(ctx context.Context, arg ListAnimeParams) ([]Anime, error) {
	rows, err := q.db.QueryContext(ctx, listAnime, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Anime{}
	for rows.Next() {
		var i Anime
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Desc,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateAnime = `-- name: UpdateAnime :one
UPDATE animes
SET name = ?, desc = ?, status = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
WHERE id = ?
RETURNING id, user_id, name, "desc", status, created_at, updated_at
`

type UpdateAnimeParams struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Status int64  `json:"status"`
	ID     int64  `json:"id"`
}

func (q *Queries) UpdateAnime(ctx context.Context, arg UpdateAnimeParams) (Anime, error) {
	row := q.db.QueryRowContext(ctx, updateAnime,
		arg.Name,
		arg.Desc,
		arg.Status,
		arg.ID,
	)
	var i Anime
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Desc,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
