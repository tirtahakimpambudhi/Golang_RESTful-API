package helper

import (
	"restful_api/model/domain"
	"restful_api/model/web"
)

func ToResponseBooks(book domain.Books) web.ResponseBooks {
	return web.ResponseBooks{
		ISBN: book.ISBN,
		Title: book.Title,
		Author: book.Author,
		Status_Borrow: book.Status_Borrow,
		Publisher: book.Publisher,
		Publication_Years: book.Publication_Years,
		Description: book.Description,
		CreatedAt: book.CreatedAt,
		UpdateAt:  book.UpdateAt,
	}
}

func ToResponsesBooks( books []domain.Books ) []web.ResponseBooks {
	var responsesBooks []web.ResponseBooks
	for _, book := range books {
		responsesBooks = append(responsesBooks, ToResponseBooks(book))
	}
	return responsesBooks
}

func ToDomainBooks(book web.RequestBody) domain.Books {
	return domain.Books{
		ISBN: book.ISBN,
		Title: book.Title,
		Author: book.Author,
		Status_Borrow: book.Status_Borrow,
		Publisher: book.Publisher,
		Publication_Years: book.Publication_Years,
		Description: book.Description,
	}
}

func ToDomainsBooks(books []web.RequestBody) []domain.Books {
	var result []domain.Books
	for _, book := range books {
		result = append(result, ToDomainBooks(book))
	}
	return result
}