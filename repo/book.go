package repo

import (
	"go-template/logger"
	"go-template/model"

	"gorm.io/gorm"
)

type BookRepo struct {
	db  *gorm.DB
	log logger.StructLogger
}

func NewBookRepo(db *gorm.DB, log logger.StructLogger) *BookRepo {
	return &BookRepo{
		db:  db,
		log: log,
	}
}

func (b *BookRepo) Create(book *model.Book) error {
	resp := b.db.Create(book)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (b *BookRepo) List(offset, limit int) ([]model.Book, error) {
	books := make([]model.Book, limit)
	resp := b.db.Offset(offset).Limit(limit).Find(&books)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return books, nil
}
