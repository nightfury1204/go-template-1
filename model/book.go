package model

import (
	"go-template/api/request"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

func FromBookRequestToBook(in *request.BookReq) *Book {
	return &Book{
		Name:   in.Name,
		Author: in.Author,
	}
}
