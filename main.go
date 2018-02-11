package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.Static("/", "synkron-web-angular/dist/")
	e.Logger.Fatal(e.Start(":8080"))
}
