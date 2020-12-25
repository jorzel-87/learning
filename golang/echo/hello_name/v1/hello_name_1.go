package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// User
	type User struct {
		Name  string `json:"name" form:"name" query:"name"`
		Email string `json:"email" form:"email" query:"email"`
	}

	// Route => handler
	//	e.GET("/", func(c echo.Context) error {
	//		return c.String(http.StatusOK, "Czesc Marta!\n")
	//	})

	//	e.GET("/hello", func(c echo.Context) (err error) {
	//		u := new(User)
	//		if err = c.Bind(u); err != nil {
	//			return
	//		}
	//		return c.JSON(http.StatusOK, u)
	//	})

	e.POST("/hello", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s!\n", u ))
	})




	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
