package router

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"project-name/app/controllers"
	"project-name/app/middlewares"
	"project-name/config"
	_ "project-name/docs" // For Swagger

	"log"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Init Router
func Init(app *echo.Echo) {
	if !config.LoadConfig().IsDesktop {
		renderer := &TemplateRenderer{
			templates: template.Must(template.ParseGlob("*.html")),
		}
		app.Renderer = renderer
	}
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())

	if !config.LoadConfig().IsDesktop {
		app.GET("/swagger/*", echoSwagger.WrapHandler)
		app.GET("/api-docs", func(c echo.Context) error {
			err := c.Render(http.StatusOK, "docs.html", map[string]interface{}{
				"BaseUrl": config.LoadConfig().BaseUrl,
				"Title":   "Api Documentation of " + config.LoadConfig().AppName,
			})
			fmt.Println("err:", err)
			return err
		})
	}

	app.Static("/assets", "assets")

	api := app.Group("/v1", middlewares.StripHTMLMiddleware, middlewares.CheckAPIKey())
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login/user", controllers.LoginUser)
			auth.POST("/login/admin", controllers.LoginAdmin)
			auth.POST("/register", controllers.Register)
			auth.POST("/forgot-password", controllers.ForgotPassword)
			auth.POST("/email-verify", controllers.SendEmailVerifyEmail, middlewares.Auth())
			auth.PUT("/activate-account/:id", controllers.AktivateAccount, middlewares.Auth())
			auth.PUT("/change-password-login", controllers.ChangePasswordLogin, middlewares.Auth())
			auth.PUT("/reset-password/:id", controllers.ResetPassword)
		}

	}

	log.Printf("Server started...")
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
