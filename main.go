package main

import (
	"net/http"
	"github.com/austincunningham/go-react/pkg/handlers"
	"github.com/labstack/echo"
	)



func main() {
	router := echo.New()
	// non database route
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!")
	})

	router.GET("/apps", handlers.GetAllApps)
	router.GET("/apps/:id", handlers.GetApp)
	router.PUT("/apps/:id", handlers.UpdateApp)
	router.POST("/apps", handlers.CreateApp)
	router.DELETE("/apps/:id", handlers.DeleteApp)

	router.Logger.Fatal(router.Start(":8001"))
}