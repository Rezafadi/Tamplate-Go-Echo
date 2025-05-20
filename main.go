package main

import (
	"log"
	"project-name/app/router"
	"project-name/config"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"gopkg.in/tylerb/graceful.v1"
)

// @title PROJECT_NAME
// @version 1.0
// @description API documentation by PROJECT_NAME

// @securityDefinitions.apikey JwtToken
// @in header
// @name Authorization

func main() {
	app := echo.New()
	config.Database()

	// config.Redis()
	router.Init(app)

	// activateCron()

	app.Server.Addr = "0.0.0.0:" + config.LoadConfig().Port
	log.Printf("Server: " + config.LoadConfig().BaseUrl)

	if !config.LoadConfig().IsDesktop {
		log.Printf("Documentation: " + config.LoadConfig().BaseUrl + "/api-docs")
	}

	graceful.ListenAndServe(app.Server, 5*time.Second)
}

func activateCron() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	job := cron.New(cron.WithLocation(loc))

	job.Start()
}
