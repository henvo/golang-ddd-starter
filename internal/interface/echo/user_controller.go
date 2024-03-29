package echo

import (
	"net/http"

	"github.com/henvo/golang-ddd-starter/internal/domain"
	"github.com/labstack/echo/v4"
)

// The main user handler.
type UserController struct {
	userService domain.UserService
}

// NewUserController creates a new user controller
func NewUserController(service domain.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// Apply adds all the necessary handlers to the main web server instance.
func (uc *UserController) Apply(ws *webServer) {
	group := ws.echo.Group("/users")

	group.GET("*", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
