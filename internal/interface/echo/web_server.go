package echo

import "github.com/labstack/echo/v4"

type webServer struct {
	echo *echo.Echo
}

type controller interface {
	Apply(*webServer)
}

// NewWebServer creates a new wrapper for the echo framework.
func NewWebServer() *webServer {
	// Create a new echo server.
	srv := echo.New()
	return &webServer{echo: srv}
}

// Add applies the handlers from the controller.
func (ws *webServer) Add(c controller) {
	c.Apply(ws)
}

// Start starts the server on a given port.
func (ws *webServer) Start(port string) {
	ws.echo.Start(port)
}
