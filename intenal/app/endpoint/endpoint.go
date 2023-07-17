package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface{
	TimeToday() string
}

type Endpoint struct{
	s Service
}

func NewEndpoint(serv Service) *Endpoint{
	return &Endpoint{
		s: serv,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error{
	date := e.s.TimeToday()
	
	msg := fmt.Sprint("fishman D.", date)
	err := ctx.String(http.StatusOK, msg)
	if err != nil{
		return err
	}
	
	return nil
}