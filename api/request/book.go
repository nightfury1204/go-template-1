package request

import "fmt"

type BookReq struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func ValidateBookCreateReq(book *BookReq) error {
	if book == nil {
		return fmt.Errorf("book is nil")
	}
	if book.Name == "" {
		return fmt.Errorf("name field should be non-empty")
	}
	if book.Author == "" {
		return fmt.Errorf("autho field should be non-empty")
	}
	return nil
}
