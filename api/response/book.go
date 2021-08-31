package response

import (
	"go-template/model"
	"time"
)

type BookResp struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func ToBookResp(in []model.Book) []BookResp {
	out := make([]BookResp, len(in))
	for _, item := range in {
		out = append(out, BookResp{
			ID:        item.ID,
			Name:      item.Name,
			Author:    item.Author,
			UpdatedAt: item.UpdatedAt,
			CreatedAt: item.CreatedAt,
		})
	}
	return out
}
