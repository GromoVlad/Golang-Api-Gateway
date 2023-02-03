package findBookResponse

import (
	"gin_tonic/internal/models/book"
)

type Response struct {
	Data    book.Book `json:"data"`
	Success bool      `json:"success"`
}
