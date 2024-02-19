package model

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

// Stock db model
type Stock struct {
	ID          uuid.UUID   `db:"id"`
	Name        null.String `db:"name"`
	IsAvailable null.Bool   `db:"is_available"`
	CreatedAt   null.Time   `db:"created_at"`
	UpdatedAt   null.Time   `db:"updated_at"`
	DeletedAt   null.Time   `db:"deleted_at"`
}
