package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"math"
	"restful_api/config"
	"restful_api/exception"
	"restful_api/helper"
	"restful_api/model/web"
	"restful_api/repository"
)

type BooksServiceImpl struct {
	BooksRepository repository.BooksRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func (service *BooksServiceImpl) FindAllBooks(ctx context.Context, currentPage string) ([]web.ResponseBooks, int) {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)
	if currentPage == "" || currentPage == "0" {
		currentPage = "1"
	}
	currentPageParse := helper.ParamsParseInt(currentPage)
	limit := config.DataPerPage
	offset := (currentPageParse - 1) * limit

	books, totalBooks := service.BooksRepository.FindAll(ctx, tx, limit, offset)

	// Log: Received request to find all books
	fmt.Println("Received request to find all books")

	return helper.ToResponsesBooks(books), totalBooks
}

func (service *BooksServiceImpl) Pagination(ctx context.Context, pageParams string, totalRecords int) web.Pagination {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	if pageParams == "" || pageParams == "0" {
		pageParams = "1"
	}

	totalPage := math.Ceil(float64(totalRecords) / float64(config.DataPerPage))
	currentPage := helper.ParamsParseInt(pageParams)

	// Log: Received request for pagination
	fmt.Println("Received request for pagination")

	return web.Pagination{
		CurrentPage: currentPage,
		TotalPage:   int(totalPage),
	}
}

func (service *BooksServiceImpl) FindBookByISBN(ctx context.Context, ISBN string) web.ResponseBooks {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BooksRepository.FindByISBN(ctx, tx, ISBN)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	// Log: Received request to find a book by ISBN
	fmt.Printf("Received request to find a book by ISBN (%v)\n", ISBN)

	return helper.ToResponseBooks(book)
}

func (service *BooksServiceImpl) CreateBook(ctx context.Context, request web.RequestBody) web.ResponseBooks {
	validationError := service.Validate.Struct(request)
	helper.PanicIFError(validationError)

	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	body := helper.ToDomainBooks(request)

	// Log: Received request to create a book
	fmt.Println("Received request to create a book")

	err = service.BooksRepository.Create(ctx, tx, body)
	if err != nil {
		exception.NewErrorDuplcated(err.Error())
	}

	return helper.ToResponseBooks(body)
}

func (service *BooksServiceImpl) UpdateBook(ctx context.Context, ISBN string, request web.RequestBody) web.ResponseBooks {
	validationError := service.Validate.Struct(request)
	helper.PanicIFError(validationError)

	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.BooksRepository.FindByISBN(ctx, tx, ISBN)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	// Log: Received request to update a book
	fmt.Printf("Received request to update a book with ISBN (%v)\n", ISBN)

	body := helper.ToDomainBooks(request)
	service.BooksRepository.Update(ctx, tx, ISBN, body)

	return helper.ToResponseBooks(body)
}

func (service *BooksServiceImpl) DeleteBook(ctx context.Context, ISBN string) {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.BooksRepository.FindByISBN(ctx, tx, ISBN)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	// Log: Received request to delete a book
	fmt.Printf("Received request to delete a book with ISBN (%v)\n", ISBN)

	service.BooksRepository.Delete(ctx, tx, ISBN)
}

func (service *BooksServiceImpl) CreateBooks(ctx context.Context, requests []web.RequestBody) []web.ResponseBooks {
	for _, request := range requests {
		validationError := service.Validate.Struct(request)
		if validationError != nil {
			helper.PanicIFError(validationError)
			break
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	// Log: Received request to create a book
	fmt.Println("Received request to create many books")

	body := helper.ToDomainsBooks(requests)
	err = service.BooksRepository.CreateMany(ctx, tx, body)
	if err != nil {
		panic(exception.NewErrorDuplcated(err.Error()))
	}

	return helper.ToResponsesBooks(body)
}

func (service *BooksServiceImpl) DeleteBooks(ctx context.Context, ISBNs []string) {
	tx, err := service.DB.Begin()
	helper.PanicIFError(err)
	defer helper.CommitOrRollback(tx)

	var found bool = true // Inisialisasi sebagai true

	for _, ISBN := range ISBNs {
		_, err := service.BooksRepository.FindByISBN(ctx, tx, ISBN)
		if err != nil {
			found = false
			break
		}
	}

	if !found {
		panic(exception.NewErrorNotFound("NOT FOUND"))
	}

	// Log: Received request to delete a book
	fmt.Printf("Received request to delete a book with ISBN (%v)\n", helper.SliceToInterface(ISBNs)...)

	service.BooksRepository.DeleteMany(ctx, tx, ISBNs)
}

func NewBooksService(BooksRepo repository.BooksRepository, Database *sql.DB, Validation *validator.Validate) BooksService {
	return &BooksServiceImpl{
		BooksRepository: BooksRepo,
		DB:              Database,
		Validate:        Validation,
	}
}
