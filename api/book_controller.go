package api

import (
	"encoding/json"
	"fmt"
	"go-template/api/request"
	"go-template/api/response"
	"go-template/logger"
	"go-template/service"
	"go-template/utils"
	"net/http"
)

type BookController struct {
	svc *service.BookService
	log logger.StructLogger
}

func NewBookController(svc *service.BookService, log logger.StructLogger) *BookController {
	return &BookController{
		svc: svc,
		log: log,
	}
}

func (b *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	tid := utils.GetTracingID(r.Context())
	b.log.Println("CreateBook", tid, "request for creating book")

	bookReq := &request.BookReq{}
	err := json.NewDecoder(r.Body).Decode(&bookReq)
	if err != nil {
		b.log.Errorf("CreateBook", tid, "failed to decode request body: %s", err)
		response.ServeJSON(w, http.StatusBadRequest, nil, nil, "failed to decode request body", nil)
		return
	}

	err = request.ValidateBookCreateReq(bookReq)
	if err != nil {
		b.log.Errorf("CreateBook", tid, "validation error: %s", err)
		response.ServeJSON(w, http.StatusBadRequest, nil, nil, fmt.Sprintf("validation error: %s", err), nil)
		return
	}

	err = b.svc.CreateBook(r.Context(), bookReq)
	if err != nil {
		b.log.Errorf("CreateBook", tid, "failed to create book: %s", err)
		response.ServeJSON(w, http.StatusInternalServerError, nil, nil, fmt.Sprintf("failed to create book: %s", err), nil)
		return
	}

	response.ServeJSON(w, http.StatusOK, nil, nil, "created book successfully", nil)
}

func (b *BookController) ListBooks(w http.ResponseWriter, r *http.Request) {
	tid := utils.GetTracingID(r.Context())
	page, skip, limit, err := parseSkipLimit(r, 10, 100)
	if err != nil {
		b.log.Errorln("ListBooks", tid, err.Error())
		response.ServeJSON(w, http.StatusBadRequest, nil, nil, err.Error(), nil)
		return
	}
	b.log.Printf("ListBooks", tid, "request for getting list of books, page=%d, limit=%d", page, limit)

	result, err := b.svc.ListBooks(r.Context(), skip, limit+1)
	if err != nil {
		b.log.Errorf("ListBooks", tid, "failed to get books list: %s", err)
		response.ServeJSON(w, http.StatusInternalServerError, nil, nil, fmt.Sprintf("failed to get books list: %s", err), nil)
		return
	}

	prev, next := getNextPreviousPagerLink(r.URL.Path, page, limit, len(result))
	if len(result) > limit {
		result = result[:limit]
	}
	response.ServeJSON(w, http.StatusOK, prev, next, "get books list is successful", result)
}
