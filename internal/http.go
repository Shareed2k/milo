package internal

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

type HttpServer interface {
	StartServer()
}

type httpServer struct {
	Settings
	e         *echo.Echo
	templates *template.Template
}

type TemplateRenderer struct {
	templates *template.Template
}

func NewHttp(s Settings) HttpServer {
	e := echo.New()
	return &httpServer{Settings: s, e: e}
}

func (h *httpServer) StartServer() {
	s := h.GetOptions()

	// Middleware
	h.e.Use(middleware.Logger())
	h.e.Use(middleware.Recover())
	h.e.Use(middleware.Gzip())
	h.e.Use(middleware.CORS())
	h.e.Use(middleware.CSRF())

	// Routes
	box := rice.MustFindBox("../ui")

	assetHandler := http.FileServer(box.HTTPBox())
	// serves the index.html from rice
	h.e.GET("/", echo.WrapHandler(assetHandler))

	// api
	api := h.e.Group("/api")
	api.Use(middleware.JWT([]byte("secret")))

	api.GET("/test", hello)

	//TODO: need to find way where to store routes

	// get file contents as string
	templateString, _ := box.String("index.html")

	// parse and execute the template
	tmpl, _ := template.New("index").Parse(templateString)

	renderer := &TemplateRenderer{
		templates: tmpl,
	}
	h.e.Renderer = renderer

	h.e.GET("/test2", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"csrf": c.Get("csrf"),
		})
	})

	// Start Server
	h.e.Start(fmt.Sprintf(":%d", s.HttpPort))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Render renders a template document
func (h *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return h.templates.ExecuteTemplate(w, name, data)
}
