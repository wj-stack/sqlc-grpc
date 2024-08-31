// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package googleuuid

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	ID int32 `json:"id"`
}

type Location struct {
	ID uuid.UUID `json:"id"`
}

type Product struct {
	ID       int32       `json:"id"`
	Category pgtype.Int4 `json:"category"`
	Name     pgtype.Text `json:"name"`
}

type User struct {
	ID       uuid.UUID   `json:"id"`
	Location pgtype.UUID `json:"location"`
	Name     pgtype.Text `json:"name"`
}
