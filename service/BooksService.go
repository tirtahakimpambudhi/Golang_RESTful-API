package service

import (
	"context"
	"restful_api/model/web"
)
type BooksService interface {
	FindAllBooks(ctx context.Context , page web.Pagination) []web.ResponseBooks
	FindBookByISBN(ctx context.Context , ISBN string) web.ResponseBooks
	CreateBook(ctx context.Context , request web.RequestBody) web.ResponseBooks
	CreateBooks(ctx context.Context , requests []web.RequestBody) []web.ResponseBooks
	UpdateBook(ctx context.Context ,  ISBN string , request web.RequestBody) web.ResponseBooks
	DeleteBook(ctx context.Context , ISBN string) 
	DeleteBooks(ctx context.Context , ISBNs []string) 
	Pagination(ctx context.Context , pageParams string) web.Pagination
}