package main

import (
    "github.com/labstack/echo/v4"
	"webapp/pkg/config"
)

func main() {
	port := config.GetEnv("PORT", "3000")

    e := echo.New()
    e.Static("/", "public")
    
    e.GET("/", func(c echo.Context) error {
        return c.File("public/views/webapp.html")
    })
    e.Start(":" + port)
}

package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"webapp/pkg/config"
)

func main() {
	port := config.GetEnv("PORT", "3000")
	logPath := config.GetEnv("LOG_PATH", "/var/log/webapp/app.log")

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

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