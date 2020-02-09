package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		time := fmt.Sprintf("+%v", c.QueryParam("time"))

		var err error

		if time == "+0" {
			cmd := exec.Command("shutdown", "-c")
			err = cmd.Start()
		} else {
			cmd := exec.Command("shutdown", "-c")
			err = cmd.Start()
		}

		if err != nil {
			log.Fatal(err)
		}

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
