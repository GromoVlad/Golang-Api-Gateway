package listBookResponse

import "gin_tonic/internal/models/book"

type Response struct {
	Data    ListBookResponse `json:"data"`
	Success bool             `json:"success"`
}

type ListBookResponse struct {
	CurrentPage int         `json:"current_page"  example:"1"`
	Limit       int         `json:"limit"         example:"10"`
	Books       []book.Book `json:"books"`
}
