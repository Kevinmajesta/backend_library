package handler

import (
	"net/http"

	"github.com/Kevinmajesta/backend_library/internal/entity"
	"github.com/Kevinmajesta/backend_library/internal/http/binder"
	"github.com/Kevinmajesta/backend_library/internal/service"
	"github.com/Kevinmajesta/backend_library/pkg/response"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	input := binder.UserCreateRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "there is an input error"))
	}

	if h.userService.EmailExists(input.Email) {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "email is already in use"))
	}

	newUser := entity.NewUser(input.Fullname, input.Email, input.Phone)
	user, err := h.userService.CreateUser(newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully created a new user, the email has been sent", user))
}
