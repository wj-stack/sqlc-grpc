// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package authors

import (
	"context"
)

type Querier interface {
	// http: POST /authors
	CreateAuthor(ctx context.Context, db DBTX, arg *CreateAuthorParams) (*Authors, error)
	// http: DELETE /authors/{id}
	DeleteAuthor(ctx context.Context, db DBTX, id int64) error
	// http: GET /authors/{id}
	GetAuthor(ctx context.Context, db DBTX, id int64) (*Authors, error)
	// http: GET /authors
	ListAuthors(ctx context.Context, db DBTX) ([]*Authors, error)
	// http: PATCH /authors/{id}/bio
	UpdateAuthorBio(ctx context.Context, db DBTX, arg *UpdateAuthorBioParams) error
}

var _ Querier = (*Queries)(nil)
