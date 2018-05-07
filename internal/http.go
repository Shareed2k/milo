package internal

import (
	"github.com/GeertJohan/go.rice"
	_ "github.com/casbin/casbin"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo-contrib/casbin"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net"
	"net/http"
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

type MiloContext struct {
	echo.Context
	Core
}

func NewHttp(c Core) HttpServer {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	return &httpServer{Core: c, Echo: e}
}

func (h *httpServer) StartServer(l net.Listener) {
	// Middleware
	h.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &MiloContext{Context: c, Core: h.Core}
			return next(cc)
		}
	})
	h.Use(middleware.Logger())
	h.Use(middleware.Recover())
	h.Use(middleware.Gzip())
	h.Use(middleware.CORS())
	h.Use(middleware.CSRF())
	//h.Use(casbinmw.Middleware(casbin.NewEnforcer("./configs/auth_model.conf", "./configs/policy.csv")))

	// Routes
	NewRoutes(h.Echo)

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

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
