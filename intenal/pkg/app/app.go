package app

import (
	"github.com/fishmanDK/good_traning/intenal/app/endpoint"
	"github.com/fishmanDK/good_traning/intenal/app/service"
	"github.com/fishmanDK/good_traning/intenal/mw"
	"github.com/labstack/echo"
)

type App struct{
	s *service.Service
	e *endpoint.Endpoint
	echo *echo.Echo
}

func NewApp() (*App, error){
	a := &App{}
	
	a.s = service.NewService()
	a.e = endpoint.NewEndpoint(a.s)

	a.echo = echo.New()
	a.echo.Use()

	a.echo.GET("test", a.e.Status, mw.RoleCheck)
	// a.Run()

	return a, nil
}

func (a *App) Run() error{
	err := a.echo.Start(":8080")
	if err != nil{
		return err
	}

	return nil
}