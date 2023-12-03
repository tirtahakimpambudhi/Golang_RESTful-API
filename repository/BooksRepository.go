package repository

import (
	"context"
	"database/sql"
	"restful_api/model/domain"
)

type BooksRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, limit, offset int) ([]domain.Books, int)
	FindByISBN(ctx context.Context, tx *sql.Tx, ISBN string) (domain.Books, error)
	Create(ctx context.Context, tx *sql.Tx, books domain.Books) error
	CreateMany(ctx context.Context, tx *sql.Tx, books []domain.Books) error
	Update(ctx context.Context, tx *sql.Tx, ISBN string, books domain.Books)
	Delete(ctx context.Context, tx *sql.Tx, ISBN string)
	DeleteMany(ctx context.Context, tx *sql.Tx, ISBNs []string)
}
