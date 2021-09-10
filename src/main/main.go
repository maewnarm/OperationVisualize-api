package main

import (
	"OperationVisualize/api"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	fmt.Println("Hey")
	api.ConnectDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/mcnamedata", api.GetDistinctMCname)
	e.GET("/recdata/:detail", api.GetRecordData)
	e.GET("/sumdata/:detail", api.GetSummaryData)

	e.Logger.Fatal(e.Start(":8080"))
}
