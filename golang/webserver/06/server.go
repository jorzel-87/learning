package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your cat name is %s\nand his type is %s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "Please make a better request..",
	})
}

func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarshalling: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your cat: %#v", cat)
	return c.String(http.StatusOK, "we got your cat")
}

func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing addDog: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your dog: %#v", dog)
	return c.String(http.StatusOK, "we got your dog")
}

func addHamster(c echo.Context) error {
	hamster := Hamster{}
	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addHamster: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("This is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster")

}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the secret admin main page!")
}

////////////////////////////// middlewares //////////////////////////////

// it will add to every response of the server, server name
// middleware gets echo.HandlerFunc and returns echo.HandlerFunc
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "JorzelSrv/1.0 ")
		c.Response().Header().Set("notReallyHeader", "thisHaveNoMeaning")

		return next(c)
	}
}

func main() {
	e := echo.New()

	e.Use(ServerHeader)

	g := e.Group("/admin")

	// middleware usage - logging server interaction
	// only for given group
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if username == "jack" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/main", mainAdmin)
	e.GET("/", hello)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/hamsters", addHamster)
	e.Start(":1323")
}
