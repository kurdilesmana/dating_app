package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/kurdilesmana/dating_app/internal/domain"
	"github.com/kurdilesmana/dating_app/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
	Validator   *validator.Validate
}

func NewUserHandler(userService *service.UserService, validator *validator.Validate) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Validator:   validator,
	}
}

// Register handles user registration
func (h *UserHandler) Register(c echo.Context) error {
	// Parse request body
	var newUser domain.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	// Validate input using the validator
	if err := h.Validator.Struct(newUser); err != nil {
		validationErrors := map[string]string{}
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors[e.Field()] = e.Tag()
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "validation failed", "details": validationErrors})
	}

	// Register user
	err := h.UserService.RegisterUser(&newUser)
	if err != nil {
		// Handle registration error
		// You may want to return more specific error messages based on the error type
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to register user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "user registered successfully"})
}
