package handler

import (
	"net/http"

	"github.com/Kevinmajesta/backend_library/internal/entity"
	"github.com/Kevinmajesta/backend_library/internal/http/binder"
	"github.com/Kevinmajesta/backend_library/internal/service"
	"github.com/Kevinmajesta/backend_library/pkg/response"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) BookHandler {
	return BookHandler{bookService: bookService}
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	input := binder.BookCreateRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}

	if h.bookService.BookExists(input.Title) {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "title is already in use"))
	}

	newBook := entity.NewBook(input.Title, input.Stock)
	book, err := h.bookService.CreateBook(newBook)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully created a new book", book))
}

func (h *BookHandler) FindBook(c echo.Context) error {
	book_ID := c.Param("book_id")
	book, err := h.bookService.FindBookByID(book_ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, "book not found"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully found the book", book))
}
