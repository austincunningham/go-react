package main

import (
	"net/http"

	"github.com/austincunningham/go-react/pkg/db"
	"github.com/austincunningham/go-react/pkg/handlers"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

var d = db.DBconnect()

func main() {
	defer d.Close()

	router := echo.New()
	// non database route
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!")
	})

	router.GET("/apps", handlers.GetAllApps)
	router.GET("/apps/:id", handlers.GetApp)
	router.PUT("/apps/:id", handlers.UpdateApp)
	router.POST("/apps", handlers.CreateApp)
	// router.DELETE("/apps/:id", DeleteApp)

	router.Logger.Fatal(router.Start(":8001"))
}

// // DeleteApp delete app by id app/:id
// func DeleteApp(c echo.Context) error {
// 	id := c.Param("id")
// 	sqlStatment := "DELETE FROM apps WHERE id = $1"
// 	res, err := d.Query(sqlStatment, id)
// 	if err != nil{
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(res)
// 		return c.JSON(http.StatusOK, "Deleted")
// 	}
// 	return c.JSON(http.StatusOK, "Deleted")
// }
