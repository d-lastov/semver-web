package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, SemVer)
}
