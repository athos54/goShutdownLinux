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
			cmd := exec.Command("shutdown", time)
			err = cmd.Start()
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Apagado en %v!", time))
		return c.String(http.StatusOK, fmt.Sprintf("Apagado en %v!", time))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
