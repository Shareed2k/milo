package internal

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
	"github.com/milo/internal/api"
)

type HttpServer interface {
	StartServer()
}

type httpServer struct {
	Settings
	*echo.Echo
	templates *template.Template
}

type TemplateRenderer struct {
	templates *template.Template
}

func NewHttp(s Settings) HttpServer {
	e := echo.New()
	cnt := e.AcquireContext()
	cnt.Set("config", s.GetOptions())
	e.ReleaseContext(cnt)
	return &httpServer{Settings: s, Echo: e}
}

func (h *httpServer) StartServer() {
	s := h.GetOptions()

	// Middleware
	h.Use(middleware.Logger())
	h.Use(middleware.Recover())
	h.Use(middleware.Gzip())
	h.Use(middleware.CORS())
	h.Use(middleware.CSRF())

	// Routes
	api.NewRoutes(h.Echo)

	box := rice.MustFindBox("../ui/dist")
	tmplBox := rice.MustFindBox("../ui/src/tmpl")

	assetHandler := http.FileServer(box.HTTPBox())

	h.GET("/static/*", echo.WrapHandler(http.StripPrefix("/", assetHandler)))

	// get file contents as string
	templateString, _ := tmplBox.String("index.html")

	// parse and execute the template
	tmpl, _ := template.New("index").Parse(templateString)

	renderer := &TemplateRenderer{
		templates: tmpl,
	}
	h.Renderer = renderer

	h.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", map[string]interface{}{
			"csrf": c.Get("csrf"),
		})
	})

	// Start Server
	h.Start(fmt.Sprintf(":%d", s.HttpPort))
}

// Render renders a template document
func (h *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return h.templates.ExecuteTemplate(w, name, data)
}
