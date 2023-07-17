package mw

import (
	"log"

	"github.com/labstack/echo"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc{
	return func(ctx echo.Context) error{
		val := ctx.Request().Header.Get("Role")
		if val == "admin"{
			log.Println("this admin")
		} else{
			log.Println("it is not admin :(")
		}

		err := next(ctx)
		if err != nil{
			return err
		}
		return nil
	}
}