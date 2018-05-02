package internal

import (
	"github.com/labstack/echo"
	"fmt"
)

type Http interface {
	StartServer()
}

type http struct {
	Settings
	e *echo.Echo
}

func NewHttp(s Settings) Http {
	e := echo.New()
	return &http{s, e}
}

func (h *http) StartServer() {
	s := h.GetOptions()

	// Start Server
	h.e.Start(fmt.Sprintf(":%d", s.HttpPort))
}
