package internal

import (
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/milo/internal/api"
	"html/template"
	"io"
	"net"
	"net/http"
	_ "github.com/casbin/casbin"
	_ "github.com/labstack/echo-contrib/casbin"
)

type HttpServer interface {
	StartServer(l net.Listener)
}

type httpServer struct {
	Core
	*echo.Echo
	templates *template.Template
}

type TemplateRenderer struct {
	templates *template.Template
}

func NewHttp(c Core) HttpServer {
	e := echo.New()
	cnt := e.AcquireContext()
	cnt.Set("config", c.GetSettings().GetOptions())
	e.ReleaseContext(cnt)
	return &httpServer{Core: c, Echo: e}
}

func (h *httpServer) StartServer(l net.Listener) {
	// Middleware
	h.Use(middleware.Logger())
	h.Use(middleware.Recover())
	h.Use(middleware.Gzip())
	h.Use(middleware.CORS())
	h.Use(middleware.CSRF())
	//h.Use(casbinmw.Middleware(casbin.NewEnforcer("./configs/auth_model.conf", "./configs/policy.csv")))

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
	h.Server.Serve(l)
}

// Render renders a template document
func (h *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return h.templates.ExecuteTemplate(w, name, data)
}
