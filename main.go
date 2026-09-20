package main

import (
	"webapp/pkg/config"
	"webapp/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := config.GetEnv("PORT", "3000")
	logPath := config.GetEnv("LOG_PATH", "/app/log/app.log")

	logFile, err := logger.CreateLogger(logPath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
	}))
	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		return c.File("public/views/webapp.html")
	})
	e.Start(":" + port)
}
