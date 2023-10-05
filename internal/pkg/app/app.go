package app

import (
	"Umbrella-test-task/internal/app/endpoint"
	"Umbrella-test-task/internal/app/mw"
	"Umbrella-test-task/internal/app/sevice"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *sevice.Service
	echo *echo.Echo
}

func (a *App) Run() error {
	fmt.Println("server running")
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func New() (*App, error) {
	a := &App{}

	a.s = sevice.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}
