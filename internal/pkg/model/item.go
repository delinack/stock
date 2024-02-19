package model

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

// Item db model
type Item struct {
	ID         uuid.UUID   `db:"id"`
	Name       null.String `db:"name"`
	Size       null.String `db:"size"`
	UniqueCode null.String `db:"unique_code"`
	Quantity   null.Int    `db:"quantity"`
	CreatedAt  null.Time   `db:"created_at"`
	UpdatedAt  null.Time   `db:"updated_at"`
	DeletedAt  null.Time   `db:"deleted_at"`
}
