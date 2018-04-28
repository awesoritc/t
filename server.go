package main

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./handler"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", handler.MainPage())
    e.GET("/test", func(c echo.Context) error {
        return c.String(http.StatusOK, "test")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
