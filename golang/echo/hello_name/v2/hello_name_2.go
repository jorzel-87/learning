package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func hello(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}

	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!\n", u))
}

func name(c echo.Context) error {
	name := c.QueryParam("name")

	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!\n", name))
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/hello", hello)
	e.POST("/name", name)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
