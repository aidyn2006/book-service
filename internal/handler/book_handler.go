package handler

import (
	"book-service/internal/entity"
	"book-service/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type BookHandler struct {
	uc usecase.BookUsecase
}

func NewBookHandler(uc usecase.BookUsecase) *BookHandler {
	return &BookHandler{uc}
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	book, err := h.uc.GetBook(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	json.NewDecoder(r.Body).Decode(&book)
	err := h.uc.AddBook(&book)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
