// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: room_restriction.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRoomRestriction = `-- name: CreateRoomRestriction :one
INSERT INTO room_restrictions (
  start_date, end_date, room_id, reservation_id, restriction
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, start_date, end_date, room_id, reservation_id, restriction, created_at, updated_at
`

type CreateRoomRestrictionParams struct {
	StartDate     pgtype.Date `json:"start_date"`
	EndDate       pgtype.Date `json:"end_date"`
	RoomID        int64       `json:"room_id"`
	ReservationID pgtype.Int8 `json:"reservation_id"`
	Restriction   Restriction `json:"restriction"`
}

func (q *Queries) CreateRoomRestriction(ctx context.Context, arg CreateRoomRestrictionParams) (RoomRestriction, error) {
	row := q.db.QueryRow(ctx, createRoomRestriction,
		arg.StartDate,
		arg.EndDate,
		arg.RoomID,
		arg.ReservationID,
		arg.Restriction,
	)
	var i RoomRestriction
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.ReservationID,
		&i.Restriction,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAllRoomRestrictions = `-- name: DeleteAllRoomRestrictions :exec
DELETE FROM room_restrictions
`

func (q *Queries) DeleteAllRoomRestrictions(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteAllRoomRestrictions)
	return err
}

const deleteRoomRestriction = `-- name: DeleteRoomRestriction :exec
DELETE FROM room_restrictions
WHERE id = $1
`

func (q *Queries) DeleteRoomRestriction(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteRoomRestriction, id)
	return err
}

const getLastRoomRestriction = `-- name: GetLastRoomRestriction :one
SELECT id, start_date, end_date, room_id, reservation_id, restriction, created_at, updated_at FROM room_restrictions
WHERE room_id = $1 
ORDER BY created_at DESC
LIMIT 1
`

func (q *Queries) GetLastRoomRestriction(ctx context.Context, roomID int64) (RoomRestriction, error) {
	row := q.db.QueryRow(ctx, getLastRoomRestriction, roomID)
	var i RoomRestriction
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.ReservationID,
		&i.Restriction,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoomRestriction = `-- name: GetRoomRestriction :one
SELECT id, start_date, end_date, room_id, reservation_id, restriction, created_at, updated_at FROM room_restrictions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRoomRestriction(ctx context.Context, id int64) (RoomRestriction, error) {
	row := q.db.QueryRow(ctx, getRoomRestriction, id)
	var i RoomRestriction
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.ReservationID,
		&i.Restriction,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listRoomRestrictions = `-- name: ListRoomRestrictions :many
SELECT id, start_date, end_date, room_id, reservation_id, restriction, created_at, updated_at FROM room_restrictions
ORDER BY room_id, start_date
LIMIT $1
OFFSET $2
`

type ListRoomRestrictionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRoomRestrictions(ctx context.Context, arg ListRoomRestrictionsParams) ([]RoomRestriction, error) {
	rows, err := q.db.Query(ctx, listRoomRestrictions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RoomRestriction{}
	for rows.Next() {
		var i RoomRestriction
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.ReservationID,
			&i.Restriction,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRoomRestriction = `-- name: UpdateRoomRestriction :exec
UPDATE room_restrictions
  set   start_date = $2,
        end_date = $3, 
        room_id = $4,
        reservation_id = $5, 
        restriction =  $6,
        updated_at = $7
WHERE id = $1
`

type UpdateRoomRestrictionParams struct {
	ID            int64              `json:"id"`
	StartDate     pgtype.Date        `json:"start_date"`
	EndDate       pgtype.Date        `json:"end_date"`
	RoomID        int64              `json:"room_id"`
	ReservationID pgtype.Int8        `json:"reservation_id"`
	Restriction   Restriction        `json:"restriction"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) UpdateRoomRestriction(ctx context.Context, arg UpdateRoomRestrictionParams) error {
	_, err := q.db.Exec(ctx, updateRoomRestriction,
		arg.ID,
		arg.StartDate,
		arg.EndDate,
		arg.RoomID,
		arg.ReservationID,
		arg.Restriction,
		arg.UpdatedAt,
	)
	return err
}
