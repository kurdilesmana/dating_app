package router

import (
	"github.com/go-playground/validator/v10"
	handler "github.com/kurdilesmana/dating_app/internal/delivery/http/handlers"
	"github.com/kurdilesmana/dating_app/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize service
	var userService service.UserService

	// Initialize validator
	validator := validator.New()

	// Initialize user handler with validator
	userHandler := handler.NewUserHandler(&userService, validator)

	// Define routes
	e.POST("/register", userHandler.Register)

	return e
}
