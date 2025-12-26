package handler

import (
	"errors"
	"net/http"

	"github.com/Kevinmajesta/backend_library/internal/http/binder"
	"github.com/Kevinmajesta/backend_library/internal/service"
	"github.com/Kevinmajesta/backend_library/pkg/response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BorrowHandler struct {
	borrowService service.BorrowService
}

func NewBorrowHandler(borrowService service.BorrowService) BorrowHandler {
	return BorrowHandler{
		borrowService: borrowService,
	}
}

func (h *BorrowHandler) BorrowBook(c echo.Context) error {
	var input binder.BorrowRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.ErrorResponse(http.StatusBadRequest, "invalid request body"),
		)
	}

	borrow, err := h.borrowService.BorrowBook(input.UserID, input.BookID)
	if err != nil {

		if errors.Is(err, service.ErrBookOutOfStock) {
			return c.JSON(http.StatusBadRequest, response.CustomErrorResponse{
				Message:        "Stok buku habis",
				ZiyadErrorCode: "ZYD-ERR-001",
				TraceID:        response.GenerateTraceID(12),
			})
		}

		// system error
		traceID := response.GenerateTraceID(12)
		return c.JSON(http.StatusInternalServerError, response.CustomErrorResponse{
			Message:        "Terjadi kesalahan pada sistem",
			ZiyadErrorCode: "ZYD-ERR-999",
			TraceID:        traceID,
		})
	}

	return c.JSON(
		http.StatusOK,
		response.SuccessResponse(
			http.StatusOK,
			"book successfully borrowed",
			borrow,
		),
	)
}

func (h *BorrowHandler) ReturnBook(c echo.Context) error {
	borrowIDParam := c.Param("borrow_id")

	borrowID, err := uuid.Parse(borrowIDParam)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.ErrorResponse(http.StatusBadRequest, "invalid borrow id"),
		)
	}

	if err := h.borrowService.ReturnBook(borrowID); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.ErrorResponse(http.StatusBadRequest, err.Error()),
		)
	}

	return c.JSON(
		http.StatusOK,
		response.SuccessResponse(
			http.StatusOK,
			"book successfully returned",
			nil,
		),
	)
}
