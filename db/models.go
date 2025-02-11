// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Restriction string

const (
	RestrictionReservation Restriction = "reservation"
	RestrictionOwnerBlock  Restriction = "owner_block"
)

func (e *Restriction) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Restriction(s)
	case string:
		*e = Restriction(s)
	default:
		return fmt.Errorf("unsupported scan type for Restriction: %T", src)
	}
	return nil
}

type NullRestriction struct {
	Restriction Restriction `json:"restriction"`
	Valid       bool        `json:"valid"` // Valid is true if Restriction is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRestriction) Scan(value interface{}) error {
	if value == nil {
		ns.Restriction, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Restriction.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRestriction) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Restriction), nil
}

type Reservation struct {
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
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type Room struct {
	ID            int64              `json:"id"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	ImageFilename string             `json:"image_filename"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

type RoomRestriction struct {
	ID            int64              `json:"id"`
	StartDate     pgtype.Date        `json:"start_date"`
	EndDate       pgtype.Date        `json:"end_date"`
	RoomID        int64              `json:"room_id"`
	ReservationID pgtype.Int8        `json:"reservation_id"`
	Restriction   Restriction        `json:"restriction"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

type User struct {
	ID          int64              `json:"id"`
	FirstName   string             `json:"first_name"`
	LastName    string             `json:"last_name"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	AccessLevel int64              `json:"access_level"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}
