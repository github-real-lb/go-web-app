// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: reservation.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createReservation = `-- name: CreateReservation :one
INSERT INTO reservations (
  code, first_name, last_name, email, phone, start_date, end_date, room_id, notes
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, code, first_name, last_name, email, phone, start_date, end_date, room_id, notes, created_at, updated_at
`

type CreateReservationParams struct {
	Code      string      `json:"code"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     pgtype.Text `json:"phone"`
	StartDate pgtype.Date `json:"start_date"`
	EndDate   pgtype.Date `json:"end_date"`
	RoomID    int64       `json:"room_id"`
	Notes     pgtype.Text `json:"notes"`
}

func (q *Queries) CreateReservation(ctx context.Context, arg CreateReservationParams) (Reservation, error) {
	row := q.db.QueryRow(ctx, createReservation,
		arg.Code,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.StartDate,
		arg.EndDate,
		arg.RoomID,
		arg.Notes,
	)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteReservation = `-- name: DeleteReservation :exec
DELETE FROM reservations
WHERE id = $1
`

func (q *Queries) DeleteReservation(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteReservation, id)
	return err
}

const getReservation = `-- name: GetReservation :one
SELECT id, code, first_name, last_name, email, phone, start_date, end_date, room_id, notes, created_at, updated_at FROM reservations
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetReservation(ctx context.Context, id int64) (Reservation, error) {
	row := q.db.QueryRow(ctx, getReservation, id)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReservationByCode = `-- name: GetReservationByCode :one
SELECT id, code, first_name, last_name, email, phone, start_date, end_date, room_id, notes, created_at, updated_at FROM reservations
WHERE code = $1 LIMIT 1
`

func (q *Queries) GetReservationByCode(ctx context.Context, code string) (Reservation, error) {
	row := q.db.QueryRow(ctx, getReservationByCode, code)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReservationByLastName = `-- name: GetReservationByLastName :one
SELECT id, code, first_name, last_name, email, phone, start_date, end_date, room_id, notes, created_at, updated_at FROM reservations
WHERE code = $1 AND last_name = $2 LIMIT 1
`

type GetReservationByLastNameParams struct {
	Code     string `json:"code"`
	LastName string `json:"last_name"`
}

func (q *Queries) GetReservationByLastName(ctx context.Context, arg GetReservationByLastNameParams) (Reservation, error) {
	row := q.db.QueryRow(ctx, getReservationByLastName, arg.Code, arg.LastName)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listReservations = `-- name: ListReservations :many
SELECT id, code, first_name, last_name, email, phone, start_date, end_date, room_id, notes, created_at, updated_at FROM reservations
ORDER BY created_at DESC
LIMIT $1
OFFSET $2
`

type ListReservationsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListReservations(ctx context.Context, arg ListReservationsParams) ([]Reservation, error) {
	rows, err := q.db.Query(ctx, listReservations, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Reservation{}
	for rows.Next() {
		var i Reservation
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.Notes,
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

const listReservationsByRoom = `-- name: ListReservationsByRoom :many
SELECT id, code, first_name, last_name, email, phone, start_date, end_date, room_id, notes, created_at, updated_at FROM reservations
WHERE room_id = $1
ORDER BY start_date, end_date DESC
LIMIT $2
OFFSET $3
`

type ListReservationsByRoomParams struct {
	RoomID int64 `json:"room_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListReservationsByRoom(ctx context.Context, arg ListReservationsByRoomParams) ([]Reservation, error) {
	rows, err := q.db.Query(ctx, listReservationsByRoom, arg.RoomID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Reservation{}
	for rows.Next() {
		var i Reservation
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.Notes,
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

const updateReservation = `-- name: UpdateReservation :exec
UPDATE reservations
  set   code = $2,
        first_name = $3,
        last_name = $4, 
        email = $5,
        phone = $6, 
        start_date =  $7,
        end_date = $8,
        room_id = $9,
        notes = $10,
        updated_at = $11
WHERE id = $1
`

type UpdateReservationParams struct {
	ID        int64              `json:"id"`
	Code      string             `json:"code"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	StartDate pgtype.Date        `json:"start_date"`
	EndDate   pgtype.Date        `json:"end_date"`
	RoomID    int64              `json:"room_id"`
	Notes     pgtype.Text        `json:"notes"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) UpdateReservation(ctx context.Context, arg UpdateReservationParams) error {
	_, err := q.db.Exec(ctx, updateReservation,
		arg.ID,
		arg.Code,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.StartDate,
		arg.EndDate,
		arg.RoomID,
		arg.Notes,
		arg.UpdatedAt,
	)
	return err
}
