package service

import (
	"context"
	"go-template/api/request"
	"go-template/api/response"
	"go-template/logger"
	"go-template/model"
	"go-template/repo"
	"go-template/utils"
)

type BookService struct {
	bookRepo *repo.BookRepo
	log      logger.StructLogger
}

func NewBookService(bookRepo *repo.BookRepo, log logger.StructLogger) *BookService {
	return &BookService{
		bookRepo: bookRepo,
		log:      log,
	}
}

func (b *BookService) CreateBook(ctx context.Context, payload *request.BookReq) error {
	tid := utils.GetTracingID(ctx)
	b.log.Println("BookService_CreateBook", tid, "creating book")

	book := model.FromBookRequestToBook(payload)
	return b.bookRepo.Create(book)
}

func (b *BookService) ListBooks(ctx context.Context, offset, limit int) ([]response.BookResp, error) {
	tid := utils.GetTracingID(ctx)
	b.log.Printf("BookService_ListBooks", tid, "listing books, offset=%d, limit=%d", offset, limit)

	books, err := b.bookRepo.List(offset, limit)
	if err != nil {
		return nil, err
	}

	return response.ToBookResp(books), nil
}
