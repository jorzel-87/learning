package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

func name(c echo.Context) error {
	name := c.QueryParam("name")

	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!\n", name))
}

func main() {

	e := echo.New()

	e.GET("/hello", name)

	e.Start(":1323")
}
