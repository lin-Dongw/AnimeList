// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: episode.sql

package db

import (
	"context"
)

const createEpisode = `-- name: CreateEpisode :one
INSERT INTO episodes (
    season_id,
    name,
    value,
    desc,
    status
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING id, season_id, name, value, "desc", status, created_at, updated_at
`

type CreateEpisodeParams struct {
	SeasonID int64  `json:"season_id"`
	Name     string `json:"name"`
	Value    int64  `json:"value"`
	Desc     string `json:"desc"`
	Status   int64  `json:"status"`
}

func (q *Queries) CreateEpisode(ctx context.Context, arg CreateEpisodeParams) (Episode, error) {
	row := q.db.QueryRowContext(ctx, createEpisode,
		arg.SeasonID,
		arg.Name,
		arg.Value,
		arg.Desc,
		arg.Status,
	)
	var i Episode
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.Name,
		&i.Value,
		&i.Desc,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteEpisode = `-- name: DeleteEpisode :exec
DELETE FROM episodes
WHERE id = ?
`

func (q *Queries) DeleteEpisode(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEpisode, id)
	return err
}

const getEpisode = `-- name: GetEpisode :one
SELECT id, season_id, name, value, "desc", status, created_at, updated_at FROM episodes
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetEpisode(ctx context.Context, id int64) (Episode, error) {
	row := q.db.QueryRowContext(ctx, getEpisode, id)
	var i Episode
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.Name,
		&i.Value,
		&i.Desc,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listEpisode = `-- name: ListEpisode :many
SELECT id, season_id, name, value, "desc", status, created_at, updated_at FROM episodes
ORDER BY id
LIMIT ?
OFFSET ?
`

type ListEpisodeParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) ListEpisode(ctx context.Context, arg ListEpisodeParams) ([]Episode, error) {
	rows, err := q.db.QueryContext(ctx, listEpisode, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Episode{}
	for rows.Next() {
		var i Episode
		if err := rows.Scan(
			&i.ID,
			&i.SeasonID,
			&i.Name,
			&i.Value,
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

const updateEpisode = `-- name: UpdateEpisode :one
UPDATE episodes
SET name = ?, value = ?, desc = ?, status = ?, updated_at = strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime')
WHERE id = ?
RETURNING id, season_id, name, value, "desc", status, created_at, updated_at
`

type UpdateEpisodeParams struct {
	Name   string `json:"name"`
	Value  int64  `json:"value"`
	Desc   string `json:"desc"`
	Status int64  `json:"status"`
	ID     int64  `json:"id"`
}

func (q *Queries) UpdateEpisode(ctx context.Context, arg UpdateEpisodeParams) (Episode, error) {
	row := q.db.QueryRowContext(ctx, updateEpisode,
		arg.Name,
		arg.Value,
		arg.Desc,
		arg.Status,
		arg.ID,
	)
	var i Episode
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.Name,
		&i.Value,
		&i.Desc,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
