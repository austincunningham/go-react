package main

import(
	"net/http"

	"github.com/labstack/echo"
)

func main(){
	router := echo.New()
	router.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "hello world!")
	})
	router.Logger.Fatal(router.Start(":1323"))
}